package main

import (
	"flag"
	"github.com/stevencorona/elastic-haproxy/elb"
	"github.com/stevencorona/elastic-haproxy/haproxy"
	"github.com/stevencorona/elastic-haproxy/statsd"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var defaultConfigFile = "config/elastic.toml"
var flagConfigFile string

func main() {

	config := haproxy.Transform()

	flag.StringVar(&flagConfigFile, "configFile", defaultConfigFile, "Path to toml file")
	flag.Parse()

	conf := LoadConfig(flagConfigFile)

	server := new(haproxy.Server)

	// We use two channels— one to send actions to the server and one to recieve
	// notifications from it. Create them right now.
	actionChan := make(chan haproxy.Action)
	notificationChan := make(chan haproxy.Event)

	// Handle signals gracefully in another goroutine
	go gracefulSignals(server, actionChan, notificationChan)

	// Start up the HAProxy Server
	go server.Start(notificationChan, actionChan)

	// Setup the ELB HTTP Handlers
	go elb.SetupApiHandlers()

	// Fire up statsd goroutine if statsd is enabled. This might be better off in
	// a seperate binary to monitor HAProxy.
	if conf.Statsd.Enabled {
		go statsd.SendMetrics(server)
	}

	// Event loop for handling events from the HAProxy server
	// (right now, it only sends start/stop notifications)
	for {
		<-notificationChan
		log.Println("Received a notification")
		time.Sleep(2 * time.Second)

		server.Socket = conf.Haproxy.Socket
		serverInfo := server.GetInfo()
		log.Println(serverInfo)
	}

}

// Capture, capture, capture those signals. We gracefully shutdown on a
// SIGQUIT (this is what docker sends w/ docker stop) and everything else
// causes a graceful (zero-downtime) restart of HAProxy. Remember, HAProxy
// does not support reloads or graceful restarts without some... trickery.
func gracefulSignals(server *haproxy.Server, actionChan chan haproxy.Action, notificationChan chan haproxy.Event) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT)

	for {
		s := <-signals
		log.Println("Received a signal", s)

		if s == syscall.SIGQUIT {
			log.Println("Caught SIGQUIT, Stopping HAProxy")

			// Tell server to stop and wait for a response
			actionChan <- haproxy.WantsStop

			// Race condition, this exits before we stop :( It should wait!
			<-notificationChan
			os.Exit(1)
		} else {
			actionChan <- haproxy.WantsReload
		}
	}
}

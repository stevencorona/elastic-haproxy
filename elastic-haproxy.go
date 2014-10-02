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
)

var defaultConfigFile = "config/elastic.toml"
var flagConfigFile string

func main() {

	// Read in the configuration file settings. I think we should instead
	// serialize to JSON and treat this like a view.
	haproxy.Transform()

	flag.StringVar(&flagConfigFile, "configFile", defaultConfigFile, "Path to toml file")
	flag.Parse()

	conf, err := LoadConfig(flagConfigFile)

	if err != nil {
		log.Fatal("Could not load configuration", err)
	}

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
	go elb.InitApiHandlers()

	// Fire up statsd goroutine if statsd is enabled. This might be better off in
	// a seperate binary to monitor HAProxy.
	if conf.Statsd.Enabled {
		go statsd.SendMetrics(server)
	}

	// Event loop for handling events from the HAProxy server
	// (right now, it only sends start/stop notifications)
	for _ = range notificationChan {
		log.Println("Received a notification")

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

	for s := range signals {
		log.Println("Received a signal", s)

		switch s {
		case syscall.SIGQUIT:
			log.Println("Caught SIGQUIT, Stopping HAProxy")
			actionChan <- haproxy.WantsStop
			<-notificationChan
			os.Exit(1)
		default:
			actionChan <- haproxy.WantsReload
		}
	}
}

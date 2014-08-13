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

	haproxy.Transform()
	os.Exit(1)

	flag.StringVar(&flagConfigFile, "configFile", defaultConfigFile, "Path to toml file")
	flag.Parse()

	conf := LoadConfig(flagConfigFile)

	server := new(haproxy.Server)

	// We use two channels— one to send actions to the server and one to recieve
	// notifications from it. Create them right now.
	actionChan := make(chan haproxy.Action)
	notificationChan := make(chan haproxy.Event)

	// Handle signals gracefully in another goroutine
	go gracefulSignals(server)

	// Start up the HAProxy Server
	go server.Start(notificationChan, actionChan)

	// Setup the ELB HTTP Handlers
	go elb.SetupApiHandlers()

	// Fire up statsd goroutine if statsd is enabled. This might be better off in
	// a seperate binary to monitor HAProxy.
	if conf.Statsd.Enabled {
		go statsd.SendMetrics(server)
	}

	for {
		<-notificationChan
		log.Println("Received a notification")
		time.Sleep(2 * time.Second)

		server.Socket = conf.Haproxy.Socket
		serverInfo := server.GetInfo()
		log.Println(serverInfo)
	}
}

func gracefulSignals(server *haproxy.Server) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT)

	for {
		s := <-signals
		log.Println("Received a signal", s)

		if s == syscall.SIGQUIT {
			log.Println("Caught SIGQUIT, Stopping HAProxy")
			server.ActionChan <- haproxy.WantsStop
			os.Exit(1)
		}

		server.ActionChan <- haproxy.WantsReload
	}
}

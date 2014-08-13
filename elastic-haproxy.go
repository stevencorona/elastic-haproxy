package main

import (
	"flag"
	"fmt"
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

	notificationChan := make(chan haproxy.Event)
	actionChan := make(chan haproxy.Action)

	// Handle signals gracefully in another goroutine
	go gracefulSignals(server)

	// Start up the HAProxy Server
	go server.Start(notificationChan, actionChan)

	// Setup the ELB HTTP Handlers
	go elb.SetupApiHandlers()

	if conf.Statsd.Enabled {
		go statsd.SendMetrics(server)
	}

	for {
		<-notificationChan
		fmt.Println("Got notification")
		time.Sleep(2 * time.Second)

		server.Socket = conf.Haproxy.Socket
		serverInfo := server.GetInfo()
		fmt.Println(serverInfo)
	}
}

func gracefulSignals(server *haproxy.Server) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT)

	for {
		s := <-signals
		log.Println("Got signal:", s)

		if s == syscall.SIGQUIT {
			fmt.Println("caught sigquit")
			server.ActionChan <- haproxy.WantsStop
			os.Exit(1)
		}

		server.ActionChan <- haproxy.WantsReload
	}
}

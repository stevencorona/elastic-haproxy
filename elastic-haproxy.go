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

	go gracefulSignals(server)
	go server.Start(notificationChan, actionChan)
	go elb.SetupApiHandlers()

	for {
		<-notificationChan
		fmt.Println("Got notification")
		time.Sleep(2 * time.Second)

		server.Socket = conf.Haproxy.Socket
		serverInfo := server.GetInfo()
		fmt.Println(serverInfo)

		if conf.Statsd.Enabled {
			go statsd.SendMetrics(server)
		}
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

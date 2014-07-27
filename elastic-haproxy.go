package main

import (
	"flag"
	"fmt"
	"github.com/stevencorona/elastic-haproxy/elb"
	"github.com/stevencorona/elastic-haproxy/haproxy"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var defaultConfigFile = "config/elastic.toml"
var flagConfigFile string

func main() {

	flag.StringVar(&flagConfigFile, "configFile", defaultConfigFile, "Path to toml file")
	flag.Parse()

	conf := LoadConfig(flagConfigFile)

	haproxy := new(haproxy.Server)

	notificationChan := make(chan int)
	shouldReloadChan := make(chan int)

	go gracefulSignals(haproxy)
	go haproxy.Start(notificationChan, shouldReloadChan)
	go elb.SetupApiHandlers()

	for {
		<-notificationChan
		fmt.Println("Got notify")
		time.Sleep(2 * time.Second)

		haproxy.Socket = conf.HaproxySocket
		serverInfo := haproxy.GetInfo()
		fmt.Println(serverInfo)
	}
}

func gracefulSignals(haproxy *haproxy.Server) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT)

	for {
		s := <-signals
		log.Println("Got signal:", s)

		if s == syscall.SIGQUIT {
			fmt.Println("caught sigquit")
			haproxy.ActionChan <- 8
			os.Exit(1)
		}

		haproxy.ActionChan <- 4
	}
}

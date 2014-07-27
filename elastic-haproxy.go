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
)

var defaultConfigFile = "config/elastic.toml"
var flagConfigFile string

func main() {

	flag.StringVar(&flagConfigFile, "configFile", defaultConfigFile, "Path to toml file")
	flag.Parse()

	conf := LoadConfig(flagConfigFile)

	haproxy := new(haproxy.Server)

	startChannel := make(chan int)
	stopChannel := make(chan int)

	go gracefulSignals(haproxy)
	go haproxy.Start(startChannel, stopChannel)

	<-startChannel

	haproxy.Socket = conf.HaproxySocket
	serverInfo := haproxy.GetInfo()
	fmt.Println(serverInfo)
	elb.SetupApiHandlers()
}

func gracefulSignals(haproxy *haproxy.Server) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	for {
		s := <-signals
		log.Println("Got signal:", s)
		startChannel := make(chan int)
		stopChannel := make(chan int)
		go haproxy.Start(startChannel, stopChannel)
	}
}

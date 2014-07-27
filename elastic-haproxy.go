package main

import (
	"flag"
	"fmt"
	"github.com/stevencorona/elastic-haproxy/haproxy"
)

var flagConfigFile string

func main() {

	flag.StringVar(&flagConfigFile, "configFile", "config.toml", "Path to toml file")
	flag.Parse()

	conf := LoadConfig(flagConfigFile)

	haproxy := new(haproxy.Server)
	haproxy.Socket = conf.HaproxySocket

	serverInfo := haproxy.GetInfo()

	fmt.Println(serverInfo)

	SetupApiHandlers()

}

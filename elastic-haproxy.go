package main

import (
	"flag"
	"fmt"
	"github.com/stevencorona/elastic-haproxy/elb"
	"github.com/stevencorona/elastic-haproxy/haproxy"
)

var defaultConfigFile = "config/elastic.toml"
var flagConfigFile string

func main() {

	flag.StringVar(&flagConfigFile, "configFile", defaultConfigFile, "Path to toml file")
	flag.Parse()

	conf := LoadConfig(flagConfigFile)

	haproxy := new(haproxy.Server)
	haproxy.Socket = conf.HaproxySocket

	serverInfo := haproxy.GetInfo()

	fmt.Println(serverInfo)

	elb.SetupApiHandlers()

}

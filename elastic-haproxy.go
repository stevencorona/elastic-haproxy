package main

import (
	"flag"
	"fmt"
)

var flagConfigFile string

func main() {

	flag.StringVar(&flagConfigFile, "configFile", "config.toml", "Path to toml file")
	flag.Parse()

	conf := LoadConfig(flagConfigFile)

	haproxy := new(Haproxy)
	haproxy.Socket = conf.HaproxySocket

	serverInfo := haproxy.GetInfo()

	fmt.Println(serverInfo)

}

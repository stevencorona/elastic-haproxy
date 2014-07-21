package main

import (
	"fmt"
)

func main() {

	conf := LoadConfig("config.toml")

	haproxy := new(Haproxy)
	haproxy.Socket = conf.HaproxySocket

	serverInfo := haproxy.GetInfo()

	fmt.Println(serverInfo)

}

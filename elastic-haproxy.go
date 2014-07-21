package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

type Config struct {
	EnableStad       bool
	EnableAutosclale bool
	EnableRoute53    bool
	HaproxySocket    string
}

func main() {
	var conf Config
	if _, err := toml.DecodeFile("config.toml", &conf); err != nil {
		fmt.Println(err)
		return
	}

	haproxy := new(Haproxy)
	haproxy.Socket = conf.HaproxySocket

	serverInfo := haproxy.GetInfo()

	fmt.Println(serverInfo)

}

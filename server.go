package main

import (
	"bufio"
	"fmt"
	"github.com/BurntSushi/toml"
	"net"
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

	conn, err := net.Dial("unix", conf.HaproxySocket)

	if err != nil {
		fmt.Println(err)
	}

	conn.Write([]byte("show info\n"))

	reader := bufio.NewReader(conn)

	for {
		status, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(status)
	}

}

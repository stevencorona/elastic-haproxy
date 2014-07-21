package main

import (
	"bufio"
	"fmt"
	"github.com/BurntSushi/toml"
	"net"
	"reflect"
	"strconv"
	"strings"
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

	server := new(HaproxyServer)

	val := reflect.ValueOf(server).Elem()

	for {
		status, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println(err)
			break
		}

		parts := strings.Split(status, ":")

		key := parts[0]
		value := ""

		if strings.TrimSpace(key) == "" {
			continue
		}

		if len(parts) == 2 {
			value = strings.TrimSpace(parts[1])
		}

		for i := 0; i < val.NumField(); i++ {
			valueField := val.Field(i)
			typeField := val.Type().Field(i)
			tag := typeField.Tag

			if tag.Get("haproxy") == key {

				if valueField.Kind() == reflect.String {
					valueField.SetString(value)
				}

				if valueField.Kind() == reflect.Int {
					i, _ := strconv.Atoi(value)
					valueField.SetInt(int64(i))
				}
			}
		}

		//fmt.Println(key, ":", value)
	}

	fmt.Println(server)

}

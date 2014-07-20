package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

type Config struct {
	EnableStad       bool
	EnableAutosclale bool
	EnableRoute53    bool
}

func main() {
	var conf Config
	if _, err := toml.DecodeFile("config.toml", &conf); err != nil {
		fmt.Println(err)
		return
	}
}

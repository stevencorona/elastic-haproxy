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

func LoadConfig(path string) (config *Config) {
	if _, err := toml.DecodeFile(flagConfigFile, &config); err != nil {
		fmt.Println(err)
		return
	}

	return config
}

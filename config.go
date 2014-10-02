package main

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Haproxy   haproxyConfig
	Cluster   clusterConfig
	Statsd    statsdConfig
	Autoscale autoscaleConfig
	Route53   route53Config
	Api       apiConfig
	Dashboard dashboardConfig
}

type haproxyConfig struct {
	Socket string
	Binary string
}

type clusterConfig struct {
	Enabled bool
}

type statsdConfig struct {
	Enabled  bool
	Hostname string
	Port     int
}

type autoscaleConfig struct {
	Enabled bool
}

type route53Config struct {
	Enabled bool
}

type apiConfig struct {
	Enabled bool
}

type dashboardConfig struct {
	Enabled bool
}

func LoadConfig(path string) (config *Config, err error) {
	if _, err := toml.DecodeFile(flagConfigFile, &config); err != nil {
		return nil, err
	}

	return config, nil
}

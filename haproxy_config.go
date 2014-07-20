package main

type ConfigFile struct {
	Blocks map[string]ConfigBlock
}

type ConfigBlock struct {
	Log      string
	Mode     string
	Options  map[string]string
	Retries  int
	Balance  string
	Timeouts map[string]int
	Servers  map[string]string
}

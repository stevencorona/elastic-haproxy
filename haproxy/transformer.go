package haproxy

import (
	"fmt"
)

func Transform() {
	config := new(Config)
	configData := make([]string, 0)

	global := ParseGlobalBlock(config.Global)
	configData = append(configData, global[0:]...)

	if len(config.Frontends) > 0 {
		frontend := ParseFrontendBlock(config.Frontends[0])
		configData = append(configData, frontend[0:]...)
	}

	fmt.Println(configData)
}

func ParseGlobalBlock(global GlobalBlock) (data []string) {
	data = append(data, "global")

	data = ConfigBool("daemon", global.Daemon, data)
	data = ConfigInt("maxconn %d", global.MaxConn, data)

	data = ConfigString("ca-base %s", global.CaBase, data)
	data = ConfigString("chroot %s", global.Chroot, data)
	data = ConfigString("crt-base %s", global.CrtBase, data)

	data = ConfigInt("gid %d", global.Gid, data)
	data = ConfigString("group %s", global.Group, data)
	data = ConfigString("user %s", global.User, data)
	data = ConfigInt("uid %d", global.Uid, data)

	data = ConfigInt("nbproc %d", global.Nbproc, data)

	data = ConfigString("log %s", global.Log, data)
	data = ConfigString("log-tag %s", global.LogTag, data)
	data = ConfigStringAllowBlank("log-send-hostname %s", global.LogSendHostname, data)
	data = ConfigString("pidfile %s", global.Pidfile, data)

	data = ConfigString("stats bind-process %s", global.StatsBind, data)
	data = ConfigString("stats socket %s", global.StatsSocket, data)
	data = ConfigString("stats timeout %s", global.StatsTimeout, data)

	data = ConfigString("ssl-default-bind-ciphers %s", global.SslDefaultBindCiphers, data)
	data = ConfigString("ssl-server-verify %s", global.SslServerVerify, data)

	data = ConfigInt("ulimit-n %d", global.UlimitN, data)

	data = ConfigString("node %s", global.Node, data)
	data = ConfigString("description %s", global.Description, data)
	data = ConfigString("unix-bind %s", global.UnixBind, data)

	return data
}

func ParseFrontendBlock(frontend FrontendBlock) (data []string) {
	data = make([]string, 0)
	data = ConfigString("frontend %s", frontend.Name, data)

	for _, bind := range frontend.Binds {
		data = ConfigString("bind %s", bind.IpAddress, data)
	}

	return data
}

// func ConfigString(format string, data []string, settings ...string)

func ConfigString(format string, setting string, data []string) []string {
	if setting != "" {
		data = append(data, fmt.Sprintf(format, setting))
	}

	return data
}

func ConfigStringAllowBlank(format string, setting string, data []string) []string {
	data = append(data, fmt.Sprintf(format, setting))
	return data
}

func ConfigBool(format string, setting bool, data []string) []string {
	if setting {
		data = append(data, fmt.Sprintf(format, setting))
	}

	return data
}

func ConfigInt(format string, setting int, data []string) []string {
	data = append(data, fmt.Sprintf(format, setting))
	return data
}

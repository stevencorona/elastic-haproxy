package haproxy

import (
	"fmt"
	"strings"
)

func Transform() {
	config := new(Config)
	configFile := make([]string, 0)

	global := ParseGlobalBlock(config.Global)

	configFile = append(configFile, global[0:]...)
	fmt.Println(strings.Join(configFile, "\n"))
}

func ParseGlobalBlock(global GlobalBlock) (data []string) {
	data = append(data, "global")

	data = ConfigBool("daemon", global.Daemon, data)
	data = ConfigInt("maxconn %d", global.Maxconn, data)
	data = ConfigString("ca-base %s", global.CaBase, data)
	data = ConfigString("chroot %s", global.Chroot, data)
	data = ConfigString("crt-base %s", global.CrtBase, data)
	data = ConfigInt("gid %d", global.Gid, data)
	data = ConfigString("group %s", global.Group, data)

	data = ConfigStringAllowBlank("log-send-hostname %s", global.LogSendHostname, data)

	data = ConfigString("log-tag %s", global.LogTag, data)
	data = ConfigInt("nbproc %d", global.Nbproc, data)
	data = ConfigString("pidfile %s", global.Pidfile, data)
	data = ConfigString("stats bind-process %s", global.StatsBind, data)
	data = ConfigString("ssl-default-bind-ciphers %s", global.SslDefaultBind)
	// LogSendHostname

	// Nbproc

	// Pidfile

	// Uid

	// UlimitN

	// User

	// Stats

	// SslServerVerify

	// Node

	// Description

	// UnixBind

	return data
}

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

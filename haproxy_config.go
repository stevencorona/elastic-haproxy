package main

type GlobalBlock struct {
	Daemon          bool
	Maxconn         int
	CaBase          string
	Chroot          string
	CrtBase         string
	Gid             int
	Group           string
	LogSendHostname string
	Nbproc          int
	Pidfile         string
	Uid             int
	UlimitN         int
	User            string
	Stats           string
	SslServerVerify bool
	Node            string
	Description     string
	UnixBind        string

	Log struct {
		Address  string
		Length   int
		Facility string
		MaxLevel string
		MinLevel string
	}
}

type DefaultsBlock struct {
	Mode    string
	Timeout TimeoutBlock
}

type TimeoutBlock struct {
	Connect int
	Client  int
	Server  int
}

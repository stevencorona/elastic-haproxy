package haproxy

type Config struct {
	Global   GlobalBlock
	Defaults DefaultBlock
	//Frontends []FrontendBlock
}

type GlobalBlock struct {
	Daemon          bool `config:"daemon %s"`
	Maxconn         int  `config:"maxconn %s"`
	CaBase          string
	Chroot          string
	CrtBase         string
	Gid             int
	Group           string
	LogSendHostname string
	LogTag string
	CpuMap string
	Nbproc          int
	Pidfile         string
	Uid             int
	UlimitN         int
	User            string
	StatsBind       string
	StatsSocket     string
	SslDefaultBindCiphers string
	SslDefaultServerCiphers string
	SslServerVerify string
	Node            string
	Description     string
	UnixBind        string

	Log struct {
		Address  string
		Length   int
		Facility string
		MaxLevel string
		MinLevel string
	} `config:"log %s %s %s %s %s"`
}

type DefaultBlock struct {
	Mode    string
	Timeout []TimeoutBlock
	Options map[string]string
}

type TimeoutBlock struct {
	Type   string
	Amount string
}

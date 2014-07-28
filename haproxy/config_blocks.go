package haproxy

type Config struct {
	Global   GlobalBlock
	Defaults DefaultBlock
	//Frontends []FrontendBlock
}

type GlobalBlock struct {
	Daemon                  bool
	CaBase                  string
	Chroot                  string
	CrtBase                 string
	Gid                     int
	Group                   string
	Log                     string
	LogSendHostname         string
	LogTag                  string
	CpuMap                  string
	Nbproc                  int
	Pidfile                 string
	Uid                     int
	UlimitN                 int
	User                    string
	StatsBind               string
	StatsSocket             string
	SslDefaultBindCiphers   string
	SslDefaultServerCiphers string
	SslServerVerify         string
	Node                    string
	Description             string
	UnixBind                string
	MaxSpreadChecks         int
	MaxConn                 int
	MaxConnRate             int
	MaxCompRate             int
	MaxCompCpuUsage         int
	MaxPipes                int
	MaxSessRate             int
	MaxSslConn              int
	MaxSslRate              int
	MaxZlibMem              int
	NoEpoll                 bool
	NoKQueue                bool
	NoPoll                  bool
	NoSplice                bool
	NoGetAddrInfo           bool
	SpreadChecks            int
	Tuning                  struct {
		Bufsize              int
		Chksize              int
		CompMaxLevel         int
		HttpCookieLen        int
		HttpMaxHeader        int
		IdleTimer            int
		MaxAccept            int
		MaxPollEvents        int
		MaxRewrite           int
		PipeSize             int
		RcvBufClient         int
		RcvBufServer         int
		SndBufClient         int
		SndBufServer         int
		SslCacheSize         int
		SslForcePrivateCache int
		SslLifetime          int
		SslMaxRecord         int
		SslDefaultDhParam    int
		ZlibMemLevel         int
		ZlibWindowSize       int
	}
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

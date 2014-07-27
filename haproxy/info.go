package haproxy

import (
	"bufio"
	"net"
	"reflect"
	"strconv"
	"strings"
)

type HaproxyServerInfo struct {
	Name                       string `haproxy:"Name""`
	Version                    string `haproxy:"Version"`
	ReleaseDate                string `haproxy:"Release_date"`
	Nbproc                     int    `haproxy:"Nbproc"`
	ProcessNum                 int    `haproxy:"Process_num"`
	Pid                        int    `haproxy:"Pid"`
	Uptime                     string `haproxy:"Uptime"`
	UptimeSec                  int    `haproxy:"Uptime_sec"`
	MemmaxMB                   int    `haproxy:"Memmax_MB"`
	UlimitN                    int    `haproxy:"Ulimit-n"`
	Maxsock                    int    `haproxy:"Maxsock"`
	Maxconn                    int    `haproxy:"Maxconn"`
	HardMaxconn                int    `haproxy:"Hard_maxconn"`
	CurrConns                  int    `haproxy:"CurrConns"`
	CumConns                   int    `haproxy:"CumConns"`
	CumReq                     int    `haproxy:"CumReq"`
	MaxSslConns                int    `haproxy:"MaxSslConns"`
	CurrSslConns               int    `haproxy:"CurrSslConns"`
	CumSslConns                int    `haproxy:"CumSslConns"`
	Maxpipes                   int    `haproxy:"Maxpipes"`
	PipesUsed                  int    `haproxy:"PipesUsed"`
	PipesFree                  int    `haproxy:"PipesFree"`
	ConnRate                   int    `haproxy:"ConnRate"`
	ConnRateLimit              int    `haproxy:"ConnRateLimit"`
	MaxConnRate                int    `haproxy:"MaxConnRate"`
	SessRate                   int    `haproxy:"SessRate"`
	SessRateLimit              int    `haproxy:"SessRateLimit"`
	MaxSessRate                int    `haproxy:"MaxSslRate"`
	SslRate                    int    `haproxy:"SslRate"`
	SslRateLimit               int    `haproxy:"SslRateLimit"`
	MaxSslRate                 int    `haproxy:"MaxSslRate"`
	SslFrontendKeyRate         int    `haproxy:"SslFrontendKeyRate"`
	SslFrontendMaxKeyRate      int    `haproxy:"SslFrontendMaxKeyRate"`
	SslFrontendSessionReusePct int    `haproxy:"SslFrontendSessionReuse_pct"`
	SslBackendKeyRate          int    `haproxy:"SslBackendKeyRate"`
	SslBackendMaxKeyRate       int    `haproxy:"SslBackendMaxKeyRate"`
	SslCacheLookups            int    `haproxy:"SslCacheLookups"`
	SslCacheMisses             int    `haproxy:"SslCacheMisses"`
	CompressBpsIn              int    `haproxy:"CompressBpsIn"`
	CompressBpsOut             int    `haproxy:"CompressBpsOut"`
	CompressBpsRateLim         int    `haproxy:"CompressBpsRateLim"`
	ZlibMemUsage               int    `haproxy:"ZlibMemUsage"`
	MaxZlibMemUSage            int    `haproxy:"MaxZlibMemUsage"`
	Tasks                      int    `haproxy:"Tasks"`
	RunQueue                   int    `haproxy:"Run_queue"`
	IdlePct                    int    `haproxy:"Idle_pct"`
	Node                       string `haproxy:"node"`
	Description                string `haproxy:"description"`
}

func (h *Server) socketCommand(command string) (data []string, err error) {
	conn, err := net.Dial("unix", h.Socket)

	if err != nil {
		return nil, err
	}

	defer conn.Close()

	// Send command to HAProxy
	conn.Write([]byte(command))

	// Read the response line by line and return it
	reader := bufio.NewReader(conn)

	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			break
		}

		data = append(data, line)
	}

	return data, nil
}

// Make a call to the HAProxy unix socket and read it into our
// struct
func (h *Server) GetInfo() (serverInfo *HaproxyServerInfo) {
	data, err := h.socketCommand("show info\n")
	serverInfo = new(HaproxyServerInfo)

	if err != nil {
		return serverInfo
	}

	strukt := reflect.ValueOf(serverInfo).Elem()

	for _, line := range data {
		parts := strings.Split(line, ":")

		if len(parts) != 2 {
			continue
		}

		k := strings.TrimSpace(parts[0])
		v := strings.TrimSpace(parts[1])

		// Skip blank/empty keys
		if k == "" {
			continue
		}

		// Loop over all of the fields on the struct
		for i := 0; i < strukt.NumField(); i++ {

			// Grab the field value
			fieldValue := strukt.Field(i)

			// Grab the field tag
			fieldTag := strukt.Type().Field(i).Tag

			// If the haproxy tag doesn't match the key that we read in, move on.
			if k != fieldTag.Get("haproxy") {
				continue
			}

			// Set and convert the value read in depending on the type
			switch fieldValue.Kind() {
			case reflect.Int:
				intVal, _ := strconv.Atoi(v)
				fieldValue.SetInt(int64(intVal))
			case reflect.String:
				fieldValue.SetString(v)
			}
		}
	}

	return serverInfo
}

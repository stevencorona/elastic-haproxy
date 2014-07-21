package main

import (
	"bufio"
	"fmt"
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

func (h *Haproxy) GetInfo() {
	conn, err := net.Dial("unix", h.Socket)

	if err != nil {
		fmt.Println(err)
	}

	conn.Write([]byte("show info\n"))

	reader := bufio.NewReader(conn)

	h.ServerInfo = new(HaproxyServerInfo)

	val := reflect.ValueOf(h.ServerInfo).Elem()

	for {
		status, err := reader.ReadString('\n')

		if err != nil {
			//fmt.Println(err)
			break
		}

		parts := strings.Split(status, ":")

		key := parts[0]
		value := ""

		if strings.TrimSpace(key) == "" {
			continue
		}

		if len(parts) == 2 {
			value = strings.TrimSpace(parts[1])
		}

		for i := 0; i < val.NumField(); i++ {
			valueField := val.Field(i)
			typeField := val.Type().Field(i)
			tag := typeField.Tag

			if tag.Get("haproxy") == key {

				if valueField.Kind() == reflect.String {
					valueField.SetString(value)
				}

				if valueField.Kind() == reflect.Int {
					i, _ := strconv.Atoi(value)
					valueField.SetInt(int64(i))
				}
			}
		}
	}
}

package main

type conf struct {
	Lang       string `json:"lang"`
	Prefix     string `json:"prefix"`
	EsPrefix   string `json:"es_prefix"`
	PullPrefix string `json:"pull_prefix"`
	Db         struct {
		Doris struct {
			Addr        string `json:"addr"`
			MaxIdleConn int    `json:"max_idle_conn"`
			MaxOpenConn int    `json:"max_open_conn"`
		} `json:"doris"`
		PullDoris struct {
			Addr        string `json:"addr"`
			MaxIdleConn int    `json:"max_idle_conn"`
			MaxOpenConn int    `json:"max_open_conn"`
		} `json:"pull_doris"`
		Report struct {
			Addr        string `json:"addr"`
			MaxIdleConn int    `json:"max_idle_conn"`
			MaxOpenConn int    `json:"max_open_conn"`
		} `json:"report"`
		Slave struct {
			Addr        string `json:"addr"`
			MaxIdleConn int    `json:"max_idle_conn"`
			MaxOpenConn int    `json:"max_open_conn"`
		} `json:"slave"`
	} `json:"db"`
	Redis struct {
		Master struct {
			Addr     []string `json:"addr"`
			Password string   `json:"password"`
			Sentinel string   `json:"sentinel"`
			Db       int      `json:"db"`
		} `json:"master"`
	} `json:"redis"`
	ReportEs struct {
		Host     []string `json:"host"`
		Username string   `json:"username"`
		Password string   `json:"password"`
	} `json:"report_es"`
	Zlog struct {
		Host string `json:"host"`
		Port int    `json:"port"`
	} `json:"zlog"`
	Port string `json:"port"`
}

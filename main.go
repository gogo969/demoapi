package main

import (
	"fmt"
	"log"
	"os"
	"reportapi/contrib/apollo"
	"reportapi/contrib/conn"
	"reportapi/contrib/session"
	"reportapi/middleware"
	"reportapi/model"
	"reportapi/router"
	"strings"

	"github.com/valyala/fasthttp"
	_ "go.uber.org/automaxprocs"
)

var (
	gitReversion   = ""
	buildTime      = ""
	buildGoVersion = ""
)

func main() {

	argc := len(os.Args)
	if argc != 3 {
		fmt.Printf("%s <etcds> <cfgPath>\r\n", os.Args[0])
		return
	}

	cfg := conf{}
	endpoints := strings.Split(os.Args[1], ",")

	apollo.New(endpoints)
	apollo.Parse(os.Args[2], &cfg)
	apollo.Close()

	mt := new(model.MetaTable)

	mt.Zlog = conn.InitFluentd(cfg.Zlog.Host, cfg.Zlog.Port)
	mt.ReportDB = conn.InitDB(cfg.Db.Report.Addr, cfg.Db.Report.MaxIdleConn, cfg.Db.Report.MaxOpenConn)
	mt.SlaveDB = conn.InitDB(cfg.Db.Slave.Addr, cfg.Db.Slave.MaxIdleConn, cfg.Db.Slave.MaxOpenConn)
	mt.DorisDB = conn.InitDB(cfg.Db.Doris.Addr, cfg.Db.Doris.MaxIdleConn, cfg.Db.Doris.MaxOpenConn)
	mt.PullDorisDB = conn.InitDB(cfg.Db.PullDoris.Addr, cfg.Db.PullDoris.MaxIdleConn, cfg.Db.PullDoris.MaxOpenConn)
	mt.MerchantRedis = conn.InitRedisSentinel(cfg.Redis.Master.Addr, cfg.Redis.Master.Password, cfg.Redis.Master.Sentinel, cfg.Redis.Master.Db)
	//mt.ReportEs = conn.InitES(cfg.ReportEs.Host, cfg.ReportEs.Username, cfg.ReportEs.Password)
	mt.Prefix = cfg.Prefix
	mt.Lang = cfg.Lang
	mt.EsPrefix = cfg.EsPrefix
	mt.PullPrefix = cfg.PullPrefix

	model.Constructor(mt)
	session.New(mt.MerchantRedis, mt.Prefix)

	defer func() {
		model.Close()
		mt = nil
	}()

	b := router.BuildInfo{
		GitReversion:   gitReversion,
		BuildTime:      buildTime,
		BuildGoVersion: buildGoVersion,
	}
	app := router.SetupRouter(b)
	srv := &fasthttp.Server{
		Handler:            middleware.Use(app.Handler),
		ReadTimeout:        router.ApiTimeout,
		WriteTimeout:       router.ApiTimeout,
		Name:               "reportApi",
		MaxRequestBodySize: 51 * 1024 * 1024,
	}
	fmt.Printf("gitReversion = %s\r\nbuildGoVersion = %s\r\nbuildTime = %s\r\n", gitReversion, buildGoVersion, buildTime)
	fmt.Println("report2 running", cfg.Port)
	if err := srv.ListenAndServe(cfg.Port); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}
}

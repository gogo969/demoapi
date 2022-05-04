package middleware

import (
	"errors"
	"fmt"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fastjson"
	"reportapi/contrib/session"
	"reportapi/model"
)

var allows = map[string]bool{
	"/merchant/report/version":            true,
	"/merchant/report/pprof/":             true,
	"/merchant/report/pprof/block":        true,
	"/merchant/report/pprof/allocs":       true,
	"/merchant/report/pprof/cmdline":      true,
	"/merchant/report/pprof/goroutine":    true,
	"/merchant/report/pprof/heap":         true,
	"/merchant/report/pprof/profile":      true,
	"/merchant/report/pprof/trace":        true,
	"/merchant/report/pprof/threadcreate": true,
}

func CheckTokenMiddleware(ctx *fasthttp.RequestCtx) error {

	path := string(ctx.Path())
	if _, ok := allows[path]; ok {
		return nil
	}

	data, err := session.Get(ctx)
	if err != nil {
		return errors.New(`{"status":false,"data":"token"}`)
	}

	gid := fastjson.GetString(data, "group_id")
	permission := model.PrivCheck(path, gid)
	if permission != nil {
		fmt.Println("path = ", path)
		fmt.Println("gid = ", gid)
		fmt.Println("permission = ", permission)
		return errors.New(`{"status":false,"data":"permission denied"}`)
	}

	ctx.SetUserValue("token", data)

	return nil
}

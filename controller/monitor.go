package controller

import (
	"github.com/valyala/fasthttp"
	"reportApi2/contrib/helper"
	"reportApi2/model"
)

type MonitorController struct{}

// Access 报表中心-实时存取监控
func (that MonitorController) List(ctx *fasthttp.RequestCtx) {

	day := string(ctx.PostArgs().Peek("date_day"))
	page := ctx.PostArgs().GetUintOrZero("page")
	pageSize := ctx.PostArgs().GetUintOrZero("page_size")

	if page < 1 {
		helper.Print(ctx, false, helper.ParamErr)
		return
	}

	if pageSize < 10 || pageSize > 200 {
		helper.Print(ctx, false, helper.ParamErr)
		return
	}
	if len(day) >= 10 {
		day = day[:10] + " 00:00:00"
	}

	data, err := model.MonitorReport(page, pageSize, day)
	if err != nil {
		helper.Print(ctx, false, err.Error())
		return
	}

	helper.Print(ctx, true, data)
}

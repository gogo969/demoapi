package controller

import (
	g "github.com/doug-martin/goqu/v9"
	"github.com/valyala/fasthttp"
	"reportapi/contrib/helper"
	"reportapi/model"
)

type DividendController struct{}

// List 红利报表
func (that DividendController) List(ctx *fasthttp.RequestCtx) {

	flag := ctx.QueryArgs().GetUintOrZero("flag")
	ty := ctx.QueryArgs().GetUintOrZero("ty") // 211 平台红利、216 维护补偿、218 活动红利、220 红利调整、221 负数清零 222 代理红利
	page := ctx.QueryArgs().GetUintOrZero("page")
	pageSize := ctx.QueryArgs().GetUintOrZero("page_size")
	startTime := string(ctx.QueryArgs().Peek("start_time")) //开始时间
	endTime := string(ctx.QueryArgs().Peek("end_time"))     //结束时间

	if page == 0 {
		page = 1
	}

	if pageSize == 0 {
		pageSize = 15
	}

	if flag < 1 || flag > 2 {
		helper.Print(ctx, false, helper.ParamErr)
		return
	}

	ex := g.Ex{
		"flag": flag,
	}

	if ty > 0 {
		dividendMap := map[int]bool{
			211: true, // 211 平台红利
			216: true, // 216 维护补偿
			218: true, // 218 活动红利
			220: true, // 220 红利调整
			221: true, // 221 负数清零
			222: true, // 222 代理红利
		}
		if _, ok := dividendMap[ty]; !ok {
			helper.Print(ctx, false, helper.DividendTypeErr)
			return
		}
		ex["report_type"] = ty
	}

	data, err := model.DividendList(ex, startTime, endTime, page, pageSize)
	if err != nil {
		helper.Print(ctx, false, err.Error())
		return
	}

	helper.Print(ctx, true, data)
}

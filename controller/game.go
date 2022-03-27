package controller

import (
	"github.com/valyala/fasthttp"
	"reportApi2/contrib/helper"
	"reportApi2/contrib/validator"
	"reportApi2/model"
)

type GameController struct{}

// 游戏报表
func (that *GameController) List(ctx *fasthttp.RequestCtx) {

	ty := ctx.PostArgs().GetUintOrZero("ty")               //1-游戏报表,2-游戏报表场馆
	flag := ctx.PostArgs().GetUintOrZero("flag")           //1-日报;2-月报
	dateFlag := ctx.PostArgs().GetUintOrZero("date_flag")  //1-投注时间;2-结算时间
	timeFlag := ctx.PostArgs().GetUintOrZero("time_flag")  //1-单天|单月,2-按时间段
	ids := string(ctx.PostArgs().Peek("ids"))              //游戏id,多个用逗号分隔
	startDate := string(ctx.PostArgs().Peek("start_date")) //开始时间
	endDate := string(ctx.PostArgs().Peek("end_date"))     //结束时间
	page := ctx.PostArgs().GetUintOrZero("page")           //页码
	pageSize := ctx.PostArgs().GetUintOrZero("page_size")  //一页多少条
	if !validator.CheckStringCommaDigit(ids) {
		helper.Print(ctx, false, helper.ParamErr)
		return
	}
	if page < 1 {
		helper.Print(ctx, false, model.ParamErr)
		return
	}

	if pageSize < 10 || pageSize > 200 {
		helper.Print(ctx, false, model.ParamErr)
		return
	}

	data, err := model.GameReport(ty, flag, dateFlag, timeFlag, startDate, endDate, ids, page, pageSize)
	if err != nil {
		helper.Print(ctx, false, err.Error())
		return
	}

	helper.Print(ctx, true, data)
}

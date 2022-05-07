package controller

import (
	"github.com/valyala/fasthttp"
	"reportApi2/contrib/helper"
	"reportApi2/model"
)

type PlatformController struct{}

// List 报表中心-平台报表
func (that PlatformController) List(ctx *fasthttp.RequestCtx) {

	flag := ctx.PostArgs().GetUintOrZero("flag")          //1-日报;2-月报
	dateFlag := ctx.PostArgs().GetUintOrZero("date_flag") //1-投注时间;2-结算时间
	timeFlag := ctx.PostArgs().GetUintOrZero("time_flag") //1-单天|单月,2-按时间段
	page := ctx.PostArgs().GetUintOrZero("page")
	pageSize := ctx.PostArgs().GetUintOrZero("page_size")
	depositStart := string(ctx.PostArgs().Peek("depositStart"))             //存款金额最小
	depositEnd := string(ctx.PostArgs().Peek("depositEnd"))                 //存款金额最大
	betAmountStart := string(ctx.PostArgs().Peek("bet_amount_start"))       //投注金额最小
	betAmountEnd := string(ctx.PostArgs().Peek("bet_amount_end"))           //投注金额最大
	depositCountStart := string(ctx.PostArgs().Peek("deposit_count_start")) //充值笔数最小
	depositCountEnd := string(ctx.PostArgs().Peek("deposit_count_end"))     //充值笔数最大
	netAmountStart := string(ctx.PostArgs().Peek("net_amount_start"))       //净输赢最小
	netAmountEnd := string(ctx.PostArgs().Peek("net_amount_end"))           //净输赢最大
	startDate := string(ctx.PostArgs().Peek("start_date"))                  //开始时间
	endDate := string(ctx.PostArgs().Peek("end_date"))                      //结束时间

	if flag < 1 || flag > 2 || dateFlag < 1 || dateFlag > 2 || timeFlag < 1 || timeFlag > 2 || page < 1 || pageSize < 10 || pageSize > 200 {
		helper.Print(ctx, false, helper.ParamErr)
		return
	}

	data, err := model.PlatformReport(page, pageSize, flag, dateFlag, timeFlag, depositStart, depositEnd, betAmountStart,
		betAmountEnd, depositCountStart, depositCountEnd, netAmountStart, netAmountEnd, startDate, endDate)
	if err != nil {
		helper.Print(ctx, false, err.Error())
		return
	}

	helper.Print(ctx, true, data)
}

// Overview 报表中心 - 综合报表
func (that PlatformController) Overview(ctx *fasthttp.RequestCtx) {

	flag := ctx.PostArgs().GetUintOrZero("flag")           //1-日报;2-月报
	startDate := string(ctx.PostArgs().Peek("start_date")) //开始时间
	endDate := string(ctx.PostArgs().Peek("end_date"))     //结束时间

	data, err := model.ComplexReport(flag, startDate, endDate)
	if err != nil {
		helper.Print(ctx, false, err.Error())
		return
	}

	helper.Print(ctx, true, data)
}

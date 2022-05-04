package controller

import (
	"reportapi/contrib/helper"
	"reportapi/contrib/validator"
	"reportapi/model"
	"strings"

	"github.com/valyala/fasthttp"
)

type BetController struct{}

// 报表中心-实时投注监控
func (that BetController) List(ctx *fasthttp.RequestCtx) {

	sGameType := string(ctx.QueryArgs().Peek("game_type")) // 游戏类型
	page := ctx.QueryArgs().GetUintOrZero("page")
	pageSize := ctx.QueryArgs().GetUintOrZero("page_size")
	startTime := string(ctx.QueryArgs().Peek("start_time")) //开始时间
	endTime := string(ctx.QueryArgs().Peek("end_time"))     //结束时间

	var gameType []string
	if sGameType != "" {
		if !validator.CheckStringCommaDigit(sGameType) {
			helper.Print(ctx, false, helper.GameTypeErr)
			return
		}

		gameType = append(gameType, strings.Split(sGameType, ",")...)
	}

	if page < 1 {
		helper.Print(ctx, false, helper.ParamErr)
		return
	}

	if pageSize < 10 || pageSize > 100 {
		helper.Print(ctx, false, helper.ParamErr)
		return
	}

	data, err := model.BetReport(gameType, startTime, endTime, page, pageSize)
	if err != nil {
		helper.Print(ctx, false, err.Error())
		return
	}

	helper.Print(ctx, true, data)
}

package controller

import (
	"github.com/valyala/fasthttp"
	"reportapi/contrib/helper"
	"reportapi/model"
)

type MemberController struct{}

// Member 报表中心-会员报表
func (that MemberController) List(ctx *fasthttp.RequestCtx) {

	flag := ctx.QueryArgs().GetUintOrZero("flag")                             //1-日报;2-月报
	sRegStartTime := string(ctx.QueryArgs().Peek("reg_start_time"))           // 注册开始时间
	sRegEndTime := string(ctx.QueryArgs().Peek("reg_end_time"))               // 注册结束时间
	sFdepositStartTime := string(ctx.QueryArgs().Peek("fdeposit_start_time")) // 首存开始时间
	sFdepositEndTime := string(ctx.QueryArgs().Peek("fdeposit_end_time"))     // 首存结束时间
	timeOutBet := ctx.QueryArgs().GetUintOrZero("time_out_bet")               // 未投注时长
	timeOutLogin := ctx.QueryArgs().GetUintOrZero("time_out_login")           // 未登录时长
	timeOutDeposit := ctx.QueryArgs().GetUintOrZero("time_out_deposit")       // 未存款时长
	parentName := string(ctx.QueryArgs().Peek("parent_name"))                 // 代理上级名称
	parentUid := string(ctx.QueryArgs().Peek("parent_uid"))                   // 上级id
	userName := string(ctx.QueryArgs().Peek("username"))                      // 会员帐号
	sStartTime := string(ctx.QueryArgs().Peek("start_date"))                  // 开始时间
	sEndTime := string(ctx.QueryArgs().Peek("end_date"))                      // 结束时间
	dateFlag := ctx.QueryArgs().GetUintOrZero("date_flag")                    // 1-投注时间;2-结算时间
	timeFlag := ctx.QueryArgs().GetUintOrZero("time_flag")                    // 1-单天;2-时间段
	//sCommissionId := string(ctx.QueryArgs().Peek("commission_id"))            // 返佣方案
	sMainId := string(ctx.QueryArgs().Peek("main_id"))     // 维护人
	page := ctx.QueryArgs().GetUintOrZero("page")          //页码
	pageSize := ctx.QueryArgs().GetUintOrZero("page_size") //一页多少条
	ty := ctx.QueryArgs().GetUintOrZero("ty")              //1会员报表2代理报表

	if page < 1 {
		helper.Print(ctx, false, model.ParamErr)
		return
	}

	if pageSize < 10 || pageSize > 200 {
		helper.Print(ctx, false, model.ParamErr)
		return
	}

	data, err := model.MemberReport(flag, dateFlag, timeFlag, page, pageSize, timeOutBet, timeOutLogin, timeOutDeposit, sRegStartTime,
		sRegEndTime, sFdepositStartTime, sFdepositEndTime, parentName, parentUid, userName, sStartTime, sEndTime, sMainId, ty)
	if err != nil {
		helper.Print(ctx, false, err.Error())
		return
	}

	helper.Print(ctx, true, data)
}

package model

import (
	"errors"
	"fmt"
	g "github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
	"github.com/shopspring/decimal"
	"reportapi/contrib/helper"
	"strings"
	"time"
)

// 投注时间/结算时间 数据结构
type Report struct {
	ReportTime          string `db:"report_time" json:"report_time"`
	ID                  string `db:"id" json:"id"`
	ApiType             string `db:"api_type" json:"api_type"`
	MemCount            int    `db:"mem_count" json:"mem_count"`
	BetCount            int    `db:"bet_count" json:"bet_count"`
	BetAmount           string `db:"bet_amount" json:"bet_amount"`
	RebateAmount        string `db:"rebate_amount" json:"rebate_amount"`
	ValidBetAmount      string `db:"valid_bet_amount" json:"valid_bet_amount"`
	CompanyNetAmount    string `db:"company_net_amount" json:"company_net_amount"`
	AvgBetAmount        string `db:"avg_bet_amount" json:"avg_bet_amount"`
	AvgValidBetAmount   string `db:"avg_valid_bet_amount" json:"avg_valid_bet_amount"`
	AvgCompanyNetAmount string `db:"avg_company_net_amount" json:"avg_company_net_amount"`
	Presettle           string `db:"presettle" json:"presettle"`
	ProfitRate          string `db:"profit_rate" json:"profit_rate"`
}

type ReportGameUserVo struct {
	ID               string `db:"id" json:"id"`
	ApiType          string `db:"api_type" json:"api_type"`
	Username         string `db:"username" json:"username"`
	BetCount         int    `db:"bet_count" json:"bet_count"`
	BetAmount        string `db:"bet_amount" json:"bet_amount"`
	RebateAmount     string `db:"rebate_amount" json:"rebate_amount"`
	ValidBetAmount   string `db:"valid_bet_amount" json:"valid_bet_amount"`
	CompanyNetAmount string `db:"company_net_amount" json:"company_net_amount"`
	Presettle        string `db:"presettle" json:"presettle"`
	ProfitRate       string `db:"profit_rate" json:"profit_rate"`
}

type GameReportUserData struct {
	D []ReportGameUserVo `json:"d"`
	T int64              `json:"t"`
	S int                `json:"s"`
}

type GameDetailReportData struct {
	D []DetailReport `json:"d"`
	T int64          `json:"t"`
	S int            `json:"s"`
}

type DetailReport struct {
	ApiType          string  `json:"api_type" db:"api_type"`
	ApiName          string  `json:"api_name" db:"api_name"`
	GameCode         string  `json:"game_code" db:"game_code"`
	GameName         string  `json:"game_name" db:"game_name"`
	GameVnName       string  `json:"game_vn_name" db:"game_vn_name"`
	MemCount         int64   `json:"mem_count" db:"mem_count"`
	BetAmount        float64 `json:"bet_amount" db:"bet_amount"`
	ValidBetAmount   float64 `json:"valid_bet_amount" db:"valid_bet_amount"`
	CompanyNetAmount float64 `json:"company_net_amount" db:"company_net_amount"`
	RebateAmount     float64 `json:"rebate_amount" db:"rebate_amount"`
	ProfitAmount     float64 `json:"profit_amount" db:"profit_amount"`
	ProfitRate       string  `json:"profit_rate" db:"profit_rate"`
}

type PlanIssues struct {
	Id             string  `json:"id" db:"id"`
	PlanId         string  `json:"plan_id" db:"plan_id"`
	Issue          string  `json:"issue" db:"issue"`
	Content        string  `json:"content" db:"content"`
	BetMemCount    int     `json:"bet_mem_count" db:"bet_mem_count"`
	BetAmountTotal float64 `json:"bet_amount_total" db:"bet_amount_total"`
	BonusTotal     float64 `json:"bonus_total" db:"bonus_total"`
}

type GamePlanBase struct {
	PlanId         string  `json:"plan_id" db:"plan_id"`
	BetMemCount    int     `json:"bet_mem_count" db:"bet_mem_count"`
	BetAmountTotal float64 `json:"bet_amount_total" db:"bet_amount_total"`
	BonusTotal     float64 `json:"bonus_total" db:"bonus_total"`
}

type PlanReportData struct {
	D []PlanIssues `json:"d"`
	T int64        `json:"t"`
	S int          `json:"s"`
}

type GameReportData struct {
	D   []Report `json:"d"`
	T   int64    `json:"t"`
	S   int      `json:"s"`
	Agg Report   `json:"agg"`
}

// GameReport Game 游戏报表
func GameReport(ty, flag, dateFlag, timeFlag int, startTime, endTime, gameIds string, page, pagesize int) (GameReportData, error) {

	var result GameReportData

	startAt := helper.DaySST(startTime, loc).Unix()

	endAt := helper.DaySET(endTime, loc).Unix()

	if startAt > endAt {
		return result, errors.New(helper.QueryTimeRangeErr)
	}
	// 游戏报表（投注时间统计）
	if dateFlag == ReportDateFlagBet && ty == ReportTyGame {
		result, err := gameReportBetTime(startAt, endAt, flag, timeFlag, gameIds, page, pagesize)
		return result, err
	}
	// 游戏报表（结算时间统计）
	if dateFlag == ReportDateFlagSettle && ty == ReportTyGame {
		result, err := gameReportSettleTime(startAt, endAt, flag, timeFlag, gameIds, page, pagesize)
		return result, err
	}
	// 场馆游戏报表（结算时间统计）
	if dateFlag == ReportDateFlagSettle && ty == ReportTyPlat {
		result, err := gamePlatReportSettleTime(startAt, endAt, flag, timeFlag, gameIds, page, pagesize)
		return result, err
	}

	return result, nil
}

// GameReportUser 游戏报表
func GameReportUser(ty, flag, dateFlag, timeFlag int, startTime, endTime, apiType string, page, pageSize int) (GameReportUserData, error) {

	var data GameReportUserData

	startAt := helper.DaySST(startTime, loc).Unix()

	endAt := helper.DaySET(endTime, loc).Unix()

	if startAt > endAt {
		return data, errors.New(helper.QueryTimeRangeErr)
	}
	ex := g.Ex{
		"report_time": g.Op{"between": exp.NewRangeVal(startAt, endAt)},
		"api_type":    apiType,
		"prefix":      meta.Prefix,
	}
	tableName := "tbl_report_game_user"
	if ty == ReportTyGame && dateFlag == ReportDateFlagBet && flag == ReportFlagDay { // 游戏投注日报
		ex["report_type"] = 1
	}

	if ty == ReportTyGame && dateFlag == ReportDateFlagSettle && flag == ReportFlagDay { // 游戏结算日报
		ex["report_type"] = 2
	}

	if ty == ReportTyGame && dateFlag == ReportDateFlagBet && flag == ReportFlagMonth { // 游戏投注日报
		ex["report_type"] = 3
	}

	if ty == ReportTyGame && dateFlag == ReportDateFlagSettle && flag == ReportFlagMonth { // 游戏结算日报
		ex["report_type"] = 4
	}

	if ty == ReportTyPlat && flag == ReportFlagDay { // 场馆结算日报
		ex["report_type"] = 5
	}

	if ty == ReportTyPlat && flag == ReportFlagMonth { // 场馆结算月报
		ex["report_type"] = 6
	}
	offset := (page - 1) * pageSize
	build := dialect.From(tableName).Where(ex)

	build = build.GroupBy("username").Select(
		"username",
		g.SUM("bet_count").As("bet_count"),
		g.SUM("bet_amount").As("bet_amount"),
		g.SUM("valid_bet_amount").As("valid_bet_amount"),
		g.SUM("company_net_amount").As("company_net_amount"),
		g.SUM("presettle").As("presettle"),
		g.SUM("profit_rate").As("profit_rate"),
	).Order(g.C("report_time").Desc()).Offset(uint(offset)).Limit(uint(pageSize))

	buildCount := dialect.From(tableName).Select(g.COUNT(g.DISTINCT("username"))).Where(ex)
	query, _, _ := buildCount.ToSQL()
	fmt.Println(query)
	err := meta.ReportDB.Get(&data.T, query)
	if err != nil {
		return data, pushLog(fmt.Errorf("%s,[%s]", err.Error(), query), "db", helper.DBErr)
	}

	query, _, _ = build.ToSQL()
	fmt.Println(query)
	err = meta.ReportDB.Select(&data.D, query)
	if err != nil {
		return data, pushLog(fmt.Errorf("%s,[%s]", err.Error(), query), "db", helper.DBErr)
	}

	data.D = reportGameUserFormat(data.D, apiType)

	return data, nil
}

// GameDetailReport 游戏子游戏报表
func GameDetailReport(flag int, startTime, endTime string, gameIds []string, page, pageSize int) (GameDetailReportData, error) {

	var result GameDetailReportData

	startAt := helper.DaySST(startTime, loc).Unix()

	endAt := helper.DaySET(endTime, loc).Unix()

	if startAt > endAt {
		return result, errors.New(helper.QueryTimeRangeErr)
	}
	ex := g.Ex{
		"report_time": g.Op{"between": exp.NewRangeVal(startAt, endAt)},
		"api_type":    gameIds,
		"prefix":      meta.Prefix,
	}
	tableName := "tbl_report_game_detail"
	if flag == ReportFlagDay { // 日报
		ex["report_type"] = 1
	}

	if flag == ReportFlagMonth { //
		ex["report_type"] = 2
	}
	offset := (page - 1) * pageSize

	buildCount := dialect.From(tableName).Select(g.COUNT(g.DISTINCT("game_code"))).Where(ex)
	query, _, _ := buildCount.ToSQL()
	err := meta.ReportDB.Get(&result.T, query)
	if err != nil {
		return result, pushLog(fmt.Errorf("%s,[%s]", err.Error(), query), "db", helper.DBErr)
	}

	build := dialect.From(tableName).Where(ex)

	build = build.GroupBy("api_type", "game_code").Select(
		"api_type", "game_code",
		g.SUM("mem_count").As("mem_count"),
		g.SUM("bet_amount").As("bet_amount"),
		g.SUM("valid_bet_amount").As("valid_bet_amount"),
		g.SUM("company_net_amount").As("company_net_amount"),
		g.SUM("rebate_amount").As("rebate_amount"),
		g.SUM("profit_amount").As("profit_amount"),
		g.SUM("profit_rate").As("profit_rate"),
	).Order(g.C("api_type").Desc()).Offset(uint(offset)).Limit(uint(pageSize))
	query, _, _ = build.ToSQL()
	err = meta.ReportDB.Select(&result.D, query)
	fmt.Println(query)
	if err != nil {
		return result, pushLog(fmt.Errorf("%s,[%s]", err.Error(), query), "db", helper.DBErr)
	}
	l := len(result.D)
	for i := 0; i < l; i++ {

		v := result.D[i]
		v.GameName = gameNameMap[v.ApiType][v.GameCode].CnName
		v.GameVnName = gameNameMap[v.ApiType][v.GameCode].VnName
		if v.ApiType == "8840968486572375835" {
			v.ApiName = "P3_CP"
		}
		if v.ApiType == "2326854765648775667" {
			v.ApiName = "AE_CP"
		}
		v.ProfitAmount, _ = decimal.NewFromFloat(v.CompanyNetAmount).Float64()
		if decimal.NewFromFloat(v.ValidBetAmount).Cmp(decimal.Zero) != 0 {
			v.ProfitRate = decimal.NewFromFloat(v.ProfitAmount).Div(decimal.NewFromFloat(v.ValidBetAmount)).StringFixed(4)
		}
		result.D[i] = v
	}

	return result, nil
}

// 结算时间
func gamePlatReportSettleTime(startAt, endAt int64, flag, timeFlag int, gameIds string, page, pageSize int) (GameReportData, error) {

	var data GameReportData
	ex := g.Ex{
		"report_time": g.Op{"between": exp.NewRangeVal(startAt, endAt)},
		"api_type":    strings.Split(gameIds, ","),
		"prefix":      meta.Prefix,
	}
	tableName := "tbl_report_game"
	if flag == ReportFlagDay { // 日报
		ex["report_type"] = 5
	}

	if flag == ReportFlagMonth { //
		ex["report_type"] = 6
	}
	offset := (page - 1) * pageSize
	build := dialect.From(tableName).Where(ex)

	if timeFlag == ReportTimeFlagPart {
		build = build.GroupBy("api_type").Select(
			"api_type",
			g.SUM("mem_count").As("mem_count"),
			g.SUM("bet_count").As("bet_count"),
			g.SUM("bet_amount").As("bet_amount"),
			g.SUM("valid_bet_amount").As("valid_bet_amount"),
			g.SUM("company_net_amount").As("company_net_amount"),
			g.SUM("presettle").As("presettle"),
			g.SUM("profit_rate").As("profit_rate"),
			g.SUM("avg_bet_amount").As("avg_bet_amount"),
			g.SUM("avg_valid_bet_amount").As("avg_valid_bet_amount"),
			g.SUM("avg_company_net_amount").As("avg_company_net_amount"),
		).Order(g.C("report_time").Desc()).Offset(uint(offset)).Limit(uint(pageSize))

		buildCount := dialect.From(tableName).Select(g.COUNT(g.DISTINCT("api_type"))).Where(ex)
		query, _, _ := buildCount.ToSQL()
		err := meta.ReportDB.Get(&data.T, query)
		if err != nil {
			return data, pushLog(fmt.Errorf("%s,[%s]", err.Error(), query), "db", helper.DBErr)
		}
	}

	if timeFlag == ReportTimeFlagSingle {
		tempCol := colReport[1:]
		tempCol = append(tempCol, g.C("report_time").As("report_time"))
		build = build.Select(tempCol...).Order(g.C("report_time").Desc()).Offset(uint(offset)).Limit(uint(pageSize))

		var t []total
		buildCount := dialect.From(tableName).Select(g.COUNT("api_type").As("count")).GroupBy("report_time", "api_type").Where(ex)
		query, _, _ := buildCount.ToSQL()
		err := meta.ReportDB.Select(&t, query)
		if err != nil {
			return data, pushLog(fmt.Errorf("%s,[%s]", err.Error(), query), "db", helper.DBErr)
		}
		data.T = int64(len(t))
		if data.T == 0 {
			return data, nil
		}
	}

	query, _, _ := build.ToSQL()
	err := meta.ReportDB.Select(&data.D, query)
	if err != nil {
		return data, pushLog(fmt.Errorf("%s,[%s]", err.Error(), query), "db", helper.DBErr)
	}

	aggQuery, _, _ := dialect.From(tableName).Select(g.SUM("mem_count").As("mem_count"),
		g.SUM("bet_count").As("bet_count"),
		g.SUM("bet_amount").As("bet_amount"),
		g.SUM("rebate_amount").As("rebate_amount"),
		g.SUM("valid_bet_amount").As("valid_bet_amount"),
		g.SUM("company_net_amount").As("company_net_amount"),
		g.SUM("presettle").As("presettle"),
		g.SUM("profit_rate").As("profit_rate"),
		g.SUM("avg_bet_amount").As("avg_bet_amount"),
		g.SUM("avg_valid_bet_amount").As("avg_valid_bet_amount"),
		g.SUM("avg_company_net_amount").As("avg_company_net_amount")).Where(ex).ToSQL()
	err = meta.ReportDB.Get(&data.Agg, aggQuery)
	if err != nil {
		return data, pushLog(fmt.Errorf("%s,[%s]", err.Error(), query), "db", helper.DBErr)
	}

	data.D = reportGameFormat(data.D)
	if timeFlag == ReportTimeFlagPart || flag == ReportFlagMonth {
		for i := 0; i < len(data.D); i++ {
			if flag == ReportFlagMonth && timeFlag == ReportTimeFlagSingle {
				data.D[i].ReportTime = parsePart(startAt, 0, "m")
			} else if flag == ReportFlagMonth && timeFlag == ReportTimeFlagPart {
				data.D[i].ReportTime = parsePart(startAt, endAt, "m")
			} else {
				data.D[i].ReportTime = parsePart(startAt, endAt, "d")
			}
		}
	}
	return data, nil
}

/**
 * @Description: gameReportBetTime 投注时间统计游戏报表
 * @Author: parker
 * @Date: 2021/4/13 12:05
 * @LastEditTime: 2021/4/13 21:00
 * @LastEditors: parker
 */
func gameReportBetTime(startAt, endAt int64, flag, timeFlag int, gameIds string, page, pageSize int) (GameReportData, error) {

	var data GameReportData
	ex := g.Ex{
		"report_time": g.Op{"between": exp.NewRangeVal(startAt, endAt)},
		"api_type":    strings.Split(gameIds, ","),
		"prefix":      meta.Prefix,
	}

	tableName := "tbl_report_game"
	if flag == ReportFlagDay {
		ex["report_type"] = 1
	}

	if flag == ReportFlagMonth {
		ex["report_type"] = 3
	}
	offset := (page - 1) * pageSize
	build := dialect.From(tableName).Where(ex)

	if timeFlag == ReportTimeFlagSingle {
		tempCol := colReport[1:]
		tempCol = append(tempCol, g.C("report_time").As("report_time"))
		build = build.Select(tempCol...).Order(g.C("report_time").Desc()).Offset(uint(offset)).Limit(uint(pageSize))

		var t []total
		buildCount := dialect.From(tableName).Select(g.C("api_type").As("count"), g.C("report_time").As("report_time")).Where(ex)
		query, _, _ := buildCount.ToSQL()
		err := meta.ReportDB.Select(&t, query)
		if err != nil {
			return data, pushLog(fmt.Errorf("%s,[%s]", err.Error(), query), "db", helper.DBErr)
		}
		data.T = int64(len(t))
		if data.T == 0 {
			return data, nil
		}
	}

	if timeFlag == ReportTimeFlagPart {
		build = build.Select("api_type",
			g.SUM("mem_count").As("mem_count"),
			g.SUM("bet_count").As("bet_count"),
			g.SUM("bet_amount").As("bet_amount"),
			g.SUM("rebate_amount").As("rebate_amount"),
			g.SUM("valid_bet_amount").As("valid_bet_amount"),
			g.SUM("company_net_amount").As("company_net_amount"),
			g.SUM("presettle").As("presettle"),
			g.SUM("profit_amount").As("profit_amount"),
			g.SUM("avg_bet_amount").As("avg_bet_amount"),
			g.SUM("avg_valid_bet_amount").As("avg_valid_bet_amount"),
			g.SUM("profit_rate").As("profit_rate"),
			g.SUM("avg_company_net_amount").As("avg_company_net_amount"),
		).GroupBy("api_type").Order(g.C("report_time").Desc()).Offset(uint(offset)).Limit(uint(pageSize))

		buildCount := dialect.From(tableName).Select(g.COUNT(g.DISTINCT("api_type"))).Where(ex)
		query, _, _ := buildCount.ToSQL()
		err := meta.ReportDB.Get(&data.T, query)
		if err != nil {
			return data, pushLog(fmt.Errorf("%s,[%s]", err.Error(), query), "db", helper.DBErr)
		}

	}

	query, _, _ := build.ToSQL()
	err := meta.ReportDB.Select(&data.D, query)
	if err != nil {
		return data, pushLog(fmt.Errorf("%s,[%s]", err.Error(), query), "db", helper.DBErr)
	}

	aggQuery, _, _ := dialect.From(tableName).Select(g.SUM("mem_count").As("mem_count"),
		g.SUM("bet_count").As("bet_count"),
		g.SUM("bet_amount").As("bet_amount"),
		g.SUM("rebate_amount").As("rebate_amount"),
		g.SUM("valid_bet_amount").As("valid_bet_amount"),
		g.SUM("company_net_amount").As("company_net_amount"),
		g.SUM("presettle").As("presettle"),
		g.SUM("avg_bet_amount").As("avg_bet_amount"),
		g.SUM("avg_valid_bet_amount").As("avg_valid_bet_amount"),
		g.SUM("profit_rate").As("profit_rate"),
		g.SUM("avg_company_net_amount").As("avg_company_net_amount")).Where(ex).ToSQL()
	err = meta.ReportDB.Get(&data.Agg, aggQuery)
	if err != nil {
		return data, pushLog(fmt.Errorf("%s,[%s]", err.Error(), query), "db", helper.DBErr)
	}

	data.D = reportGameFormat(data.D)
	if timeFlag == ReportTimeFlagPart || flag == ReportFlagMonth {
		for i := 0; i < len(data.D); i++ {
			if flag == ReportFlagMonth && timeFlag == ReportTimeFlagSingle {
				data.D[i].ReportTime = parsePart(startAt, 0, "m")
			} else if flag == ReportFlagMonth && timeFlag == ReportTimeFlagPart {
				data.D[i].ReportTime = parsePart(startAt, endAt, "m")
			} else {
				data.D[i].ReportTime = parsePart(startAt, endAt, "d")
			}
		}
	}
	return data, nil
}

/**
 * @Description: gameReportSettleTime 按结算时间统计游戏报表
 * @Author: parker
 * @Date: 2021/4/13 12:05
 * @LastEditTime: 2021/4/13 21:00
 * @LastEditors: parker
 */
func gameReportSettleTime(startAt, endAt int64, flag, timeFlag int, gameIds string, page, pageSize int) (GameReportData, error) {

	var data GameReportData
	ex := g.Ex{
		"report_time": g.Op{"between": exp.NewRangeVal(startAt, endAt)},
		"api_type":    strings.Split(gameIds, ","),
		"prefix":      meta.Prefix,
	}

	tableName := "tbl_report_game"
	if flag == ReportFlagDay {
		ex["report_type"] = 2
	}

	if flag == ReportFlagMonth {
		ex["report_type"] = 4
	}
	offset := (page - 1) * pageSize
	build := dialect.From(tableName).Where(ex)
	if timeFlag == ReportTimeFlagSingle {
		tempCol := colReport[1:]
		tempCol = append(tempCol, g.C("report_time").As("report_time"))
		build = build.Select(tempCol...).Order(g.C("report_time").Desc()).Offset(uint(offset)).Limit(uint(pageSize))
		var t []total
		buildCount := dialect.From(tableName).Select(g.C("api_type").As("count"), g.C("report_time").As("report_time")).Where(ex)
		query, _, _ := buildCount.ToSQL()
		err := meta.ReportDB.Select(&t, query)
		if err != nil {
			return data, pushLog(fmt.Errorf("%s,[%s]", err.Error(), query), "db", helper.DBErr)
		}
		data.T = int64(len(t))
		if data.T == 0 {
			return data, nil
		}
	}

	if timeFlag == ReportTimeFlagPart {

		build = build.Select(
			"api_type",
			g.SUM("mem_count").As("mem_count"),
			g.SUM("bet_count").As("bet_count"),
			g.SUM("bet_amount").As("bet_amount"),
			g.SUM("rebate_amount").As("rebate_amount"),
			g.SUM("valid_bet_amount").As("valid_bet_amount"),
			g.SUM("company_net_amount").As("company_net_amount"),
			g.SUM("avg_bet_amount").As("avg_bet_amount"),
			g.SUM("avg_valid_bet_amount").As("avg_valid_bet_amount"),
			g.SUM("avg_company_net_amount").As("avg_company_net_amount"),
			g.SUM("presettle").As("presettle"),
			g.SUM("profit_rate").As("profit_rate"),
		).GroupBy("api_type").Order(g.C("report_time").Desc()).Offset(uint(offset)).Limit(uint(pageSize))

		buildCount := dialect.From(tableName).Select(g.COUNT(g.DISTINCT("api_type"))).Where(ex)
		query, _, _ := buildCount.ToSQL()
		fmt.Println(query)
		err := meta.ReportDB.Get(&data.T, query)
		if err != nil {
			return data, pushLog(fmt.Errorf("%s,[%s]", err.Error(), query), "db", helper.DBErr)
		}
	}

	query, _, _ := build.ToSQL()
	err := meta.ReportDB.Select(&data.D, query)
	if err != nil {
		return data, pushLog(fmt.Errorf("%s,[%s]", err.Error(), query), "db", helper.DBErr)
	}

	if data.T > 0 {
		aggQuery, _, _ := dialect.From(tableName).Select(
			g.SUM("mem_count").As("mem_count"),
			g.SUM("bet_count").As("bet_count"),
			g.SUM("bet_amount").As("bet_amount"),
			g.SUM("rebate_amount").As("rebate_amount"),
			g.SUM("valid_bet_amount").As("valid_bet_amount"),
			g.SUM("company_net_amount").As("company_net_amount"),
			g.SUM("presettle").As("presettle"),
			g.SUM("avg_bet_amount").As("avg_bet_amount"),
			g.SUM("avg_valid_bet_amount").As("avg_valid_bet_amount"),
			g.SUM("profit_rate").As("profit_rate"),
			g.SUM("avg_company_net_amount").As("avg_company_net_amount")).Where(ex).ToSQL()
		fmt.Println(aggQuery)
		err = meta.ReportDB.Get(&data.Agg, aggQuery)
		if err != nil {
			return data, pushLog(fmt.Errorf("%s,[%s]", err.Error(), query), "db", helper.DBErr)
		}
	}

	data.D = reportGameFormat(data.D)
	if timeFlag == ReportTimeFlagPart || flag == ReportFlagMonth {
		for i := 0; i < len(data.D); i++ {
			if flag == ReportFlagMonth && timeFlag == ReportTimeFlagSingle {
				data.D[i].ReportTime = parsePart(startAt, 0, "m")
			} else if flag == ReportFlagMonth && timeFlag == ReportTimeFlagPart {
				data.D[i].ReportTime = parsePart(startAt, endAt, "m")
			} else {
				data.D[i].ReportTime = parsePart(startAt, endAt, "d")
			}
		}
	}
	return data, nil
}

func reportGameUserFormat(data []ReportGameUserVo, id string) []ReportGameUserVo {

	for k, v := range data {

		data[k].ApiType = id
		val, _ := decimal.NewFromString(v.BetAmount)
		data[k].BetAmount = val.StringFixed(4)

		ValidBetAmount, _ := decimal.NewFromString(v.ValidBetAmount)
		data[k].ValidBetAmount = ValidBetAmount.StringFixed(4)

		CompanyNetAmount, _ := decimal.NewFromString(v.CompanyNetAmount)
		data[k].CompanyNetAmount = CompanyNetAmount.StringFixed(4)

		Presettle, _ := decimal.NewFromString(v.Presettle)
		data[k].Presettle = Presettle.StringFixed(4)

		RebateAmount, _ := decimal.NewFromString(v.RebateAmount)
		data[k].RebateAmount = RebateAmount.StringFixed(4)

		data[k].ProfitRate = (CompanyNetAmount.Add(Presettle)).Div(ValidBetAmount).StringFixed(4)
	}

	return data
}

func reportGameFormat(data []Report) []Report {

	for k, v := range data {

		data[k].ReportTime = parseDay(v.ReportTime)

		val, _ := decimal.NewFromString(v.BetAmount)
		data[k].BetAmount = val.StringFixed(4)

		ValidBetAmount, _ := decimal.NewFromString(v.ValidBetAmount)
		data[k].ValidBetAmount = ValidBetAmount.StringFixed(4)

		CompanyNetAmount, _ := decimal.NewFromString(v.CompanyNetAmount)
		data[k].CompanyNetAmount = CompanyNetAmount.StringFixed(4)

		Presettle, _ := decimal.NewFromString(v.Presettle)
		data[k].Presettle = Presettle.StringFixed(4)

		RebateAmount, _ := decimal.NewFromString(v.RebateAmount)
		data[k].Presettle = RebateAmount.StringFixed(4)

		data[k].ProfitRate = CompanyNetAmount.Add(Presettle).Div(ValidBetAmount).StringFixed(4)
	}

	return data
}

// GamePlanReport 游戏计划报表
func GamePlanReport(id string, page, pageSize int) (PlanReportData, error) {

	var result PlanReportData

	ex := g.Ex{
		"plan_id": id,
		//"lott_id": lottId,
		//"play_id": playId,
		//"prefix":  meta.Prefix,
		"created_at": g.Op{"between": exp.NewRangeVal(0, time.Now().Unix())},
	}
	tableName := "tbl_vncp_plan_issues"

	offset := (page - 1) * pageSize

	buildCount := dialect.From(tableName).Select(g.COUNT("id")).Where(ex)
	query, _, _ := buildCount.ToSQL()
	err := meta.SlaveDB.Get(&result.T, query)
	fmt.Println(query)
	if err != nil {
		return result, pushLog(fmt.Errorf("%s,[%s]", err.Error(), query), "db", helper.DBErr)
	}

	build := dialect.From(tableName).Where(ex)

	build = build.Select(
		"id", "plan_id", "issue", "content",
	).Order(g.C("created_at").Desc()).Offset(uint(offset)).Limit(uint(pageSize))
	query, _, _ = build.ToSQL()
	err = meta.SlaveDB.Select(&result.D, query)
	if err != nil {
		return result, pushLog(fmt.Errorf("%s,[%s]", err.Error(), query), "db", helper.DBErr)
	}
	l := len(result.D)
	if l > 0 {
		var ids []string
		for _, v := range result.D {
			ids = append(ids, v.Id)
		}

		var gpb []GamePlanBase
		gpbMap := map[string]GamePlanBase{}
		exRp := g.Ex{
			"plan_id": ids,
			"prefix":  meta.Prefix,
		}

		build = dialect.From("tbl_report_game_plan").Where(exRp).Select(
			"plan_id",
			"bet_mem_count",
			"bet_amount_total",
			"bonus_total",
		)
		query, _, _ = build.ToSQL()
		err = meta.ReportDB.Select(&gpb, query)
		if err != nil {
			return result, pushLog(fmt.Errorf("%s,[%s]", err.Error(), query), "db", helper.DBErr)
		}

		if len(gpb) > 0 {
			for _, v := range gpb {
				gpbMap[v.PlanId] = v
			}
		}

		for i := 0; i < len(result.D); i++ {
			result.D[i].BetMemCount = gpbMap[result.D[i].Id].BetMemCount
			result.D[i].BetAmountTotal = gpbMap[result.D[i].Id].BetAmountTotal
			result.D[i].BonusTotal, _ = decimal.NewFromFloat(gpbMap[result.D[i].Id].BonusTotal).Sub(
				decimal.NewFromFloat(gpbMap[result.D[i].Id].BetAmountTotal)).Float64()
		}
	}

	return result, nil
}

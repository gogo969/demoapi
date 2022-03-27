package model

import (
	"errors"
	"fmt"
	g "github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
	"github.com/shopspring/decimal"
	"reportApi2/contrib/helper"
	"strings"
)

// 投注时间/结算时间 数据结构
type Report struct {
	ReportTime          string `db:"report_time" json:"report_time"`
	ID                  string `db:"id" json:"id"`
	ApiType             string `db:"api_type" json:"api_type"`
	MemCount            int    `db:"mem_count" json:"mem_count"`
	BetCount            int    `db:"bet_count" json:"bet_count"`
	BetAmount           string `db:"bet_amount" json:"bet_amount"`
	ValidBetAmount      string `db:"valid_bet_amount" json:"valid_bet_amount"`
	CompanyNetAmount    string `db:"company_net_amount" json:"company_net_amount"`
	AvgBetAmount        string `db:"avg_bet_amount" json:"avg_bet_amount"`
	AvgValidBetAmount   string `db:"avg_valid_bet_amount" json:"avg_valid_bet_amount"`
	AvgCompanyNetAmount string `db:"avg_company_net_amount" json:"avg_company_net_amount"`
	Presettle           string `db:"presettle" json:"presettle"`
	ProfitRate          string `db:"profit_rate" json:"profit_rate"`
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

		buildCount := dialect.From(tableName).Select(g.COUNT("api_type")).Where(ex)
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
		"bet_time": g.Op{"between": exp.NewRangeVal(startAt, endAt)},
		"api_type": strings.Split(gameIds, ","),
		"prefix":   meta.Prefix,
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
			g.SUM("valid_bet_amount").As("valid_bet_amount"),
			g.SUM("company_net_amount").As("company_net_amount"),
			g.SUM("presettle").As("presettle"),
			g.SUM("profit_amount").As("profit_amount"),
			g.SUM("avg_bet_amount").As("avg_bet_amount"),
			g.SUM("avg_valid_bet_amount").As("avg_valid_bet_amount"),
			g.SUM("profit_rate").As("profit_rate"),
			g.SUM("avg_company_net_amount").As("avg_company_net_amount"),
		).GroupBy("api_type").Order(g.C("report_time").Desc()).Offset(uint(offset)).Limit(uint(pageSize))

		buildCount := dialect.From(tableName).Select(g.COUNT("api_type")).Where(ex)
		query, _, _ := buildCount.ToSQL()
		err := meta.ReportDB.Get(&data.T, query)
		if err != nil {
			return data, pushLog(fmt.Errorf("%s,[%s]", err.Error(), query), "db", helper.DBErr)
		}

	}

	query, _, _ := build.ToSQL()
	err := meta.ReportDB.Select(&data, query)
	if err != nil {
		return data, pushLog(fmt.Errorf("%s,[%s]", err.Error(), query), "db", helper.DBErr)
	}

	aggQuery, _, _ := dialect.From(tableName).Select(g.SUM("mem_count").As("mem_count"),
		g.SUM("bet_count").As("bet_count"),
		g.SUM("bet_amount").As("bet_amount"),
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
			g.SUM("valid_bet_amount").As("valid_bet_amount"),
			g.SUM("company_net_amount").As("company_net_amount"),
			g.SUM("avg_bet_amount").As("avg_bet_amount"),
			g.SUM("avg_valid_bet_amount").As("avg_valid_bet_amount"),
			g.SUM("avg_company_net_amount").As("avg_company_net_amount"),
			g.SUM("presettle").As("presettle"),
			g.SUM("profit_rate").As("profit_rate"),
		).GroupBy("api_type").Order(g.C("report_time").Desc()).Offset(uint(offset)).Limit(uint(pageSize))

		buildCount := dialect.From(tableName).Select(g.COUNT("api_type")).Where(ex)
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
	return data, nil
}

func reportGameFormat(data []Report) []Report {

	for k, v := range data {

		data[k].ReportTime = parseDay(v.ReportTime)

		val, _ := decimal.NewFromString(v.BetAmount)
		data[k].BetAmount = val.StringFixed(4)

		val, _ = decimal.NewFromString(v.ValidBetAmount)
		data[k].ValidBetAmount = val.StringFixed(4)

		val, _ = decimal.NewFromString(v.CompanyNetAmount)
		data[k].CompanyNetAmount = val.StringFixed(4)

		val, _ = decimal.NewFromString(v.Presettle)
		data[k].Presettle = val.StringFixed(4)

		val, _ = decimal.NewFromString(v.ProfitRate)
		data[k].ProfitRate = val.StringFixed(4)
	}

	return data
}

package model

import (
	"errors"
	"fmt"
	g "github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
	"github.com/shopspring/decimal"
	"reportapi/contrib/helper"
)

//数据库查出来的
type RealTimeBetData struct {
	BetTime          string  `json:"bet_time" db:"bet_time"`
	GameType         string  `json:"game_type" db:"game_type"`
	Prefix           string  `json:"prefix" db:"prefix"`
	MemCount         int64   `json:"mem_count" db:"mem_count"`
	BetAmount        float64 `json:"bet_amount" db:"bet_amount"`
	ValidBetAmount   float64 `json:"valid_bet_amount" db:"valid_bet_amount"`
	CompanyNetAmount float64 `json:"company_net_amount" db:"company_net_amount"`
	Presettle        float64 `json:"presettle" db:"presettle"`
}

//代码经过计算再插入的
type TRealTimeReport struct {
	Id                   string  `json:"id" db:"id"`
	BetTime              int64   `json:"bet_time" db:"bet_time"`
	GameType             string  `json:"game_type" db:"game_type"`
	Prefix               string  `json:"prefix" db:"prefix"`
	MemCount             int64   `json:"mem_count" db:"mem_count"`
	BetAmount            float64 `json:"bet_amount" db:"bet_amount"`
	ValidBetAmount       float64 `json:"valid_bet_amount" db:"valid_bet_amount"`
	CompanyNetAmount     float64 `json:"company_net_amount" db:"company_net_amount"`
	Presettle            float64 `json:"presettle" db:"presettle"`
	MemCountRate         float64 `json:"mem_count_rate" db:"mem_count_rate"`
	BetAmountRate        float64 `json:"bet_amount_rate" db:"bet_amount_rate"`
	CompanyNetAmountRate float64 `json:"company_net_amount_rate" db:"company_net_amount_rate"`
}

// RealTimeReport 实时投注监控(表：t_real_time_report) table struct
type RealTimeReport struct {
	ID                   string `db:"id" json:"id"`
	Prefix               string `json:"prefix" db:"prefix"`
	BetTime              string `db:"bet_time" json:"bet_time"`                               // 下注时间
	GameType             string `db:"game_type" json:"game_type"`                             // 游戏类型
	MemCount             int64  `db:"mem_count" json:"mem_count"`                             // 时段投注人数
	MemCountRate         string `db:"mem_count_rate" json:"mem_count_rate"`                   // 时段投注人数涨幅
	BetAmount            string `db:"bet_amount" json:"bet_amount"`                           // 投注金额
	BetAmountRate        string `db:"bet_amount_rate" json:"bet_amount_rate"`                 // 投注金额涨幅
	ValidBetAmount       string `db:"valid_bet_amount" json:"valid_bet_amount"`               // 有效投注金额
	CompanyNetAmount     string `db:"company_net_amount" json:"company_net_amount"`           // 公司输赢
	CompanyNetAmountRate string `db:"company_net_amount_rate" json:"company_net_amount_rate"` // 公司输赢涨幅
	Presettle            string `db:"presettle" json:"presettle"`                             // 提前结算
}

// BetReportData 实时投注监控 response struct
type BetReportData struct {
	D []RealTimeReport `json:"d"`
	T int64            `json:"t"`
	S uint16           `json:"s"`
}

// BetReport 实时投注监控列表
func BetReport(gameType []string, startTime, endTime string, page, pageSize int) (BetReportData, error) {

	data := BetReportData{}

	startAt, err := helper.TimeToLoc(startTime, loc)
	if err != nil {
		return data, errors.New(helper.DateTimeErr)
	}

	endAt, err := helper.TimeToLoc(endTime, loc)
	if err != nil {
		return data, errors.New(helper.DateTimeErr)
	}

	if startAt >= endAt {
		return data, errors.New(helper.QueryTimeRangeErr)
	}

	ex := g.Ex{"prefix": meta.Prefix}
	ex["bet_time"] = g.Op{"between": exp.NewRangeVal(startAt, endAt)}
	if len(gameType) > 0 {
		ex["game_type"] = g.Op{"in": gameType}
	}

	if page == 1 {
		totalQuery, _, _ := dialect.From("tbl_report_bet").Select(g.COUNT(1)).Where(ex).ToSQL()
		err = meta.ReportDB.Get(&data.T, totalQuery)
		if err != nil {
			return data, pushLog(fmt.Errorf("%s,[%s]", err.Error(), totalQuery), "db", helper.DBErr)
		}

		if data.T == 0 {
			return data, nil
		}
	}

	offset := (page - 1) * pageSize
	query, _, _ := dialect.From("tbl_report_bet").Select(colRealTimeReport...).
		Where(ex).Order(g.C("bet_time").Desc()).Offset(uint(offset)).Limit(uint(pageSize)).ToSQL()

	err = meta.ReportDB.Select(&data.D, query)
	if err != nil {
		return data, pushLog(fmt.Errorf("%s,[%s]", err.Error(), query), "db", helper.DBErr)
	}

	// 涨幅保留2位小数（涨幅50.76%数据库中保存的是0.5076所以需要保留4位小数）
	for k, v := range data.D {

		val, _ := decimal.NewFromString(data.D[k].CompanyNetAmountRate)
		data.D[k].CompanyNetAmountRate = val.StringFixed(4)

		val, _ = decimal.NewFromString(data.D[k].BetAmountRate)
		data.D[k].BetAmountRate = val.StringFixed(4)

		val, _ = decimal.NewFromString(data.D[k].MemCountRate)
		data.D[k].MemCountRate = val.StringFixed(4)

		val, _ = decimal.NewFromString(data.D[k].Presettle)
		data.D[k].Presettle = val.StringFixed(4)

		val, _ = decimal.NewFromString(data.D[k].CompanyNetAmount)
		data.D[k].CompanyNetAmount = val.StringFixed(4)

		val, _ = decimal.NewFromString(data.D[k].ValidBetAmount)
		data.D[k].ValidBetAmount = val.StringFixed(4)

		val, _ = decimal.NewFromString(data.D[k].BetAmount)
		data.D[k].BetAmount = val.StringFixed(4)

		data.D[k].BetTime = parseTime(v.BetTime)
	}

	data.S = uint16(pageSize)
	return data, nil
}

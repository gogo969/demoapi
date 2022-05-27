package model

import (
	"database/sql"
	"errors"
	"fmt"
	"reportapi/contrib/helper"

	g "github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
)

// DividendReport 红利报表
type DividendReport struct {
	ID         string  `db:"id"          json:"id"          `
	ReportTime int     `db:"report_time" json:"report_time" ` // 投注日期 yyyy-MM-dd 00:00:00的时间戳
	ReportType int     `db:"report_type" json:"report_type" ` // 211 平台红利、212 升级红利、213 生日红利、214 每月红利、215 红包红利、216 维护补偿、217 存款优惠、218 活动红利、219 推荐红利、220 红利调整、221 负数清零
	Amount     float64 `db:"amount"      json:"amount"      ` // 发放金额
	Prefix     string  `db:"prefix"      json:"prefix"      ` // 站点前缀
}

type DividendReportData struct {
	D      []DividendReport `json:"d"`
	T      int64            `json:"t"`
	S      int              `json:"s"`
	Amount float64          `json:"amount"`
}

func DividendList(ex g.Ex, startTime, endTime string, page, pageSize int) (DividendReportData, error) {

	data := DividendReportData{}
	ex["prefix"] = meta.Prefix

	if startTime != "" && endTime != "" {

		startAt, err := helper.TimeToLoc(startTime, loc)
		if err != nil {
			return data, errors.New(helper.DateTimeErr)
		}

		endAt, err := helper.TimeToLoc(endTime, loc)
		if err != nil {
			return data, errors.New(helper.DateTimeErr)
		}
		ex["report_time"] = g.Op{"between": exp.NewRangeVal(startAt, endAt)}
	}

	if page == 1 {
		// 拼装sql
		countQuery, _, _ := dialect.From("tbl_report_dividend").Select(g.COUNT(1)).Where(ex).ToSQL()
		err := meta.ReportDB.Get(&data.T, countQuery)
		if err != nil && err != sql.ErrNoRows {
			fmt.Println(err, countQuery)
			return data, pushLog(err, helper.DBErr)
		}

		if data.T == 0 {
			return data, nil
		}
	}

	offset := (page - 1) * pageSize
	dataQuery, _, _ := dialect.From("tbl_report_dividend").Select("report_time", "report_type", "amount", "prefix", "id").
		Where(ex).Offset(uint(offset)).Limit(uint(pageSize)).Order(g.L("report_time").Desc()).ToSQL()
	err := meta.ReportDB.Select(&data.D, dataQuery)
	if err != nil {
		return data, pushLog(err, helper.DBErr)
	}

	// 计算总计
	dataQuery, _, _ = dialect.From("tbl_report_dividend").Select(g.SUM("amount")).Where(ex).ToSQL()
	err = meta.ReportDB.Get(&data.Amount, dataQuery)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println(err, dataQuery)
		return data, pushLog(err, helper.DBErr)
	}

	data.S = pageSize
	return data, nil
}

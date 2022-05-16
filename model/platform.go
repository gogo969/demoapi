package model

import (
	"errors"
	"fmt"
	g "github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
	"github.com/shopspring/decimal"
	"reportapi/contrib/helper"
	"time"
)

type PlatReportItem struct {
	ReportTime                 string  `db:"report_time" json:"report_time"`                   //投注日期 yyyy-MM-dd 00:00:00的时间戳
	RegistCount                string  `db:"regist_count" json:"regist_count"`                 //注册人数
	DepositCount               string  `db:"deposit_count" json:"deposit_count"`               //首存人数
	ActiveCount                string  `db:"active_count" json:"active_count"`                 //活跃人数
	ConversionRate             string  `db:"conversion_rate" json:"conversion_rate"`           //转化率
	FirstDepositAmount         string  `db:"first_deposit_amount" json:"first_deposit_amount"` //首存额
	SecondDepositAmount        float64 `json:"second_deposit_amount" db:"second_deposit_amount"`
	AvgFirstDepositAmount      string  `db:"avg_first_deposit_amount" json:"avg_first_deposit_amount"` //人均首存
	DepositMemCount            string  `db:"deposit_mem_count" json:"deposit_mem_count"`               //存款人数
	WithdrawalMemCount         string  `db:"withdrawal_mem_count" json:"withdrawal_mem_count"`         //取款人数
	DepositAmount              string  `db:"deposit_amount" json:"deposit_amount"`                     //存款额
	WithdrawalAmount           string  `db:"withdrawal_amount" json:"withdrawal_amount"`               //取款额
	DepositWithdrawalSub       string  `db:"deposit_withdrawal_sub" json:"deposit_withdrawal_sub"`     //存取差
	DepositWithdrawalRate      string  `db:"deposit_withdrawal_rate" json:"deposit_withdrawal_rate"`   //提存率
	BetMemCount                string  `db:"bet_mem_count" json:"bet_mem_count"`                       //投注人數
	IpCount                    string  `db:"ip_count" json:"ip_count"`                                 //投注人數
	DeviceCount                string  `db:"device_count" json:"device_count"`                         //投注人數
	BetAmount                  string  `db:"bet_amount" json:"bet_amount"`                             //投注额
	ValidBetAmount             string  `db:"valid_bet_amount" json:"valid_bet_amount"`                 //有效投注额
	CompanyNetAmount           string  `db:"company_net_amount" json:"company_net_amount"`             //公司输赢
	ProfitAmount               string  `db:"profit_amount" json:"profit_amount"`                       //盈余比例
	AdjustAmount               string  `db:"adjust_amount" json:"adjust_amount"`                       //分数调整
	DividendAmount             string  `db:"dividend_amount" json:"dividend_amount"`                   //红利
	RebateAmount               string  `db:"rebate_amount" json:"rebate_amount"`                       //返水
	AgentAmount                string  `db:"agent_amount" json:"agent_amount"`                         //代理佣金
	Presettle                  string  `db:"presettle" json:"presettle"`                               //提前结算
	CompanyRevenue             string  `db:"company_revenue" json:"company_revenue"`                   //公司收入
	EfficientActiveCount       string  `db:"efficient_active_count" json:"efficient_active_count"`     //有效日活
	FirstDepositCount          string  `db:"first_deposit_count" json:"first_deposit_count"`           //首存人数
	SecondDepositCount         string  `json:"second_deposit_count" db:"second_deposit_count"`
	AvgCompanyNetAmount        string  `db:"avg_company_net_amount" json:"avg_company_net_amount"`               //公司平均输赢
	ReturnRate                 string  `json:"return_rate" db:"return_rate"`                                     //回存率
	OneTimeUser                string  `json:"one_time_user" db:"one_time_user"`                                 //一次性用户
	OneTimeShare               string  `json:"one_time_share" db:"one_time_share"`                               //一次性占比
	FirstDepositConversionRate string  `json:"first_deposit_conversion_rate" db:"first_deposit_conversion_rate"` //首存转化率
	KillRate                   string  `json:"kill_rate" db:"kill_rate"`                                         //杀率
	Multiple                   string  `json:"multiple" db:"multiple"`                                           //倍数
}

type PlatformReportData struct {
	D   []PlatReportItem `json:"d"`
	T   int64            `json:"t"`
	S   int              `json:"s"`
	Agg PlatReportItem   `json:"agg"`
}

type ComplexReportData struct {
	WinRate             string `db:"win_rate" json:"win_rate"`                         //盈率
	RegistCount         string `db:"regist_count" json:"regist_count"`                 //注册人数
	DepositCount        string `db:"deposit_count" json:"deposit_count"`               //首存人数
	FirstDepositAmount  string `db:"first_deposit_amount" json:"first_deposit_amount"` //首存额
	SecondDepositAmount string `db:"second_deposit_amount" json:"second_deposit_amount"`
	SecondDepositCount  string `db:"second_deposit_count" json:"second_deposit_count"`
	DepositMemCount     string `db:"deposit_mem_count" json:"deposit_mem_count"`       //存款人数
	WithdrawalMemCount  string `db:"withdrawal_mem_count" json:"withdrawal_mem_count"` //取款人数
	DepositAmount       string `db:"deposit_amount" json:"deposit_amount"`             //存款额
	WithdrawalAmount    string `db:"withdrawal_amount" json:"withdrawal_amount"`       //取款额
	BetMemCount         string `db:"bet_mem_count" json:"bet_mem_count"`               //投注人數
	BetAmount           string `db:"bet_amount" json:"bet_amount"`                     //投注额
	ValidBetAmount      string `db:"valid_bet_amount" json:"valid_bet_amount"`         //有效投注额
	CompanyNetAmount    string `db:"company_net_amount" json:"company_net_amount"`     //公司输赢
	Presettle           string `db:"presettle" json:"presettle"`                       //提前结算
	AdjustAmount        string `db:"adjust_amount" json:"adjust_amount"`               //分数调整
	DividendAmount      string `db:"dividend_amount" json:"dividend_amount"`           //红利
	RebateAmount        string `db:"rebate_amount" json:"rebate_amount"`               //返水
	CompanyRevenue      string `db:"company_revenue" json:"company_revenue"`           //公司收入
	FirstDepositCount   string `db:"first_deposit_count" json:"first_deposit_count"`   //首存人数
	MomoAmount          string `db:"momo_amount" json:"momo_amount"`                   //momo金额
	ZaloAmount          string `db:"zalo_amount" json:"zalo_amount"`                   //zalo金额
	OnlineAmount        string `db:"online_amount" json:"online_amount"`               //银行直连金额
	CardsAmount         string `db:"cards_amount" json:"cards_amount"`                 //卡转卡金额
	CoinpayAmount       string `db:"coinpay_amount" json:"coinpay_amount"`             //数字货币金额
	ViettelpayAmount    string `db:"viettelpay_amount" json:"viettelpay_amount"`       //viettelpay金额
	QrbankingAmount     string `db:"qrbanking_amount" json:"qrbanking_amount"`         //银行扫码金额
	UsdtAmount          string `db:"usdt_amount" json:"usdt_amount"`                   //线下usdt金额
	ManualAmount        string `db:"manual_amount" json:"manual_amount"`               //线下存款金额
	MomoCount           string `db:"momo_count" json:"momo_count"`                     //momo人数
	ZaloCount           string `db:"zalo_count" json:"zalo_count"`                     //zalo人数
	OnlineCount         string `db:"online_count" json:"online_count"`                 //银行直连人数
	CardsCount          string `db:"cards_count" json:"cards_count"`                   //卡转卡人数
	CoinpayCount        string `db:"coinpay_count" json:"coinpay_count"`               //数字货币人数
	ViettelpayCount     string `db:"viettelpay_count" json:"viettelpay_count"`         //viettelpay人数
	QrbankingCount      string `db:"qrbanking_count" json:"qrbanking_count"`           //银行扫码人数
	UsdtCount           string `db:"usdt_count" json:"usdt_count"`                     //线下usdt人数
	ManualCount         string `db:"manual_count" json:"manual_count"`                 //线下存款人数
}

type DepositsChannel struct {
	Id            string  `db:"id" json:"id"`
	ReportTime    int64   `db:"report_time" json:"report_time"`
	Prefix        string  `db:"prefix" json:"prefix"`
	ChannelId     int     `db:"channel_id" json:"channel_id"`
	DepositNum    int64   `db:"deposit_num" json:"deposit_num"`
	DepositAmount float64 `db:"deposit_amount" json:"deposit_amount"`
}

// PlatformReport 报表中心-平台报表
func PlatformReport(page, pageSize, flag, dateFlag, timeFlag int, depositStart, depositEnd, betAmountStart,
	betAmountEnd, depositCountStart, depositCountEnd, netAmountStart, netAmountEnd, startDate, endDate string) (PlatformReportData, error) {

	ex := g.Ex{"prefix": meta.Prefix}
	data := PlatformReportData{S: pageSize}

	startAt, err := helper.TimeToLoc(startDate, loc) // 秒级时间戳
	if err != nil {
		return data, err
	}

	endAt, err := helper.TimeToLoc(endDate, loc)
	if err != nil {
		return data, errors.New(helper.DateTimeErr)
	}

	if startAt > endAt {
		return data, errors.New(helper.QueryTimeRangeErr)
	}
	if len(depositStart) > 0 {
		ex["deposit_amount"] = g.Op{"gte": depositStart}
	}
	if len(depositEnd) > 0 {
		ex["deposit_amount"] = g.Op{"lte": depositEnd}
	}
	if len(depositStart) > 0 && len(depositEnd) > 0 {
		ex["deposit_amount"] = g.Op{"between": exp.NewRangeVal(depositStart, depositEnd)}
	}

	if len(betAmountStart) > 0 {
		ex["bet_amount"] = g.Op{"gte": betAmountStart}
	}
	if len(betAmountEnd) > 0 {
		ex["bet_amount"] = g.Op{"lte": betAmountEnd}
	}
	if len(betAmountStart) > 0 && len(betAmountEnd) > 0 {
		ex["bet_amount"] = g.Op{"between": exp.NewRangeVal(betAmountStart, betAmountEnd)}
	}

	if len(depositCountStart) > 0 {
		ex["deposit_count"] = g.Op{"gte": depositCountStart}
	}
	if len(depositCountEnd) > 0 {
		ex["deposit_count"] = g.Op{"lte": depositCountEnd}
	}

	if len(depositCountStart) > 0 && len(depositCountEnd) > 0 {
		ex["deposit_count"] = g.Op{"between": exp.NewRangeVal(depositCountStart, depositCountEnd)}
	}

	if len(netAmountStart) > 0 {
		ex["company_net_amount"] = g.Op{"gte": netAmountStart}
	}
	if len(netAmountEnd) > 0 {
		ex["company_net_amount"] = g.Op{"lte": netAmountEnd}
	}
	if len(netAmountStart) > 0 && len(netAmountEnd) > 0 {
		ex["company_net_amount"] = g.Op{"between": exp.NewRangeVal(netAmountStart, netAmountEnd)}
	}

	tableName := "tbl_report_platform"

	if flag == ReportFlagDay && timeFlag == ReportTimeFlagSingle { //单日日报
		if dateFlag == 1 {
			ex["report_type"] = 1
		}
		if dateFlag == 2 {
			ex["report_type"] = 2
		}
		return platformReportSingleDay(startAt, endAt, page, pageSize, tableName, ex)
	}

	if flag == ReportFlagDay && timeFlag == ReportTimeFlagPart { //按时间段日报
		if dateFlag == 1 {
			ex["report_type"] = 1
		}
		if dateFlag == 2 {
			ex["report_type"] = 2
		}
		return platformReportPeriodDay(startAt, endAt, page, pageSize, tableName, ex)
	}

	if flag == ReportFlagMonth && timeFlag == ReportTimeFlagSingle { //单月月报
		if dateFlag == 1 {
			ex["report_type"] = 3
		}
		if dateFlag == 2 {
			ex["report_type"] = 4
		}
		return platformReportSingleMonth(startAt, endAt, page, pageSize, tableName, ex)
	}

	if flag == ReportFlagMonth && timeFlag == ReportTimeFlagPart { //按时间段月报
		if dateFlag == 1 {
			ex["report_type"] = 3
		}
		if dateFlag == 2 {
			ex["report_type"] = 4
		}
		return platformReportPeriodMonth(startAt, endAt, page, pageSize, tableName, ex)
	}

	return data, errors.New(helper.ParamErr)
}

func platformReportCol() []interface{} {
	return []interface{}{
		g.SUM("regist_count").As("regist_count"),
		g.SUM("deposit_count").As("deposit_count"),
		g.SUM("conversion_rate").As("conversion_rate"),
		g.SUM("active_count").As("active_count"),
		g.SUM("efficient_active_count").As("efficient_active_count"),
		g.SUM("first_deposit_amount").As("first_deposit_amount"),
		g.SUM("first_deposit_count").As("first_deposit_count"),
		g.SUM("second_deposit_amount").As("second_deposit_amount"),
		g.SUM("second_deposit_count").As("second_deposit_count"),
		g.AVG("avg_first_deposit_amount").As("avg_first_deposit_amount"),
		g.SUM("deposit_mem_count").As("deposit_mem_count"),
		g.SUM("withdrawal_mem_count").As("withdrawal_mem_count"),
		g.SUM("deposit_amount").As("deposit_amount"),
		g.SUM("withdrawal_amount").As("withdrawal_amount"),
		g.SUM("deposit_withdrawal_sub").As("deposit_withdrawal_sub"),
		g.SUM("deposit_withdrawal_rate").As("deposit_withdrawal_rate"),
		g.SUM("bet_mem_count").As("bet_mem_count"),
		g.SUM("device_count").As("device_count"),
		g.SUM("ip_count").As("ip_count"),
		g.SUM("bet_amount").As("bet_amount"),
		g.SUM("valid_bet_amount").As("valid_bet_amount"),
		g.SUM("company_net_amount").As("company_net_amount"),
		g.SUM("avg_company_net_amount").As("avg_company_net_amount"),
		g.SUM("profit_amount").As("profit_amount"),
		g.SUM("adjust_amount").As("adjust_amount"),
		g.SUM("dividend_amount").As("dividend_amount"),
		g.SUM("rebate_amount").As("rebate_amount"),
		g.SUM("agent_amount").As("agent_amount"),
		g.SUM("presettle").As("presettle"),
		g.SUM("company_revenue").As("company_revenue")}
}

func platformFinanceReportCol() []interface{} {
	return []interface{}{
		g.SUM("regist_count").As("regist_count"),
		g.SUM("deposit_count").As("deposit_count"),
		g.SUM("first_deposit_amount").As("first_deposit_amount"),
		g.SUM("first_deposit_count").As("first_deposit_count"),
		g.SUM("second_deposit_amount").As("second_deposit_amount"),
		g.SUM("second_deposit_count").As("second_deposit_count"),
		g.SUM("deposit_mem_count").As("deposit_mem_count"),
		g.SUM("withdrawal_mem_count").As("withdrawal_mem_count"),
		g.SUM("deposit_amount").As("deposit_amount"),
		g.SUM("withdrawal_amount").As("withdrawal_amount"),
		g.SUM("bet_mem_count").As("bet_mem_count"),
		g.SUM("bet_amount").As("bet_amount"),
		g.SUM("valid_bet_amount").As("valid_bet_amount"),
		g.SUM("company_net_amount").As("company_net_amount"),
		g.SUM("adjust_amount").As("adjust_amount"),
		g.SUM("dividend_amount").As("dividend_amount"),
		g.SUM("rebate_amount").As("rebate_amount"),
		g.SUM("presettle").As("presettle"),
		g.SUM("company_revenue").As("company_revenue")}
}

func platformReportSingleDay(startAt, endAt int64, page, pageSize int, tableName string, ex g.Ex) (PlatformReportData, error) {

	data := PlatformReportData{S: pageSize}
	ex["report_time"] = g.Op{"between": exp.NewRangeVal(startAt, endAt)}

	if page == 1 {

		var t []total
		totalQuery, _, _ := dialect.From(tableName).
			Select(g.C("report_time").As("report_time"), g.COUNT("id").As("count")).Where(ex).GroupBy("report_time").ToSQL()
		err := meta.ReportDB.Select(&t, totalQuery)
		if err != nil {
			return data, pushLog(fmt.Errorf("%s,[%s]", err.Error(), totalQuery), "db", helper.DBErr)
		}

		data.T = int64(len(t))
		if data.T == 0 {
			return data, nil
		}
	}

	col := platformReportCol()
	offset := (page - 1) * pageSize
	col = append(col, g.C("report_time").As("report_time"))
	query, _, _ := dialect.From(tableName).Select(col...).Where(ex).
		GroupBy("report_time").Order(g.C("report_time").Desc()).Offset(uint(offset)).Limit(uint(pageSize)).ToSQL()
	fmt.Println(query)
	err := meta.ReportDB.Select(&data.D, query)
	if err != nil {
		return data, pushLog(fmt.Errorf("%s,[%s]", err.Error(), query), "db", helper.DBErr)
	}

	col = platformReportCol()
	aggQuery, _, _ := dialect.From(tableName).Select(col...).Where(ex).ToSQL()
	err = meta.ReportDB.Get(&data.Agg, aggQuery)
	if err != nil {
		return data, pushLog(fmt.Errorf("%s,[%s]", err.Error(), query), "db", helper.DBErr)
	}

	for k, v := range data.D {

		data.D[k] = reportPlatItemFormat(v)
		data.D[k].ReportTime = parseDay(v.ReportTime)
	}

	data.Agg = reportPlatItemFormat(data.Agg)

	return data, nil
}

func platformReportPeriodDay(startAt, endAt int64, page, pageSize int, tableName string, ex g.Ex) (PlatformReportData, error) {

	var t []total
	data := PlatformReportData{S: pageSize}
	ex["report_time"] = g.Op{"between": exp.NewRangeVal(startAt, endAt)}
	totalQuery, _, _ := dialect.From(tableName).
		Select(g.C("prefix").As("prefix"), g.COUNT("id").As("count")).Where(ex).GroupBy("prefix").ToSQL()
	err := meta.ReportDB.Select(&t, totalQuery)
	if err != nil {
		return data, pushLog(fmt.Errorf("%s,[%s]", err.Error(), totalQuery), "db", helper.DBErr)
	}

	data.T = int64(len(t))
	if data.T == 0 {
		return data, nil
	}

	col := platformReportCol()
	offset := (page - 1) * pageSize
	col = append(col, g.V("0").As("report_time"))
	query, _, _ := dialect.From(tableName).Select(col...).
		GroupBy("prefix").Where(ex).Offset(uint(offset)).Limit(uint(pageSize)).ToSQL()

	err = meta.ReportDB.Select(&data.D, query)
	if err != nil {
		return data, pushLog(fmt.Errorf("%s,[%s]", err.Error(), query), "db", helper.DBErr)
	}

	col = platformReportCol()
	agg, _, _ := dialect.From(tableName).Select(col...).Where(ex).ToSQL()
	err = meta.ReportDB.Get(&data.Agg, agg)
	if err != nil {
		return data, pushLog(fmt.Errorf("%s,[%s]", err.Error(), query), "db", helper.DBErr)
	}

	for k, v := range data.D {

		data.D[k] = reportPlatItemFormat(v)
		data.D[k].ReportTime = parsePart(startAt, endAt, "d")
	}

	data.Agg = reportPlatItemFormat(data.Agg)

	return data, nil
}

func platformReportSingleMonth(startAt, endAt int64, page, pageSize int, tableName string, ex g.Ex) (PlatformReportData, error) {

	var t []total
	data := PlatformReportData{S: pageSize}
	ex["report_month"] = g.Op{"between": exp.NewRangeVal(startAt, endAt)}
	totalQuery, _, _ := dialect.From(tableName).
		Select(g.C("report_month").As("report_time"), g.COUNT("id").As("count")).Where(ex).GroupBy("report_month").ToSQL()
	err := meta.ReportDB.Select(&t, totalQuery)
	if err != nil {
		return data, pushLog(fmt.Errorf("%s,[%s]", err.Error(), totalQuery), "db", helper.DBErr)
	}

	data.T = int64(len(t))
	if data.T == 0 {
		return data, nil
	}

	col := platformReportCol()
	offset := (page - 1) * pageSize
	col = append(col, g.C("report_month").As("report_time"))
	query, _, _ := dialect.From(tableName).Select(col...).
		GroupBy("report_month").Where(ex).Order(g.C("report_month").Desc()).Offset(uint(offset)).Limit(uint(pageSize)).ToSQL()

	err = meta.ReportDB.Select(&data.D, query)
	if err != nil {
		return data, pushLog(fmt.Errorf("%s,[%s]", err.Error(), query), "db", helper.DBErr)
	}

	col = platformReportCol()
	agg, _, _ := dialect.From(tableName).Select(col...).Where(ex).ToSQL()
	err = meta.ReportDB.Get(&data.Agg, agg)
	if err != nil {
		return data, pushLog(fmt.Errorf("%s,[%s]", err.Error(), query), "db", helper.DBErr)
	}

	for k, v := range data.D {

		data.D[k] = reportPlatItemFormat(v)
		data.D[k].ReportTime = parseMonth(v.ReportTime)
	}

	data.Agg = reportPlatItemFormat(data.Agg)

	return data, nil
}

func platformReportPeriodMonth(startAt, endAt int64, page, pageSize int, tableName string, ex g.Ex) (PlatformReportData, error) {

	var t []total
	data := PlatformReportData{S: pageSize}
	ex["report_month"] = g.Op{"between": exp.NewRangeVal(startAt, endAt)}
	totalQuery, _, _ := dialect.From(tableName).
		Select(g.C("prefix").As("prefix"), g.COUNT("id").As("count")).Where(ex).GroupBy("prefix").ToSQL()
	err := meta.ReportDB.Select(&t, totalQuery)
	if err != nil {
		return data, pushLog(fmt.Errorf("%s,[%s]", err.Error(), totalQuery), "db", helper.DBErr)
	}

	data.T = int64(len(t))
	if data.T == 0 {
		return data, nil
	}

	col := platformReportCol()
	offset := (page - 1) * pageSize
	col = append(col, g.V("0").As("report_time"))
	query, _, _ := dialect.From(tableName).Select(col...).
		GroupBy("prefix").Where(ex).Offset(uint(offset)).Limit(uint(pageSize)).ToSQL()

	err = meta.ReportDB.Select(&data.D, query)
	if err != nil {
		return data, pushLog(fmt.Errorf("%s,[%s]", err.Error(), query), "db", helper.DBErr)
	}

	col = platformReportCol()
	agg, _, _ := dialect.From(tableName).Select(col...).Where(ex).ToSQL()
	err = meta.ReportDB.Get(&data.Agg, agg)
	if err != nil {
		return data, pushLog(fmt.Errorf("%s,[%s]", err.Error(), query), "db", helper.DBErr)
	}

	for k, v := range data.D {

		data.D[k] = reportPlatItemFormat(v)
		data.D[k].ReportTime = parsePart(startAt, endAt, "m")
	}

	data.Agg = reportPlatItemFormat(data.Agg)

	return data, nil
}

func reportPlatItemFormat(data PlatReportItem) PlatReportItem {

	min, _ := decimal.NewFromString("0")
	rate := decimal.NewFromInt(100)

	firstDepositCount, _ := decimal.NewFromString(data.FirstDepositCount)
	depositMemCount, _ := decimal.NewFromString(data.DepositMemCount)
	registCount, _ := decimal.NewFromString(data.RegistCount)
	if !registCount.Equal(min) {
		val := depositMemCount.Div(registCount)
		val = val.Mul(rate)
		data.ConversionRate = val.StringFixed(4)
		data.FirstDepositConversionRate = firstDepositCount.Div(registCount).Mul(rate).StringFixed(4)
	} else {
		data.ConversionRate = "0.0000"
		data.FirstDepositConversionRate = "0.0000"
	}

	firstDepositAmount, _ := decimal.NewFromString(data.FirstDepositAmount)
	secondDepositCount, _ := decimal.NewFromString(data.SecondDepositCount)
	data.OneTimeUser = firstDepositCount.Sub(secondDepositCount).StringFixed(4)

	if !firstDepositCount.Equal(min) {
		val := firstDepositAmount.Div(firstDepositCount)
		data.AvgFirstDepositAmount = val.StringFixed(4)
		data.ReturnRate = secondDepositCount.Div(firstDepositCount).Mul(rate).StringFixed(4)
		data.OneTimeShare = (firstDepositCount.Sub(secondDepositCount)).Div(firstDepositCount).StringFixed(4)
	} else {
		data.AvgFirstDepositAmount = "0.0000"
		data.ReturnRate = "0.000"
		data.OneTimeShare = "0.000"
	}

	depositAmount, _ := decimal.NewFromString(data.DepositAmount)
	withdrawalAmount, _ := decimal.NewFromString(data.WithdrawalAmount)

	if !depositAmount.Equal(min) {
		val := withdrawalAmount.Div(depositAmount)
		val = val.Mul(rate)
		data.DepositWithdrawalRate = val.StringFixed(4)
		betAmount, _ := decimal.NewFromString(data.BetAmount)
		data.Multiple = betAmount.Div(depositAmount).Mul(rate).StringFixed(4)
	} else {
		data.DepositWithdrawalRate = "0.0000"
		data.Multiple = "0.0000"
	}

	val := depositAmount.Sub(withdrawalAmount)
	data.DepositWithdrawalSub = val.StringFixed(4)

	companyNetAmount, _ := decimal.NewFromString(data.CompanyNetAmount)
	presettle, _ := decimal.NewFromString(data.Presettle)
	val = companyNetAmount.Add(presettle)

	betAmount, _ := decimal.NewFromString(data.BetAmount)
	if !betAmount.Equal(min) {
		val = val.Div(betAmount)
		val = val.Mul(rate)
		data.ProfitAmount = val.StringFixed(4)
		companyRevenue, _ := decimal.NewFromString(data.CompanyRevenue)
		data.KillRate = companyRevenue.Div(betAmount).Mul(rate).StringFixed(4)

	} else {
		data.ProfitAmount = "0.0000"
		data.KillRate = "0.0000"
	}

	val, _ = decimal.NewFromString(data.FirstDepositAmount)
	data.FirstDepositAmount = val.StringFixed(4)

	val, _ = decimal.NewFromString(data.DepositAmount)
	data.DepositAmount = val.StringFixed(4)

	val, _ = decimal.NewFromString(data.WithdrawalAmount)
	data.WithdrawalAmount = val.StringFixed(4)

	val, _ = decimal.NewFromString(data.BetAmount)
	data.BetAmount = val.StringFixed(4)

	val, _ = decimal.NewFromString(data.ValidBetAmount)
	data.ValidBetAmount = val.StringFixed(4)

	val, _ = decimal.NewFromString(data.CompanyNetAmount)
	data.CompanyNetAmount = val.StringFixed(4)

	val, _ = decimal.NewFromString(data.AdjustAmount)
	data.AdjustAmount = val.StringFixed(4)

	val, _ = decimal.NewFromString(data.DividendAmount)
	data.DividendAmount = val.StringFixed(4)

	val, _ = decimal.NewFromString(data.RebateAmount)
	data.RebateAmount = val.StringFixed(4)

	val, _ = decimal.NewFromString(data.AgentAmount)
	data.AgentAmount = val.StringFixed(4)

	val, _ = decimal.NewFromString(data.Presettle)
	data.Presettle = val.StringFixed(4)

	val, _ = decimal.NewFromString(data.CompanyRevenue)
	data.CompanyRevenue = val.StringFixed(4)

	return data
}

func ComplexReport(flag int, startDate, endDate string) (ComplexReportData, error) {

	ex := g.Ex{"prefix": meta.Prefix}
	if flag == 1 {
		ex["report_type"] = 2
	} else {
		ex["report_type"] = 4
	}
	data := ComplexReportData{}
	var financeList []DepositsChannel
	financeMap := map[int]DepositsChannel{}
	startAt, err := helper.TimeToLoc(startDate, loc) // 秒级时间戳
	if err != nil {
		return data, err
	}

	endAt, err := helper.TimeToLoc(endDate, loc)
	if err != nil {
		return data, errors.New(helper.DateTimeErr)
	}

	if startAt > endAt {
		return data, errors.New(helper.QueryTimeRangeErr)
	}
	if startAt > time.Now().Unix() {
		return data, errors.New(helper.QueryTimeRangeErr)
	}
	ex["report_time"] = g.Op{"between": exp.NewRangeVal(startAt, endAt)}

	col := platformFinanceReportCol()
	query, _, _ := dialect.From("tbl_report_platform").Select(col...).Where(ex).ToSQL()
	fmt.Println(query)
	err = meta.ReportDB.Get(&data, query)
	if err != nil {
		return data, pushLog(fmt.Errorf("%s,[%s]", err.Error(), query), "db", helper.DBErr)
	}

	cm, _ := decimal.NewFromString(data.CompanyRevenue)
	vb, _ := decimal.NewFromString(data.ValidBetAmount)
	if vb.GreaterThan(decimal.Zero) {
		data.WinRate = cm.Div(vb).StringFixed(3)
	}

	delete(ex, "report_type")
	query2, _, _ := dialect.From("tbl_report_finance").Select(g.SUM("deposit_num").As("deposit_num"),
		g.SUM("deposit_amount").As("deposit_amount"), g.C("channel_id").As("channel_id")).Where(ex).GroupBy("channel_id").ToSQL()

	err = meta.ReportDB.Select(&financeList, query2)
	if err != nil {
		return data, pushLog(fmt.Errorf("%s,[%s]", err.Error(), query), "db", helper.DBErr)
	}

	for _, v := range financeList {
		financeMap[v.ChannelId] = v
	}

	data.MomoAmount = fmt.Sprintf(`%f`, financeMap[1].DepositAmount)
	data.MomoCount = fmt.Sprintf(`%d`, financeMap[1].DepositNum)

	data.ZaloAmount = fmt.Sprintf(`%f`, financeMap[2].DepositAmount)
	data.ZaloCount = fmt.Sprintf(`%d`, financeMap[2].DepositNum)

	data.OnlineAmount = fmt.Sprintf(`%f`, financeMap[3].DepositAmount)
	data.OnlineCount = fmt.Sprintf(`%d`, financeMap[3].DepositNum)

	data.CardsAmount = fmt.Sprintf(`%f`, financeMap[4].DepositAmount)
	data.CardsCount = fmt.Sprintf(`%d`, financeMap[4].DepositNum)

	data.CoinpayAmount = fmt.Sprintf(`%f`, financeMap[5].DepositAmount)
	data.CoinpayCount = fmt.Sprintf(`%d`, financeMap[5].DepositNum)

	data.ViettelpayAmount = fmt.Sprintf(`%f`, financeMap[6].DepositAmount)
	data.ViettelpayCount = fmt.Sprintf(`%d`, financeMap[6].DepositNum)

	data.QrbankingAmount = fmt.Sprintf(`%f`, financeMap[8].DepositAmount)
	data.QrbankingCount = fmt.Sprintf(`%d`, financeMap[8].DepositNum)

	data.UsdtAmount = fmt.Sprintf(`%f`, financeMap[10].DepositAmount)
	data.UsdtCount = fmt.Sprintf(`%d`, financeMap[10].DepositNum)

	data.ManualAmount = fmt.Sprintf(`%f`, financeMap[9].DepositAmount)
	data.ManualCount = fmt.Sprintf(`%d`, financeMap[9].DepositNum)

	return data, nil

}

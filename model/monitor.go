package model

import (
	"errors"
	"fmt"
	g "github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
	"github.com/shopspring/decimal"
	"reportApi2/contrib/helper"
)

type MonitorItem struct {
	ID                           string `db:"id" json:"id"`                                                             //id
	ReportTime                   string `db:"report_time" json:"report_time"`                                           //
	Prefix                       string `db:"prefix" json:"prefix"`                                                     //站点前缀
	RegistCount                  string `db:"regist_count" json:"regist_count"`                                         //注册人数
	DepositCount                 string `db:"deposit_count" json:"deposit_count"`                                       //首存人数
	ConversionRate               string `db:"conversion_rate" json:"conversion_rate"`                                   //转化率
	DepositAmount                string `db:"deposit_amount" json:"deposit_amount"`                                     //存款额
	WithdrawalAmount             string `db:"withdrawal_amount" json:"withdrawal_amount"`                               //取款额
	DepositWithdrawalSub         string `db:"deposit_withdrawal_sub" json:"deposit_withdrawal_sub"`                     //存取差
	DepositSuccessRate           string `db:"deposit_success_rate" json:"deposit_success_rate"`                         //存款成功率
	DepositChainRatio            string `db:"deposit_chain_ratio" json:"deposit_chain_ratio"`                           //
	WithdrawalChainRatio         string `db:"withdrawal_chain_ratio" json:"withdrawal_chain_ratio"`                     //
	DepositYesterdayRatio        string `db:"deposit_yesterday_ratio" json:"deposit_yesterday_ratio"`                   //
	BetMemCount                  string `db:"bet_mem_count" json:"bet_mem_count"`                                       //投注人數
	DepositBetRate               string `db:"deposit_bet_rate" json:"deposit_bet_rate"`                                 //存投比
	ValidBetAmount               string `db:"valid_bet_amount" json:"valid_bet_amount"`                                 //有效投注额
	ValidBetAmountSub            string `db:"valid_bet_amount_sub" json:"valid_bet_amount_sub"`                         //
	ValidBetAmountRate           string `db:"valid_bet_amount_rate" json:"valid_bet_amount_rate"`                       //
	ValidBetAmountYesterdayRatio string `db:"valid_bet_amount_yesterday_ratio" json:"valid_bet_amount_yesterday_ratio"` //
	CompanyNetAmount             string `db:"company_net_amount" json:"company_net_amount"`                             //公司输赢
	Presettle                    string `db:"presettle" json:"presettle"`                                               //提前结算
}

type MonitorReportData struct {
	D []MonitorItem `json:"d"`
	T int64         `json:"t"`
	S int           `json:"s"`
}

type RegisterCount struct {
	ReportTime  string `db:"report_time" json:"report_time"`
	Prefix      string `db:"prefix" json:"prefix"`
	RegistCount int64  `db:"regist_count" json:"regist_count"`
}

type DepositCount struct {
	ReportTime   string `json:"report_time" db:"report_time"`
	Prefix       string `json:"prefix" db:"prefix"`
	DepositCount int64  `json:"deposit_count" db:"deposit_count"`
}

type DepositHour struct {
	ReportTime         string  `json:"report_time" db:"report_time"`
	Prefix             string  `json:"prefix" db:"prefix"`
	DepositAmount      float64 `json:"deposit_amount" db:"deposit_amount"`
	TotalDepositAmount float64 `json:"total_deposit_amount" db:"total_deposit_amount"`
}

type OrderHour struct {
	ReportTime       string  `json:"report_time" db:"report_time"`
	Prefix           string  `json:"prefix" db:"prefix"`
	BetMemCount      int64   `json:"bet_mem_count" db:"bet_mem_count"`
	ValidBetAmount   float64 `json:"valid_bet_amount" db:"valid_bet_amount"`
	CompanyNetAmount float64 `json:"company_net_amount" db:"company_net_amount"`
	Presettle        float64 `json:"presettle" db:"presettle"`
}

type WithdrawalHour struct {
	ReportTime       string  `json:"report_time" db:"report_time"`
	Prefix           string  `json:"prefix" db:"prefix"`
	WithdrawalAmount float64 `json:"withdrawal_amount" db:"withdrawal_amount"`
}

type DepositWithdrawal struct {
	Id                           string  `json:"id" db:"id"`
	ReportTime                   int64   `json:"report_time" db:"report_time"`
	Prefix                       string  `json:"prefix" db:"prefix"`
	RegistCount                  int64   `json:"regist_count" db:"regist_count"`
	DepositCount                 int64   `json:"deposit_count" db:"deposit_count"`
	ConversionRate               float64 `json:"conversion_rate" db:"conversion_rate"`
	DepositAmount                float64 `json:"deposit_amount" db:"deposit_amount"`
	WithdrawalAmount             float64 `json:"withdrawal_amount" db:"withdrawal_amount"`
	DepositWithdrawalSub         float64 `json:"deposit_withdrawal_sub" db:"deposit_withdrawal_sub"`
	DepositSuccessRate           float64 `json:"deposit_success_rate" db:"deposit_success_rate"`
	DepositChainRatio            float64 `json:"deposit_chain_ratio" db:"deposit_chain_ratio"`
	WithdrawalChainRatio         float64 `json:"withdrawal_chain_ratio" db:"withdrawal_chain_ratio"`
	DepositYesterdayRatio        float64 `json:"deposit_yesterday_ratio" db:"deposit_yesterday_ratio"`
	BetMemCount                  int64   `json:"bet_mem_count" db:"bet_mem_count"`
	DepositBetRate               float64 `json:"deposit_bet_rate" db:"deposit_bet_rate"`
	ValidBetAmount               float64 `json:"valid_bet_amount" db:"valid_bet_amount"`
	ValidBetAmountSub            float64 `json:"valid_bet_amount_sub" db:"valid_bet_amount_sub"`
	ValidBetAmountRate           float64 `json:"valid_bet_amount_rate" db:"valid_bet_amount_rate"`
	ValidBetAmountYesterdayRatio float64 `json:"valid_bet_amount_yesterday_ratio" db:"valid_bet_amount_yesterday_ratio"`
	CompanyNetAmount             float64 `json:"company_net_amount" db:"company_net_amount"`
	Presettle                    float64 `json:"presettle" db:"presettle"`
}

var (
	accessKeyMap = map[string]AccessKey{}      //商会汇总数据map
	registerMap  = map[string]RegisterCount{}  //注册数map
	depositCMap  = map[string]DepositCount{}   //存款次数map
	depositHMap  = map[string]DepositHour{}    //存款金额map
	withdrawMap  = map[string]WithdrawalHour{} //提款金额map
	orderMap     = map[string]OrderHour{}      //投注数据map
)

// ReportAccess 报表中心-实时存取监控
func MonitorReport(page, pageSize int, date string) (MonitorReportData, error) {

	ex := g.Ex{"prefix": meta.Prefix}
	data := MonitorReportData{}

	startAt, err := helper.TimeToLoc(date, loc)
	if err != nil {
		return data, errors.New(helper.DateTimeErr)
	}

	endAt := startAt + 24*3600

	ex["report_time"] = g.Op{"between": exp.NewRangeVal(startAt, endAt)}

	if page == 1 {

		totalQuery, _, _ := dialect.From("tbl_report_deposit_withdrawal").Select(g.COUNT(1)).Where(ex).ToSQL()
		err := meta.ReportDB.Get(&data.T, totalQuery)
		if err != nil {
			return data, pushLog(fmt.Errorf("%s,[%s]", err.Error(), totalQuery), "db", helper.DBErr)
		}

		if data.T == 0 {
			return data, nil
		}
	}

	offset := (page - 1) * pageSize
	query, _, _ := dialect.From("tbl_report_deposit_withdrawal").Select(accessCol...).
		Where(ex).Offset(uint(offset)).Order(g.C("report_time").Desc()).Limit(uint(pageSize)).ToSQL()
	err = meta.ReportDB.Select(&data.D, query)
	if err != nil {
		return data, pushLog(fmt.Errorf("%s,[%s]", err.Error(), query), "db", helper.DBErr)
	}

	for k, v := range data.D {

		conversionRate, _ := decimal.NewFromString(data.D[k].ConversionRate)
		data.D[k].ConversionRate = conversionRate.StringFixed(5)

		depositAmount, _ := decimal.NewFromString(data.D[k].DepositAmount)
		data.D[k].DepositAmount = depositAmount.StringFixed(3)

		withdrawalAmount, _ := decimal.NewFromString(data.D[k].WithdrawalAmount)
		data.D[k].WithdrawalAmount = withdrawalAmount.StringFixed(3)

		wepositWithdrawalSub, _ := decimal.NewFromString(data.D[k].DepositWithdrawalSub)
		data.D[k].DepositWithdrawalSub = wepositWithdrawalSub.StringFixed(3)

		depositSuccessRate, _ := decimal.NewFromString(data.D[k].DepositSuccessRate)
		data.D[k].DepositSuccessRate = depositSuccessRate.StringFixed(5)

		depositChainRatio, _ := decimal.NewFromString(data.D[k].DepositChainRatio)
		data.D[k].DepositChainRatio = depositChainRatio.StringFixed(5)

		withdrawalChainRatio, _ := decimal.NewFromString(data.D[k].WithdrawalChainRatio)
		data.D[k].WithdrawalChainRatio = withdrawalChainRatio.StringFixed(5)

		depositYesterdayRatio, _ := decimal.NewFromString(data.D[k].DepositYesterdayRatio)
		data.D[k].DepositYesterdayRatio = depositYesterdayRatio.StringFixed(5)

		depositBetRate, _ := decimal.NewFromString(data.D[k].DepositBetRate)
		data.D[k].DepositBetRate = depositBetRate.StringFixed(5)

		validBetAmount, _ := decimal.NewFromString(data.D[k].ValidBetAmount)
		data.D[k].ValidBetAmount = validBetAmount.StringFixed(3)

		validBetAmountSub, _ := decimal.NewFromString(data.D[k].ValidBetAmountSub)
		data.D[k].ValidBetAmountSub = validBetAmountSub.StringFixed(3)

		validBetAmountRate, _ := decimal.NewFromString(data.D[k].ValidBetAmountRate)
		data.D[k].ValidBetAmountRate = validBetAmountRate.StringFixed(5)

		validBetAmountYesterdayRatio, _ := decimal.NewFromString(data.D[k].ValidBetAmountYesterdayRatio)
		data.D[k].ValidBetAmountYesterdayRatio = validBetAmountYesterdayRatio.StringFixed(5)

		companyNetAmount, _ := decimal.NewFromString(data.D[k].CompanyNetAmount)
		data.D[k].CompanyNetAmount = companyNetAmount.StringFixed(3)
		data.D[k].ReportTime = parseHourTime(v.ReportTime)
	}

	data.S = pageSize

	return data, nil
}

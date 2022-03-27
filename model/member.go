package model

import (
	"fmt"
	g "github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
	"github.com/shopspring/decimal"
	"reportApi2/contrib/helper"
	"strconv"
	"strings"
)

type MemberReportData struct {
	D   []MemberReportShow `json:"d"`
	T   int64              `json:"t"`
	S   int                `json:"s"`
	Agg MemberReportShow   `json:"agg"`
}

type MemberReportShow struct {
	MemberReportItem
}

type MemberReportItem struct {
	ReportTime            string `db:"report_time" json:"report_time"` //投注日期 yyyy-MM-dd 00:00:00的时间戳
	Id                    string `json:"id" db:"id"`
	TopName               string `json:"top_name" db:"top_name"`
	TopUid                string `json:"top_uid" db:"top_uid"`
	ParentName            string `json:"parent_name" db:"parent_name"`
	ParentUid             string `json:"parent_uid" db:"parent_uid"`
	CommissionsId         string `json:"commissions_id" db:"commissions_id"`
	CommissionName        string `json:"commission_name" db:"commission_name"`
	MemCount              string `json:"mem_count" db:"mem_count"`
	SubordinateCount      string `json:"subordinate_count" db:"subordinate_count"`
	LoginCount            string `json:"login_count" db:"login_count"`
	RegistCount           string `json:"regist_count" db:"regist_count"`
	ActiveCount           string `json:"active_count" db:"active_count"`
	ConversionRate        string `json:"conversion_rate" db:"conversion_rate"`
	FirstDepositCount     string `json:"first_deposit_count" db:"first_deposit_count"`
	FirstDepositAmount    string `json:"first_deposit_amount" db:"first_deposit_amount"`
	AvgFirstDepositAmount string `json:"avg_first_deposit_amount" db:"avg_first_deposit_amount"`
	DepositMemCount       string `json:"deposit_mem_count" db:"deposit_mem_count"`
	WithdrawalMemCount    string `json:"withdrawal_mem_count" db:"withdrawal_mem_count"`
	DepositAmount         string `json:"deposit_amount" db:"deposit_amount"`
	WithdrawalAmount      string `json:"withdrawal_amount" db:"withdrawal_amount"`
	AdjustAmount          string `json:"adjust_amount" db:"adjust_amount"`
	BetMemCount           string `json:"bet_mem_count" db:"bet_mem_count"`
	BetAmount             string `json:"bet_amount" db:"bet_amount"`
	ValidBetAmount        string `json:"valid_bet_amount" db:"valid_bet_amount"`
	CompanyNetAmount      string `json:"company_net_amount" db:"company_net_amount"`
	ProfitAmount          string `json:"profit_amount" db:"profit_amount"`
	DividendAmount        string `json:"dividend_amount" db:"dividend_amount"`
	RebateAmount          string `json:"rebate_amount" db:"rebate_amount"`
	AgentAmount           string `json:"agent_amount" db:"agent_amount"`
	AdjustSystemAmount    string `json:"adjust_system_amount" db:"adjust_system_amount"`
	AdjustWinAmount       string `json:"adjust_win_amount" db:"adjust_win_amount"`
	Prefix                string `json:"prefix" db:"prefix"`
	Presettle             string `json:"presettle" db:"presettle"`
	CompanyRevenue        string `json:"company_revenue" db:"company_revenue"`
	Uid                   string `json:"uid" db:"uid"`
	DepositCount          string `json:"deposit_count" db:"deposit_count"`
	CreatedAt             string `json:"created_at" db:"created_at"`
	FirstDepositAt        string `json:"first_deposit_at" db:"first_deposit_at"`
	TimeOutDeposit        string `json:"time_out_deposit" db:"time_out_deposit"`
	TimeOutBet            string `json:"time_out_bet" db:"time_out_bet"`
	Username              string `json:"username" db:"username"`
	TimeOutLogin          string `json:"time_out_login" db:"time_out_login"`
	State                 string `json:"state" db:"state"`
	RegDevice             string `json:"reg_device" db:"reg_device"`
	SourceId              string `json:"source_id" db:"source_id"`
	FirstBetAt            string `json:"first_bet_at" db:"first_bet_at"`
	FirstBetAmount        string `json:"first_bet_amount" db:"first_bet_amount"`
	LastLoginDevice       string `json:"last_login_device" db:"last_login_device"`
	LastLoginAt           string `json:"last_login_at" db:"last_login_at"`
	TagNames              string `json:"tag_names" db:"tag_names"`
	Balance               string `json:"balance" db:"balance"`
	RegUrl                string `json:"reg_url" db:"reg_url"`
}

type MemberData struct {
	ReportTime            string `db:"report_time" json:"report_time"` //投注日期 yyyy-MM-dd 00:00:00的时间戳
	Id                    string `json:"id" db:"id"`
	TopName               string `json:"top_name" db:"top_name"`
	TopUid                string `json:"top_uid" db:"top_uid"`
	ParentName            string `json:"parent_name" db:"parent_name"`
	ParentUid             string `json:"parent_uid" db:"parent_uid"`
	CommissionsId         string `json:"commissions_id" db:"commissions_id"`
	CommissionName        string `json:"commission_name" db:"commission_name"`
	MemCount              string `json:"mem_count" db:"mem_count"`
	SubordinateCount      string `json:"subordinate_count" db:"subordinate_count"`
	LoginCount            string `json:"login_count" db:"login_count"`
	RegistCount           string `json:"regist_count" db:"regist_count"`
	ActiveCount           string `json:"active_count" db:"active_count"`
	ConversionRate        string `json:"conversion_rate" db:"conversion_rate"`
	FirstDepositCount     string `json:"first_deposit_count" db:"first_deposit_count"`
	FirstDepositAmount    string `json:"first_deposit_amount" db:"first_deposit_amount"`
	AvgFirstDepositAmount string `json:"avg_first_deposit_amount" db:"avg_first_deposit_amount"`
	DepositMemCount       string `json:"deposit_mem_count" db:"deposit_mem_count"`
	WithdrawalMemCount    string `json:"withdrawal_mem_count" db:"withdrawal_mem_count"`
	DepositAmount         string `json:"deposit_amount" db:"deposit_amount"`
	WithdrawalAmount      string `json:"withdrawal_amount" db:"withdrawal_amount"`
	AdjustAmount          string `json:"adjust_amount" db:"adjust_amount"`
	BetMemCount           string `json:"bet_mem_count" db:"bet_mem_count"`
	BetAmount             string `json:"bet_amount" db:"bet_amount"`
	ValidBetAmount        string `json:"valid_bet_amount" db:"valid_bet_amount"`
	CompanyNetAmount      string `json:"company_net_amount" db:"company_net_amount"`
	ProfitAmount          string `json:"profit_amount" db:"profit_amount"`
	DividendAmount        string `json:"dividend_amount" db:"dividend_amount"`
	RebateAmount          string `json:"rebate_amount" db:"rebate_amount"`
	AgentAmount           string `json:"agent_amount" db:"agent_amount"`
	AdjustSystemAmount    string `json:"adjust_system_amount" db:"adjust_system_amount"`
	AdjustWinAmount       string `json:"adjust_win_amount" db:"adjust_win_amount"`
	Prefix                string `json:"prefix" db:"prefix"`
	Presettle             string `json:"presettle" db:"presettle"`
	CompanyRevenue        string `json:"company_revenue" db:"company_revenue"`
	Uid                   string `json:"uid" db:"uid"`
	DepositCount          string `json:"deposit_count" db:"deposit_count"`
	CreatedAt             string `json:"created_at" db:"created_at"`
	FirstDepositAt        string `json:"first_deposit_at" db:"first_deposit_at"`
	TimeOutDeposit        string `json:"time_out_deposit" db:"time_out_deposit"`
	TimeOutBet            string `json:"time_out_bet" db:"time_out_bet"`
	Username              string `json:"username" db:"username"`
	TimeOutLogin          string `json:"time_out_login" db:"time_out_login"`
}

// MemberReport 报表中心-会员报表
func MemberReport(flag, dateFlag, timeFlag, page, pageSize int, timeOutBet, timeOutLogin, timeOutDeposit int,
	sRegStartTime, sRegEndTime, sFdepositStartTime, sFdepositEndTime string, parentName, parentUid, userName string,
	sStartTime, sEndTime, sCommissionId, sMainId string, ty int) (MemberReportData, error) {

	var data MemberReportData
	var regStartTime int64
	if len(sRegStartTime) > 0 {
		//判断日期
		t, err := helper.TimeToLoc(sRegStartTime, loc) // 秒级时间戳
		if err != nil {
			return data, err
		}
		regStartTime = t
	}
	var regEndTime int64
	if len(sRegEndTime) > 0 {
		t, err := helper.TimeToLoc(sRegEndTime, loc) // 秒级时间戳
		if err != nil {
			return data, err
		}
		regEndTime = t
	}
	var FdepositStartTime int64
	if len(sFdepositStartTime) > 0 {
		t, err := helper.TimeToLoc(sFdepositStartTime, loc) // 秒级时间戳
		if err != nil {
			return data, err
		}
		FdepositStartTime = t
	}
	var FdepositEndTime int64
	if len(sFdepositEndTime) > 0 {
		t, err := helper.TimeToLoc(sFdepositEndTime, loc) // 秒级时间戳
		if err != nil {
			return data, err
		}
		FdepositEndTime = t
	}

	var startTime int64
	if len(sStartTime) > 0 {
		t, err := helper.TimeToLoc(sStartTime, loc) // 秒级时间戳
		if err != nil {
			return data, err
		}
		startTime = t
	}
	var endTime int64
	if len(sEndTime) > 0 {
		t, err := helper.TimeToLoc(sEndTime, loc) // 秒级时间戳
		if err != nil {
			return data, err
		}
		endTime = t
	}
	data, err := reportMemberData(flag, dateFlag, timeFlag, page, pageSize, timeOutBet, timeOutLogin, timeOutDeposit, FdepositStartTime, FdepositEndTime,
		regStartTime, regEndTime, parentName, parentUid, userName, startTime, endTime, sCommissionId, sMainId, ty)
	if err != nil {
		return MemberReportData{}, err
	}

	return data, nil
}

func reportMemberData(flag, dateFlag, timeFlag, page, pageSize int, timeOutBet, timeOutLogin, timeOutDeposit int,
	FdepositStartTime, FdepositEndTime, regStartTime, regEndTime int64, parentName, parentUid, userName string, startTime,
	endTime int64, sCommissionId, sMainId string, ty int) (MemberReportData, error) {

	data := MemberReportData{}
	var list []MemberData
	var reportType int
	if flag == ReportFlagDay && dateFlag == ReportDateFlagBet {
		reportType = 1
	}

	if flag == ReportFlagDay && dateFlag == ReportDateFlagSettle {
		reportType = 2
	}
	if flag == ReportFlagMonth && dateFlag == ReportDateFlagBet {
		reportType = 3
	}

	if flag == ReportFlagMonth && dateFlag == ReportDateFlagSettle {
		reportType = 4
	}

	ex := g.Ex{
		"report_time":      g.Op{"between": exp.NewRangeVal(startTime, endTime)},
		"prefix":           meta.Prefix,
		"report_type":      reportType,
		"time_out_login":   g.Op{"gte": timeOutLogin},
		"time_out_bet":     g.Op{"gte": timeOutBet},
		"time_out_deposit": g.Op{"gte": timeOutDeposit},
	}

	if regStartTime != 0 && regEndTime != 0 {
		ex["created_at"] = g.Op{"between": exp.NewRangeVal(regStartTime, regEndTime)}
	}
	if FdepositStartTime != 0 && FdepositEndTime != 0 {
		ex["first_deposit_at"] = g.Op{"between": exp.NewRangeVal(FdepositStartTime, FdepositEndTime)}
	}
	if ty == 2 && len(parentName) == 0 && len(userName) == 0 {
		ex["parent_uid"] = "0"
	}

	if parentName != "" && len(parentName) > 0 {
		ex["parent_name"] = parentName
	}
	if sCommissionId != "" && len(sCommissionId) > 0 {
		ex["comission_id"] = sCommissionId
	}
	if parentName == "" && userName != "" && len(userName) > 0 {
		ex["username"] = userName
	}

	if page == 1 {
		if timeFlag != ReportTimeFlagSingle {
			var t []User
			totalQuery, _, _ := dialect.From("tbl_report_agency").Select(g.C("uid").As("uid"),
				g.V("0").As("report_time")).Where(ex).GroupBy("uid").ToSQL()
			if ty == 1 {
				totalQuery = strings.ReplaceAll(totalQuery, "WHERE", "WHERE uid=parent_uid and ")
			} else if userName != "" {
				totalQuery = strings.ReplaceAll(totalQuery, "WHERE", "WHERE uid!=parent_uid and ")
			}
			fmt.Println(totalQuery)
			err := meta.ReportDB.Select(&t, totalQuery)
			if err != nil {
				return data, pushLog(err, "db", DBErr)
			}
			data.T = int64(len(t))
			if data.T == 0 {
				return data, nil
			}
		} else {
			var t []User
			totalQuery, _, _ := dialect.From("tbl_report_agency").Select(g.C("uid").As("uid"),
				g.C("report_time").As("report_time")).Where(ex).ToSQL()
			if ty == 1 {
				totalQuery = strings.ReplaceAll(totalQuery, "WHERE", "WHERE uid=parent_uid and ")
			} else if userName != "" {
				totalQuery = strings.ReplaceAll(totalQuery, "WHERE", "WHERE uid!=parent_uid and ")
			}
			fmt.Println(totalQuery)
			err := meta.ReportDB.Select(&t, totalQuery)
			if err != nil {
				return data, pushLog(err, "db", DBErr)
			}
			data.T = int64(len(t))
			if data.T == 0 {
				return data, nil
			}
		}

		aggQuery, _, _ := dialect.From("tbl_report_agency").Select(g.MAX("time_out_bet").As("time_out_bet"),
			g.MAX("time_out_login").As("time_out_login"),
			g.MAX("time_out_deposit").As("time_out_deposit"),
			g.SUM("mem_count").As("mem_count"),
			g.SUM("subordinate_count").As("subordinate_count"),
			g.SUM("login_count").As("login_count"),
			g.SUM("regist_count").As("regist_count"),
			g.SUM("active_count").As("active_count"),
			g.SUM("conversion_rate").As("conversion_rate"),
			g.SUM("first_deposit_count").As("first_deposit_count"),
			g.SUM("first_deposit_amount").As("first_deposit_amount"),
			g.SUM("avg_first_deposit_amount").As("avg_first_deposit_amount"),
			g.SUM("deposit_mem_count").As("deposit_mem_count"),
			g.SUM("withdrawal_mem_count").As("withdrawal_mem_count"),
			g.SUM("deposit_amount").As("deposit_amount"),
			g.SUM("withdrawal_amount").As("withdrawal_amount"),
			g.SUM("adjust_amount").As("adjust_amount"),
			g.SUM("bet_mem_count").As("bet_mem_count"),
			g.SUM("bet_amount").As("bet_amount"),
			g.SUM("valid_bet_amount").As("valid_bet_amount"),
			g.SUM("company_net_amount").As("company_net_amount"),
			g.SUM("profit_amount").As("profit_amount"),
			g.SUM("dividend_amount").As("dividend_amount"),
			g.SUM("rebate_amount").As("rebate_amount"),
			g.SUM("agent_amount").As("agent_amount"),
			g.SUM("adjust_system_amount").As("adjust_system_amount"),
			g.SUM("adjust_win_amount").As("adjust_win_amount"),
			g.SUM("presettle").As("presettle"),
			g.SUM("company_revenue").As("company_revenue"),
			g.SUM("deposit_count").As("deposit_count")).Where(ex).ToSQL()
		err := meta.ReportDB.Get(&data.Agg, aggQuery)
		if ty == 1 {
			aggQuery = strings.ReplaceAll(aggQuery, "WHERE", "WHERE uid=parent_uid and ")
		} else if userName != "" {
			aggQuery = strings.ReplaceAll(aggQuery, "WHERE", "WHERE uid!=parent_uid and ")
		}
		if err != nil {
			return data, pushLog(fmt.Errorf("%s,[%s]", err.Error(), aggQuery), "db", helper.DBErr)
		}
		cr, _ := decimal.NewFromString(data.Agg.RegistCount)
		fdc, _ := decimal.NewFromString(data.Agg.FirstDepositCount)
		if cr.Cmp(decimal.Zero) != 0 {
			data.Agg.ConversionRate = fdc.Div(cr).StringFixed(4)
		} else {
			data.Agg.ConversionRate = "0.0000"
		}
		if fdc.Cmp(decimal.Zero) != 0 {
			fda, _ := decimal.NewFromString(data.Agg.FirstDepositAmount)
			data.Agg.AvgFirstDepositAmount = fda.Div(fdc).StringFixed(4)
		}
		cna, _ := decimal.NewFromString(data.Agg.CompanyNetAmount)
		ba, _ := decimal.NewFromString(data.Agg.BetAmount)
		if ba.Cmp(decimal.Zero) != 0 {
			data.Agg.ProfitAmount = cna.Div(ba).StringFixed(4)
		}
	}
	build := dialect.From("tbl_report_agency").Where(ex)
	offset := (page - 1) * pageSize
	build = build.Select(
		g.C("report_time").As("report_time"),
		g.C("uid").As("uid"),
		g.C("top_name").As("top_name"),
		g.C("top_uid").As("top_uid"),
		g.C("parent_name").As("parent_name"),
		g.C("parent_uid").As("parent_uid"),
		g.C("commissions_id").As("commissions_id"),
		g.C("commission_name").As("commission_name"),
		g.C("username").As("username"),
		g.C("prefix").As("prefix"),
		g.C("created_at").As("created_at"),
		g.MAX("time_out_bet").As("time_out_bet"),
		g.MAX("time_out_login").As("time_out_login"),
		g.MAX("time_out_deposit").As("time_out_deposit"),
		g.SUM("mem_count").As("mem_count"),
		g.SUM("subordinate_count").As("subordinate_count"),
		g.SUM("login_count").As("login_count"),
		g.SUM("regist_count").As("regist_count"),
		g.SUM("active_count").As("active_count"),
		g.SUM("conversion_rate").As("conversion_rate"),
		g.SUM("first_deposit_count").As("first_deposit_count"),
		g.SUM("first_deposit_amount").As("first_deposit_amount"),
		g.SUM("avg_first_deposit_amount").As("avg_first_deposit_amount"),
		g.SUM("deposit_mem_count").As("deposit_mem_count"),
		g.SUM("withdrawal_mem_count").As("withdrawal_mem_count"),
		g.SUM("deposit_amount").As("deposit_amount"),
		g.SUM("withdrawal_amount").As("withdrawal_amount"),
		g.SUM("adjust_amount").As("adjust_amount"),
		g.SUM("bet_mem_count").As("bet_mem_count"),
		g.SUM("bet_amount").As("bet_amount"),
		g.SUM("valid_bet_amount").As("valid_bet_amount"),
		g.SUM("company_net_amount").As("company_net_amount"),
		g.SUM("profit_amount").As("profit_amount"),
		g.SUM("dividend_amount").As("dividend_amount"),
		g.SUM("rebate_amount").As("rebate_amount"),
		g.SUM("agent_amount").As("agent_amount"),
		g.SUM("adjust_system_amount").As("adjust_system_amount"),
		g.SUM("adjust_win_amount").As("adjust_win_amount"),
		g.SUM("presettle").As("presettle"),
		g.SUM("company_revenue").As("company_revenue"),
		g.SUM("deposit_count").As("deposit_count"),
	).GroupBy("uid", "username", "prefix", "parent_uid", "parent_name", "commissions_id", "commission_name", "report_time").Offset(uint(offset)).Limit(uint(pageSize))
	query, _, _ := build.ToSQL()
	if ty == 1 {
		query = strings.ReplaceAll(query, "WHERE", "WHERE uid=parent_uid and ")
		query = strings.ReplaceAll(query, ", `report_time`", "")
	} else if userName != "" {
		query = strings.ReplaceAll(query, "WHERE", "WHERE uid!=parent_uid and ")
	}
	fmt.Println(query)
	err := meta.ReportDB.Select(&list, query)
	if err != nil {
		return data, pushLog(err, "db", DBErr)
	}
	var (
		uids   []interface{}
		suids  []string
		unames []string
	)
	for _, v := range list {
		member := MemberReportItem{
			ReportTime:            parseDay(v.ReportTime),
			Id:                    v.Id,
			TopName:               v.TopName,
			TopUid:                v.TopUid,
			ParentUid:             v.ParentUid,
			ParentName:            v.ParentName,
			CommissionsId:         v.CommissionsId,
			CommissionName:        v.CommissionName,
			MemCount:              v.MemCount,
			SubordinateCount:      v.SubordinateCount,
			LoginCount:            v.LoginCount,
			RegistCount:           v.RegistCount,
			ActiveCount:           v.ActiveCount,
			ConversionRate:        v.ConversionRate,
			FirstDepositCount:     v.FirstDepositCount,
			FirstDepositAmount:    v.FirstDepositAmount,
			AvgFirstDepositAmount: v.AvgFirstDepositAmount,
			DepositMemCount:       v.DepositMemCount,
			WithdrawalMemCount:    v.WithdrawalMemCount,
			DepositAmount:         v.DepositAmount,
			WithdrawalAmount:      v.WithdrawalAmount,
			AdjustAmount:          v.AdjustAmount,
			BetMemCount:           v.BetMemCount,
			BetAmount:             v.BetAmount,
			ValidBetAmount:        v.ValidBetAmount,
			CompanyNetAmount:      v.CompanyNetAmount,
			ProfitAmount:          v.ProfitAmount,
			DividendAmount:        v.DividendAmount,
			RebateAmount:          v.RebateAmount,
			AgentAmount:           v.AgentAmount,
			AdjustSystemAmount:    v.AdjustSystemAmount,
			AdjustWinAmount:       v.AdjustWinAmount,
			Prefix:                v.Prefix,
			Presettle:             v.Presettle,
			CompanyRevenue:        v.CompanyRevenue,
			Uid:                   v.Uid,
			DepositCount:          v.DepositCount,
			CreatedAt:             v.CreatedAt,
			FirstDepositAt:        v.FirstDepositAt,
			TimeOutDeposit:        v.TimeOutDeposit,
			TimeOutBet:            v.TimeOutBet,
			Username:              v.Username,
			TimeOutLogin:          v.TimeOutLogin,
		}
		if ty == 2 && parentName == member.Username {
			member.ParentName = ""
		}
		data.D = append(data.D, MemberReportShow{MemberReportItem: member})
		unames = append(unames, v.Username)
		uids = append(uids, v.Uid)
		suids = append(suids, v.Uid)
	}

	mbs, _ := MemberMCache(unames)
	for k, v := range data.D {
		// 获取会员代理信息
		if mb, ok := mbs[v.Username]; ok {
			data.D[k].State = strconv.Itoa(mb.State)
			data.D[k].LastLoginDevice = strconv.Itoa(mb.LastLoginSource)
			data.D[k].FirstBetAmount = strconv.FormatFloat(mb.FirstBetAmount, 'f', -1, 64)
			data.D[k].FirstBetAt = strconv.FormatInt(mb.FirstBetAt, 10)
			data.D[k].RegUrl = mb.RegUrl
			data.D[k].Balance = strconv.FormatFloat(mb.Balance, 'f', -1, 64)
			data.D[k].LastLoginAt = strconv.FormatInt(mb.LastLoginAt, 10)
		}
	}

	return data, nil
}

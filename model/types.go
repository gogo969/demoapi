package model

//报表查询时间类型
const (
	ReportDateFlagBet    = 1 //投注时间
	ReportDateFlagSettle = 2 //结算时间
)

//报表统计方式(单天或时间段)
const (
	ReportTimeFlagSingle = 1 //单天或单月
	ReportTimeFlagPart   = 2 //时间段
)

//日报或月报
const (
	ReportFlagDay   = 1 //日报
	ReportFlagMonth = 2 //月报
)

//报表统计类型游戏或场馆
const (
	ReportTyGame = 1
	ReportTyPlat = 2
)

type Member struct {
	Uid                int64   `json:"uid" db:"uid"`
	Username           string  `json:"username" db:"username"`
	Password           string  `json:"password" db:"password"`
	Prefix             string  `json:"prefix" db:"prefix"`
	Regip              string  `json:"regip" db:"regip"`
	RegDevice          string  `json:"reg_device" db:"reg_device"`
	RegUrl             string  `json:"reg_url" db:"reg_url"`
	CreatedAt          int64   `json:"created_at" db:"created_at"`
	LastLoginIp        string  `json:"last_login_ip" db:"last_login_ip"`
	LastLoginAt        int64   `json:"last_login_at" db:"last_login_at"`
	SourceId           int64   `json:"source_id" db:"source_id"`
	FirstBetAt         int64   `json:"first_bet_at" db:"first_bet_at"`
	FirstBetAmount     float64 `json:"first_bet_amount" db:"first_bet_amount"`
	FirstDepositAt     int64   `json:"first_deposit_at" db:"first_deposit_at"`
	FirstDepositAmount float64 `json:"first_deposit_amount" db:"first_deposit_amount"`
	TopUid             int     `json:"top_uid" db:"top_uid"`
	TopName            string  `json:"top_name" db:"top_name"`
	ParentUid          int64   `json:"parent_uid" db:"parent_uid"`
	ParentName         string  `json:"parent_name" db:"parent_name"`
	BankcardTotal      int     `json:"bankcard_total" db:"bankcard_total"`
	LastLoginDevice    string  `json:"last_login_device" db:"last_login_device"`
	LastLoginSource    int     `json:"last_login_source" db:"last_login_source"`
	Remarks            string  `json:"remarks" db:"remarks"`
	Balance            float64 `json:"balance" db:"balance"`
	LockAmount         float64 `json:"lock_amount" db:"lock_amount"`
	Commission         float64 `json:"commission" db:"commission"`
	State              int     `json:"state" db:"state"`
}

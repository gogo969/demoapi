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

type gameName struct {
	CnName string
	VnName string
}

var gameNameMap = map[int64]map[string]gameName{
	8840968486572375835: {

		"1": {VnName: "Hà Nội",
			CnName: "河内"},

		"2": {VnName: "Quảng Ninh",
			CnName: "广宁"},

		"3": {VnName: "Bắc Ninh",
			CnName: "北宁"},

		"4": {VnName: "Hải Phòng",
			CnName: "海防"},

		"5": {VnName: "Nam Định",
			CnName: "南定"},

		"6": {VnName: "Thái Bình",
			CnName: "太平"},

		"7": {VnName: "Thừa Thiên Huế",
			CnName: "顺化"},

		"8": {VnName: "Phú Yên",
			CnName: "富安"},

		"9": {VnName: "Tỉnh Đắk Lắk",
			CnName: "多乐"},

		"10": {VnName: "Quảng Nam",
			CnName: "广南"},

		"11": {VnName: "Khánh Hoà",
			CnName: "庆和"},

		"12": {VnName: "Đà Nẵng",
			CnName: "岘港"},

		"13": {VnName: "Bình Định",
			CnName: "平定"},

		"14": {VnName: "Quảng Bình",
			CnName: "广平"},

		"15": {VnName: "Quảng Trị",
			CnName: "广治"},

		"16": {VnName: "Ninh Thuận",
			CnName: "宁顺"},

		"17": {VnName: "Gia Lai",
			CnName: "嘉莱"},

		"18": {VnName: "Đak Nông",
			CnName: "达农"},

		"20": {VnName: "Khánh Hòa",
			CnName: "庆和"},

		"21": {VnName: "Kon Tum",
			CnName: "昆嵩"},

		"22": {VnName: "Đồng Tháp",
			CnName: "同塔"},

		"23": {VnName: "HỒ CHÍ MINH",
			CnName: "胡志明市"},

		"24": {VnName: "Cà Mau",
			CnName: "金瓯"},

		"25": {VnName: "Vũng Tàu",
			CnName: "头顿"},

		"26": {VnName: "Bến Tre",
			CnName: "槟知"},

		"27": {VnName: "Bạc Liêu",
			CnName: "薄寮"},

		"28": {VnName: "Cần Thơ",
			CnName: "芹苴"},

		"29": {VnName: "Sóc Trăng",
			CnName: "朔庄"},

		"30": {VnName: "Đồng Nai",
			CnName: "同奈"},

		"31": {VnName: "An Giang",
			CnName: "安江"},

		"32": {VnName: "Tây Ninh",
			CnName: "西宁"},

		"33": {VnName: "Bình Thuận",
			CnName: "平顺"},

		"34": {VnName: "Vĩnh Long",
			CnName: "永隆"},

		"35": {VnName: "Bình Dương",
			CnName: "平阳"},

		"36": {VnName: "Trà Vinh",
			CnName: "茶荣"},

		"37": {VnName: "Long An",
			CnName: "隆安"},

		"38": {VnName: "Bình Phước",
			CnName: "平福"},

		"39": {VnName: "Hậu Giang",
			CnName: "后江"},

		"40": {VnName: "Kiên Giang",
			CnName: "坚江"},

		"41": {VnName: "Tiền Giang",
			CnName: "前江"},

		"42": {VnName: "Đà Lạt",
			CnName: "大叻"},

		"45": {VnName: "Quảng Ngãi",
			CnName: "广义"},

		"47": {VnName: "Xổ số nhanh Hà Nội",
			CnName: "河内极速彩"},

		"48": {VnName: "Hà Nội 1 phút",
			CnName: "河内1分彩"},

		"49": {VnName: "Hà Nội 3 phút",
			CnName: "河内3分彩"},

		"50": {VnName: "Hà Nội 5 phút",
			CnName: "河内5分彩"},

		"51": {VnName: "Xổ số nhanh HỒ CHÍ MINH",
			CnName: "胡志明极速彩"},

		"52": {VnName: "HỒ CHÍ MINH 1 phút",
			CnName: "胡志明1分彩"},

		"53": {VnName: "HỒ CHÍ MINH 3 phút",
			CnName: "胡志明3分彩"},

		"54": {VnName: "HỒ CHÍ MINH 5 phút",
			CnName: "胡志明5分彩"},

		"55": {VnName: "Xổ số nhanh Đà Nẵng",
			CnName: "岘港极速彩"},

		"56": {VnName: "Đà Nẵng 1 phút",
			CnName: "岘港1分彩"},

		"57": {VnName: "Đà Nẵng 3 phút",
			CnName: "岘港3分彩"},

		"58": {VnName: "Đà Nẵng 5 phút",
			CnName: "岘港5分彩"},
	},
}

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
	Uid                string  `json:"uid" db:"uid"`
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

var gameNameMap = map[string]map[string]gameName{
	"8840968486572375835": {

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
	"2326854765648775667": {
		"AHK3": {
			CnName: "安徽快3",
			VnName: "An Huy nhanh 3"},
		"JLK3": {
			CnName: "吉林快3",
			VnName: "Cát Lâm nhanh 3",
		}, "JSK3": {
			CnName: "江苏快3",
			VnName: "Giang Tô nhanh 3",
		}, "HBK3": {
			CnName: "湖北快3",
			VnName: "Hồ Bắc nhanh 3",
		}, "NYK35F": {
			CnName: "新3D快三",
			VnName: "Mới 3D nhanh 3",
		}, "NYK35F2": {
			CnName: "澳门快三",
			VnName: "Ma Cao nhanh 3",
		}, "NYK35F3": {
			CnName: "凤凰快三",
			VnName: "Phoenix nhanh 3",
		}, "NYK35F4": {
			CnName: "新加坡快三5分",
			VnName: "Singapore 5 phút",
		}, "NYK35F5": {
			CnName: "香港快三",
			VnName: "Hồng Kông nhanh 3",
		}, "NYK32F": {
			CnName: "澳门快三2分",
			VnName: "Ma Cao 2 phút",
		}, "NYK310F": {
			CnName: "新加坡快三",
			VnName: "Singapore nhanh 3",
		}, "NYK310F2": {
			CnName: "台湾风彩",
			VnName: "Đài Loan nhanh 3",
		}, "NYK31F": {
			CnName: "极速骰宝",
			VnName: "Tốc độ sicbo",
		}, "NYK33F": {
			CnName: "幸运骰宝",
			VnName: "may mắn sicbo",
		}, "TXFFC": {
			CnName: "騰訊分分彩",
			VnName: "Tencent SSC",
		}, "TX3FC": {
			CnName: "騰訊三分彩",
			VnName: "Tencent SSC 3 phút",
		}, "TX5FC": {
			CnName: "騰訊五分彩",
			VnName: "Tencent SSC 5 phút",
		}, "TX10FC": {
			CnName: "騰訊十分彩",
			VnName: "Tencent SSC 10min",
		}, "TW5FC": {
			CnName: "台灣五分彩",
			VnName: "Taiwan SSC",
		}, "HNFFC": {
			CnName: "河內分分彩",
			VnName: "Hanoi Quick1",
		}, "HN5FC": {
			CnName: "河內五分彩",
			VnName: "Hanoi Quick5",
		}, "AZXY5": {
			CnName: "澳洲幸運5",
			VnName: "Lucky 5 Ball",
		}, "NYSSC30S": {
			CnName: "浙江30秒",
			VnName: "Zhejiang SSC 30s",
		}, "NYSSC1F": {
			CnName: "新腾讯分分彩",
			VnName: "New Tencent SSC",
		}, "NYSSC15F": {
			CnName: "上海1.5分彩",
			VnName: "Shanghai SSC 90s",
		}, "NYSSC3F": {
			CnName: "重庆3分彩",
			VnName: "Chongqing SSC 3min",
		}, "NYSSC3F2": {
			CnName: "幸运3分彩",
			VnName: "Lucky SSC",
		}, "NYSSC3F3": {
			CnName: "3分时时彩",
			VnName: "AE SSC 3min",
		}, "NYSSC5F": {
			CnName: "广东5分彩",
			VnName: "Guangdong SSC 5min",
		}, "NYSSC5F2": {
			CnName: "江西5分彩",
			VnName: "Jiangxi SSC 5min",
		}, "NYSSC5F3": {
			CnName: "重庆5分彩",
			VnName: "Chongqing SSC 5min",
		}, "NYSSC5F4": {
			CnName: "5分时时彩",
			VnName: "AE SSC 5min",
		}, "NYSSC10F": {
			CnName: "十分时时彩",
			VnName: "AE SSC 10min",
		}, "BJPK10": {
			CnName: "北京賽車",
			VnName: "Beijing PK10",
		}, "XYFT": {
			CnName: "幸運飛艇",
			VnName: "Lucky Airship",
		}, "AZXY10": {
			CnName: "澳洲幸運10",
			VnName: "Lucky 10 Ball",
		}, "NYSC30S": {
			CnName: "空战风云",
			VnName: "Air Wars",
		}, "NYSC1F": {
			CnName: "空战风云60秒",
			VnName: "Air Wars 60s",
		}, "NYSC75S": {
			CnName: "极速赛车",
			VnName: "Speed PK10",
		}, "NYSC15F": {
			CnName: "竞速1.5分",
			VnName: "Fast PK10",
		}, "NYSC3F": {
			CnName: "竞速3分",
			VnName: "Quick PK10",
		}, "NYSC5F": {
			CnName: "墨西哥摩托",
			VnName: "Mexico PK10",
		}, "NYSC5F2": {
			CnName: "澳门PK10",
			VnName: "Macao PK10",
		}, "NYSC5F3": {
			CnName: "香港赛车",
			VnName: "Hong Kong PK10",
		}, "SDSYXW": {
			CnName: "山東11選5",
			VnName: "Shandong 11x5",
		}, "JXSYXW": {
			CnName: "江西11選5",
			VnName: "Jiangxi 11x5",
		}, "GDSYXW": {
			CnName: "廣東11選5",
			VnName: "Guangdong 11x5",
		}, "JSSYXW": {
			CnName: "江苏11选5",
			VnName: "Jiangsu 11x5",
		}, "SHSYXW": {
			CnName: "上海11选5",
			VnName: "Shanghai 11x5",
		}, "BJ28": {
			CnName: "北京28",
			VnName: "Beijing 28",
		}, "XGLHC": {
			CnName: "香港六合彩",
			VnName: "HG MarkSix",
		}, "NYLHC75S": {
			CnName: "極速六合彩",
			VnName: "Speed MarkSix",
		}, "NYLHC3F": {
			CnName: "3分六合彩",
			VnName: "MarkSix 3min",
		}, "NYLHC5F": {
			CnName: "5分六合彩",
			VnName: "MarkSix 5min",
		}, "THAI": {
			CnName: "泰国官彩",
			VnName: "Thai Gov. Lottery",
		}, "NYTHAIFFC": {
			CnName: "泰国快乐彩",
			VnName: "Thai Happy Lottery",
		}, "NYTHAI3FC": {
			CnName: "泰国金币彩",
			VnName: "Thai Money Lottery",
		}, "NYTHAI5FC": {
			CnName: "泰国幸运彩",
			VnName: "Thai Lucky Lottery",
		}, "NYTHAI10FC": {
			CnName: "泰国闪耀彩",
			VnName: "Thai Shiny Lottery",
		}, "LAODL": {
			CnName: "老挝彩",
			VnName: "LAO Gov. Lottery",
		}, "MAGNUM4D": {
			CnName: "Malaysia Magnum 4D",
			VnName: "4D",
		}, "DAMA4D": {
			CnName: "大马彩4D",
			VnName: "Malaysia DaMaCai 4D",
		}, "SIG4D": {
			CnName: "新加坡4D",
			VnName: "Singapore 4D",
		}, "YEEKEE": {
			CnName: "YeeKee",
			VnName: "YeeKee YeeKee",
		}, "SUPERYK": {
			CnName: "Super",
			VnName: "Happy YeeKee",
		}, "HAPPYYK": {
			CnName: "Happy",
			VnName: "Happy YeeKee",
		}, "THAISETL": {
			CnName: "泰國SET綜合指數",
			VnName: "Thai stock market Noon",
		}, "THAISETN": {
			CnName: "泰國SET綜合指數",
			VnName: "Thai stock market Evening",
		}, "JPNIKL": {
			CnName: "日經225指數",
			VnName: "Nikkei 225 Morning",
		}, "JPNIKN": {
			CnName: "日經225指數",
			VnName: "Nikkei 225 Noon",
		}, "HKHSIL": {
			CnName: "香港恆生指數",
			VnName: "Hang Seng Index Morning",
		}, "HKHSIN": {
			CnName: "香港恆生指數",
			VnName: "Hang Seng Index Noon",
		}, "SZSECIL": {
			CnName: "深证成指",
			VnName: "Chinese stocks Morning",
		}, "SZSECIN": {
			CnName: "深证成指",
			VnName: "Chinese stocks Noon",
		},
		"NYYK1F": {
			CnName: "YeeKee",
			VnName: "YeeKee Lotto",
		},
	},
}

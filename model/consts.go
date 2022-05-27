package model

const (
	TransactionTypeAgAddAutoQuota  = 51 //代理自动刷新额度增加
	TransactionTypeAgAddFixedQuota = 52 //代理固定额度增加
	TransactionTypeAgAddGrandQuota = 53 //代理累计额度增加
	TransactionTypeAgDivQuota      = 54 //代理额度扣除
	TransactionTypeMbAddAutoQuota  = 55 //会员自动刷新额度增加
	TransactionTypeMbFixedQuota    = 56 //会员固定额度增加
	TransactionTypeMbGrandQuota    = 57 //会员累计额度增加
	TransactionTypeMbDivQuota      = 58 //会员额度扣除
	TransactionBet                 = 59 //投注
	TransactionBetCancel           = 60 //投注取消
	TransactionPayout              = 61 //派彩
	TransactionResettlePlus        = 62 //重新结算加币
	TransactionResettleDeduction   = 63 //重新结算减币
	TransactionCancelPayout        = 64 //取消派彩
	TransactionEVOPrize            = 65 //游戏奖金(EVO)
	TransactionEVOPromote          = 66 //推广(EVO)
	TransactionEVOJackpot          = 67 //头奖(EVO)
)

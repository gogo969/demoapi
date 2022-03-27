package model

import (
	"github.com/shopspring/decimal"
)

const (
	TypeCsv  = 1
	TypeXlsx = 2
)

func exportConvertAmount(sAmount string) string {

	amount, _ := decimal.NewFromString(sAmount)
	amount = amount.Mul(decimal.NewFromInt(1000))

	return amount.String()
}

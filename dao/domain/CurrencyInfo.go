package domain

import (
	"encoding/json"
	"unsafe"

	"github.com/shopspring/decimal"
)

/**
* 货币信息
 */
type CurrencyInfo struct {
	BaseModel

	CurrencyCode string          `json:"currency_code"`
	CurrencyType int             `json:"currency_type"`
	CurrencyName string          `json:"currency_name"`
	NominalValue decimal.Decimal `json:"nominal_value"`
	WeightInGram decimal.Decimal `json:"weight_in_gram"`
}

func (c *CurrencyInfo) String() string {
	byteArray, err := json.Marshal(c)
	if err == nil {
		return *(*string)(unsafe.Pointer(&byteArray))
	}
	return ""
}

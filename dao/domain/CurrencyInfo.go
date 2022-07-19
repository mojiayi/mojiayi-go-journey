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

	CurrencyCode string          `json:"currencyCode"`
	CurrencyType int             `json:"currencyType"`
	CurrencyName string          `json:"currencyName"`
	NominalValue decimal.Decimal `json:"nominalValue"`
	WeightInGram decimal.Decimal `json:"weightInGram"`
}

func (c *CurrencyInfo) String() string {
	byteArray, err := json.Marshal(c)
	if err == nil {
		return *(*string)(unsafe.Pointer(&byteArray))
	}
	return ""
}

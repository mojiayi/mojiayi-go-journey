package domain

import (
	"encoding/json"
	"time"
	"unsafe"
)

/**
* 外汇牌价信息
 */
type ForexPrice struct {
	BaseModel

	SrcCurrencyCode  string    `json:"srcCurrencyCode"`
	DestCurrencyCode string    `json:"destCurrencyCode"`
	BasePrice        float64   `json:"basePrice"`
	ExchangeDate     time.Time `json:"exchangeDate"`
}

func (c *ForexPrice) String() string {
	byteArray, err := json.Marshal(c)
	if err == nil {
		return *(*string)(unsafe.Pointer(&byteArray))
	}
	return ""
}

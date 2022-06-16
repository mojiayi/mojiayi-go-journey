package vo

import (
	"encoding/json"
	"unsafe"

	"github.com/shopspring/decimal"
)

/**
* 纸币信息
 */
type CurrencyInfoVO struct {
	/**
	* 主键id
	 */
	ID int
	/**
	* 货币代号
	 */
	CurrencyCode string
	/**
	* 货币类型，具体码值查看 CurrencyType
	 */
	CurrencyType int
	/**
	 * 货币名称，比如rmb-人民币，usd-美元
	 */
	CurrencyName string
	/**
	 * 货币单位，比如100表示100元一张，50表示50元一张
	 */
	NominalValue decimal.Decimal
	/**
	* 每张纸币或每枚硬币的重量，单位为 克
	 */
	WeightInGram decimal.Decimal
}

func (c *CurrencyInfoVO) String() string {
	byteArray, err := json.Marshal(c)
	if err == nil {
		return *(*string)(unsafe.Pointer(&byteArray))
	}
	return ""
}

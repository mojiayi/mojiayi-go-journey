package vo

/**
* 纸币信息
 */
type CurrencyInfoVO struct {
	/**
	* 货币代号
	 */
	CurrencyCode string `json:"CurrencyCode"`
	/**
	* 货币名称，比如rmb-人民币，usd-美元
	 */
	CurrencyName string `json:"CurrencyName"`
	/**
	 * 货币单位，比如100表示100元一张，50表示50元一张
	 */
	NominalValue int `json:"NominalValue"`
}

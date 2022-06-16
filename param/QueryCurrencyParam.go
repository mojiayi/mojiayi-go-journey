package param

/**
* 分页查询货币信息的参数
 */
type QueryCurrencyParam struct {
	PageParam

	CurrencyCode string `json:"currency_code"`
}

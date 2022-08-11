package crawler

type ExchangeItem struct {
	ExchangeDate string  `json:"exchangeDate"`
	BasePrice    float64 `json:"basePrice"`
	CurrType     string  `json:"currType"`
}

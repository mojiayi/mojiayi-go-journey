package crawler

type ForexResponseBody struct {
	Data         ExchangeList `json:"data"`
	ResponseCode string       `json:"responseCode"`
	ResponseMsg  string       `json:"responseMsg"`
}

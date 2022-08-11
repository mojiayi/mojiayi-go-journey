package crawler

import (
	"encoding/json"
	"io"
	"mojiayi-go-journey/dao/domain"
	"mojiayi-go-journey/dao/mapper"
	"mojiayi-go-journey/setting"
	"net/http"
	"strings"
	"time"
)

var (
	forexPriceMapper = *new(mapper.ForexPriceMapper)
)

type ForexExchangeCrawler struct{}

func (f *ForexExchangeCrawler) GetLatestExchangePrice() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		currentDate := time.Now().Format("2006-01-02")

		url := "https://bank.pingan.com.cn/rmb/account/cmp/cust/acct/forex/exchange/qryFoexPriceExchangeList.do?pageIndex=1&pageSize=100&realFlag=1&currencyCode=&exchangeDate=" + currentDate + "&access_source=PC"

		response, err := http.Get(url)
		if err != nil {
			setting.MyLogger.Warn("抓取外汇牌价失败", err)
			break
		}

		body, err := io.ReadAll(response.Body)
		response.Body.Close()

		if response.StatusCode > 299 || err != nil {
			setting.MyLogger.Warn("提取返回结果失败,statusCode={}", response.StatusCode, err)
			break
		}
		setting.MyLogger.Info("body=", body)
		var forexResponseBody ForexResponseBody
		json.Unmarshal(body, &forexResponseBody)
		if strings.Compare("000000", forexResponseBody.ResponseCode) != 0 {
			break
		}
		if forexResponseBody.Data.Count <= 0 {
			break
		}

		for _, value := range forexResponseBody.Data.ExchangeList {
			var forexPrice = *new(domain.ForexPrice)
			forexPrice.BasePrice = value.BasePrice
			forexPrice.DestCurrencyCode = "RMB"
			forexPrice.SrcCurrencyCode = value.CurrType
			forexPrice.ExchangeDate, _ = time.Parse("2006-01-02", value.ExchangeDate)

			existForexPrice, err := forexPriceMapper.SelectByCurrencyCode(forexPrice.SrcCurrencyCode, forexPrice.DestCurrencyCode)
			if err == nil {
				// 已经存在指定货币兑换组合的记录时，如果牌价发生变化，就修改已有数据
				if forexPrice.BasePrice == existForexPrice.BasePrice {
					continue
				}
				forexPrice.ID = existForexPrice.ID
				forexPriceMapper.Modify(forexPrice)
			} else {
				// 不存在指定货币兑换组合的记录时，插入新数据
				forexPriceMapper.Insert(forexPrice)
			}
		}
	}
}

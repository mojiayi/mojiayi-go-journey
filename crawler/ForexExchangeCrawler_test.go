package crawler_test

import (
	"mojiayi-go-journey/crawler"
	"testing"
)

var (
	forexCrawler = *new(crawler.ForexExchangeCrawler)
)

func TestGetForexPrice(t *testing.T) {
	forexCrawler.GetLatestExchangePrice()
}

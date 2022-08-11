package main

import (
	"strconv"

	"mojiayi-go-journey/crawler"
	"mojiayi-go-journey/message"
	"mojiayi-go-journey/routers"
	"mojiayi-go-journey/setting"
)

var forexCrawler = *new(crawler.ForexExchangeCrawler)

func main() {
	setting.Setup()

	router := routers.InitRouter(setting.WebSetting.ContextPath)

	// 默认使用8080端口，配置文件中可以指定新端口
	addr := ":8080"
	if setting.WebSetting.Port >= 8080 {
		addr = ":" + strconv.Itoa(setting.WebSetting.Port)
	}

	message.Subscribe(setting.KafkaSetting.ConsumerGroup, setting.KafkaSetting.Topic)

	forexCrawler.GetLatestExchangePrice()

	err := router.Run(addr)
	if err != nil {
		setting.MyLogger.Info("启动失败,err=", err)
		return
	}
}

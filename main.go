package main

import (
	"fmt"
	"strconv"

	"mojiayi-go-journey/routers"
	"mojiayi-go-journey/setting"
)

func main() {
	setting.Setup()

	router := routers.InitRouter(setting.WebSetting.ContextPath)

	addr := ":8080"
	if setting.WebSetting.Port >= 8080 {
		addr = ":" + strconv.Itoa(setting.WebSetting.Port)
	}

	err := router.Run(addr)
	if err != nil {
		fmt.Printf("启动失败,err=%v", err)
	}
}

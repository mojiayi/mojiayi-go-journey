package utils

import (
	"mojiayi-go-journey/setting"

	"github.com/gin-gonic/gin"
)

type HeaderUtil struct {
}

func (h *HeaderUtil) GetHeaderValue(headerKey string, context *gin.Context) (headerValue string) {
	headerValue = context.Request.Header.Get(headerKey)
	if len(headerValue) == 0 {
		setting.MyLogger.Info("header中没有headerKey=", headerKey)
	} else {
		setting.MyLogger.Info("从header中取得headerKey=", headerKey, ",headerValue=", headerValue)
	}
	return headerValue
}

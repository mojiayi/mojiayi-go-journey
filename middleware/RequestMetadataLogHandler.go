package middleware

import (
	"encoding/json"
	"mojiayi-go-journey/setting"
	"time"
	"unsafe"

	"github.com/gin-gonic/gin"
)

type RequestMetadata struct {
	Cost   int64  `json:"cost"`
	IP     string `json:"ip"`
	Method string `json:"method"`
	URI    string `json:"uri"`
}

func (c *RequestMetadata) String() string {
	byteArray, err := json.Marshal(c)
	if err == nil {
		return *(*string)(unsafe.Pointer(&byteArray))
	}
	return ""
}

func RecordCostTime() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		startTime := time.Now()

		ctx.Next()

		var requestMetadata = new(RequestMetadata)
		requestMetadata.Cost = time.Since(startTime).Milliseconds()
		requestMetadata.IP = ctx.ClientIP()
		requestMetadata.Method = ctx.Request.Method
		requestMetadata.URI = ctx.Request.RequestURI

		setting.MetadataLogger.Info(requestMetadata)
	}
}

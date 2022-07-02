package middleware

import (
	"mojiayi-go-journey/setting"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func RecordCostTime() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		startTime := time.Now()

		ctx.Next()

		setting.MetadataLogger.WithFields(logrus.Fields{
			"cost":   time.Since(startTime).Milliseconds(),
			"ip":     ctx.ClientIP(),
			"method": ctx.Request.Method,
			"uri":    ctx.Request.RequestURI,
			"usage":  "metadata",
		}).Info("requestMetadata")
	}
}

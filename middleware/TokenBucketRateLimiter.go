package middleware

import (
	"mojiayi-go-journey/constants"
	"mojiayi-go-journey/setting"
	"time"

	"github.com/gin-gonic/gin"
)

/**
* 模拟一个接口在1秒内最多只能访问3次
 */
func CheckFrequency(ctx *gin.Context) {
	var uri = ctx.Request.RequestURI

	isAcquire := tryAcquire(uri)

	if isAcquire {
		ctx.Next()
		return
	}
	panic("访问太快了")
}

func tryAcquire(uri string) bool {
	// 以uri作为key，一个接口在指定时间内只能访问指定次数
	var key = constants.URI_PREFIX + uri

	// 从缓存查询上次的访问时间，默认为0
	lastRefillTime, err := setting.RedisClient.HGet(key, constants.LAST_REFILL_TIME).Int64()
	if err != nil {
		lastRefillTime = 0
	}

	var isAcquire = true

	// 如果没有查到上次访问时间，作为第一次访问处理
	if lastRefillTime == 0 {
		remainToken := setting.RateLimitSetting.Qps - 1
		setting.RedisClient.HSet(key, constants.LAST_REFILL_TIME, time.Now().UnixMilli())
		setting.RedisClient.HSet(key, constants.REMAIN_TOKEN, remainToken)
		return isAcquire
	}

	// 计算上次访问到现在的时间间隔
	difference := time.Since(time.UnixMilli(lastRefillTime)).Milliseconds()

	// 如果已经超出了指定的时间，作为第一次访问处理
	if difference >= setting.RateLimitSetting.Interval {
		remainToken := setting.RateLimitSetting.Qps - 1
		setting.RedisClient.HSet(key, constants.LAST_REFILL_TIME, time.Now().UnixMilli())
		setting.RedisClient.HSet(key, constants.REMAIN_TOKEN, remainToken)
		return isAcquire
	}

	// 从缓存查询截止上次访问时用掉的次数
	remainToken, err := setting.RedisClient.HGet(key, constants.REMAIN_TOKEN).Int()
	if err != nil {
		remainToken = setting.RateLimitSetting.Qps
	}

	remainToken = int(float64(difference)/float64(setting.RateLimitSetting.Interval)*float64(setting.RateLimitSetting.Qps)) + remainToken

	isAcquire = remainToken > 0

	if isAcquire {
		remainToken = remainToken - 1
	} else {
		remainToken = 0
	}
	if remainToken >= setting.RateLimitSetting.Qps {
		remainToken = setting.RateLimitSetting.Qps - 1
	}
	setting.RedisClient.HSet(key, constants.LAST_REFILL_TIME, time.Now().UnixMilli())
	setting.RedisClient.HSet(key, constants.REMAIN_TOKEN, remainToken)

	return isAcquire
}

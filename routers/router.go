package routers

import (
	middlewire "mojiayi-go-journey/middleware"
	"mojiayi-go-journey/routers/api"

	"github.com/gin-gonic/gin"
)

var (
	currencyInfoApi api.CurrencyInfoApi
)

func InitRouter(contextPath string) *gin.Engine {
	router := gin.New()

	router.Use(middlewire.HandleError)
	router.Use(middlewire.EnableTraceIdHook)
	router.Use(middlewire.RecordCostTime())
	router.Use(middlewire.CheckFrequency)

	currency := router.Group(contextPath + "/currency")
	{
		currency.POST("", currencyInfoApi.AddCurrency)
		currency.DELETE("/:currencyCode/:nominalValue", currencyInfoApi.DeleteCurrency)
		currency.PUT("", currencyInfoApi.ModifyCurrency)
		currency.GET("", currencyInfoApi.QueryAvailableCurrency)
		currency.GET("/:currencyCode/:nominalValue", currencyInfoApi.QuerySpecifiedCurrency)
	}

	return router
}

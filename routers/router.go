package routers

import (
	"mojiayi-go-journey/routers/api"

	"github.com/gin-gonic/gin"
)

var currencyInfoApi api.CurrencyInfoApi

func InitRouter(contextPath string) *gin.Engine {
	router := gin.New()

	currency := router.Group(contextPath + "/currency")
	{
		currency.POST("/", currencyInfoApi.AddCurrency)
		currency.DELETE("/:currencyCode", currencyInfoApi.DeleteCurrency)
		currency.PUT("/", currencyInfoApi.ModifyCurrency)
		currency.GET("/", currencyInfoApi.QueryAvailableCurrency)
		currency.GET("/:currencyCode", currencyInfoApi.QuerySpecifiedCurrency)
	}

	return router
}

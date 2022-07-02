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
	router.Use(middlewire.RecordCostTime())
	router.Use(middlewire.CheckFrequency)

	currency := router.Group(contextPath + "/currency")
	{
		currency.POST("/", currencyInfoApi.AddCurrency)
		currency.DELETE("/:CurrencyCode/:NominalValue", currencyInfoApi.DeleteCurrency)
		currency.PUT("/", currencyInfoApi.ModifyCurrency)
		currency.GET("/", currencyInfoApi.QueryAvailableCurrency)
		currency.GET("/:CurrencyCode/:NominalValue", currencyInfoApi.QuerySpecifiedCurrency)
	}

	return router
}

package routers

import (
	middlewire "mojiayi-go-journey/middleware"
	"mojiayi-go-journey/routers/api"
	"mojiayi-go-journey/setting"

	"github.com/gin-gonic/gin"
)

var (
	currencyInfoApi api.CurrencyInfoApi
)

func InitRouter(contextPath string) *gin.Engine {
	router := gin.New()

	router.Use(setting.PutTraceIdIntoLocalStorage())
	router.Use(middlewire.RecordCostTime())
	router.Use(middlewire.HandleError)

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

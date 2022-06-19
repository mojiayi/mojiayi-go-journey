package api

import (
	"mojiayi-go-journey/constants"
	"mojiayi-go-journey/param"
	"mojiayi-go-journey/service"
	"mojiayi-go-journey/setting"
	"mojiayi-go-journey/utils"
	"mojiayi-go-journey/vo"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

type CurrencyInfoApi struct {
	respUtil utils.RespUtil
}

var (
	currencyInfoService = *new(service.CurrencyInfoService)
	paginateUtil        = *new(utils.PaginateUtil)
)

/**
* 新增可用的货币信息，通过request body形式传入新数据
 */
func (c *CurrencyInfoApi) AddCurrency(ctx *gin.Context) {
	newCurrency := vo.CurrencyInfoVO{}
	err := ctx.BindJSON(&newCurrency)
	if err != nil {
		setting.MyLogger.Info("传入的新货币信息不正确,err=", err)
		c.respUtil.IllegalArgumentErrorResp("传入的新货币信息不正确", ctx)
		return
	}
	id, err := currencyInfoService.AddCurrency(newCurrency)
	if err != nil || id == constants.INT_ZERO {
		setting.MyLogger.Info("插入货币信息失败,newCurrency=" + newCurrency.String())
		c.respUtil.ErrorResp(http.StatusInternalServerError, err.Error(), ctx)
		return
	}
	c.respUtil.SuccessResp(id, ctx)
}

/**
* 删除货币信息，通过path variable的形式传入指定的货币代号
 */
func (c *CurrencyInfoApi) DeleteCurrency(ctx *gin.Context) {
	currencyCode := ctx.Param("CurrencyCode")
	nominalValueStr := ctx.Param("NominalValue")
	if len(currencyCode) == 0 || len(nominalValueStr) == 0 {
		panic("必须指定货币代号和面值")
	}
	nominalValue, err := decimal.NewFromString(nominalValueStr)
	if err != nil {
		panic(err.Error())
	}

	err = currencyInfoService.DeleteCurrency(currencyCode, nominalValue)

	if err != nil {
		setting.MyLogger.Info("删除货币" + currencyCode + "(" + nominalValueStr + ")失败")
		panic(err.Error())
	}
	c.respUtil.SuccessResp(true, ctx)
}

/**
* 修改可用的货币信息，通过request body形式传入新数据
 */
func (c *CurrencyInfoApi) ModifyCurrency(ctx *gin.Context) {
	newCurrency := vo.CurrencyInfoVO{}
	err := ctx.BindJSON(&newCurrency)
	if err != nil {
		setting.MyLogger.Info("传入的新货币信息不正确,err=", err)
		c.respUtil.IllegalArgumentErrorResp("传入的新货币信息不正确", ctx)
		return
	}

	err = currencyInfoService.ModifyCurrency(newCurrency)

	if err != nil {
		setting.MyLogger.Info("修改货币" + newCurrency.CurrencyCode + "(" + newCurrency.NominalValue.String() + ")失败")
		c.respUtil.ErrorResp(http.StatusInternalServerError, err.Error(), ctx)
		return
	}
	c.respUtil.SuccessResp(true, ctx)
}

/**
* 查询可用的货币信息，通过path variable的形式传入指定的货币代号
 */
func (c *CurrencyInfoApi) QuerySpecifiedCurrency(ctx *gin.Context) {
	currencyCode := ctx.Param("CurrencyCode")
	nominalValueStr := ctx.Param("NominalValue")
	if len(currencyCode) == 0 || len(nominalValueStr) == 0 {
		c.respUtil.IllegalArgumentErrorResp("必须指定货币代号和面值", ctx)
		return
	}
	nominalValue, err := decimal.NewFromString(nominalValueStr)
	if err != nil {
		c.respUtil.IllegalArgumentErrorResp("货币面值必须是数字", ctx)
		return
	}
	targetCurrency, err := currencyInfoService.QuerySpecifiedCurrency(currencyCode, nominalValue)

	if err != nil {
		setting.MyLogger.Info("要查询的货币不存在,CurrencyCode=", currencyCode)
		c.respUtil.ErrorResp(http.StatusNotFound, "货币不存在", ctx)
	} else {
		c.respUtil.SuccessResp(targetCurrency, ctx)
	}
}

/**
* 查询可用的货币信息，通过request parameter的形式传入指定的货币代号
 */
func (c *CurrencyInfoApi) QueryAvailableCurrency(ctx *gin.Context) {
	var param = *new(param.QueryCurrencyParam)
	currencyCode := ctx.Query("CurrencyCode")
	if len(currencyCode) > 0 {
		param.CurrencyCode = currencyCode
	}

	param.CurrentPage = paginateUtil.GetCurrentPage(ctx)
	param.PageSize = paginateUtil.GetPageSize(ctx)

	var targetCurrencyList = currencyInfoService.QueryAvailableCurrency(param)

	c.respUtil.SuccessResp(targetCurrencyList, ctx)
}

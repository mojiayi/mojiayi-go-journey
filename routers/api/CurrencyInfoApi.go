package api

import (
	"bytes"
	"container/list"
	"encoding/json"
	"mojiayi-go-journey/utils"
	"mojiayi-go-journey/vo"
	"strings"

	"github.com/gin-gonic/gin"
)

type CurrencyInfoApi struct {
	respUtil utils.RespUtil
}

var currencyList = list.New()

/**
* 新增可用的货币信息，通过request body形式传入新数据
 */
func (c *CurrencyInfoApi) AddCurrency(ctx *gin.Context) {
	newCurrency := vo.CurrencyInfoVO{}
	err := ctx.BindJSON(&newCurrency)
	if err != nil {
		c.respUtil.IllegalArgumentErrorResp("传入的新货币信息不正确", ctx)
		return
	}
	var isExist = false
	for currency := currencyList.Front(); currency != nil; currency = currency.Next() {
		var item = currency.Value.(*vo.CurrencyInfoVO)
		if strings.EqualFold(newCurrency.CurrencyCode, item.CurrencyCode) {
			isExist = true
			break
		}
	}
	if isExist {
		c.respUtil.IllegalArgumentErrorResp("要添加的货币已存在", ctx)
		return
	}
	currencyList.PushBack(&newCurrency)

	jsonRsult, _ := marshalList(currencyList)
	c.respUtil.SuccessResp(jsonRsult, ctx)
}

/**
* 删除货币信息，通过path variable的形式传入指定的货币代号
 */
func (c *CurrencyInfoApi) DeleteCurrency(ctx *gin.Context) {
	currencyCode := ctx.Param("currencyCode")
	if currencyList.Len() == 0 {
		c.respUtil.IllegalArgumentErrorResp("没有可删除的货币信息", ctx)
		return
	}
	var isExist = false
	for currency := currencyList.Front(); currency != nil; currency = currency.Next() {
		var item = currency.Value.(*vo.CurrencyInfoVO)
		if strings.EqualFold(currencyCode, item.CurrencyCode) {
			isExist = true
			currencyList.Remove(currency)
		}
	}

	if isExist {
		jsonRsult, _ := marshalList(currencyList)
		c.respUtil.SuccessResp(jsonRsult, ctx)
	} else {
		c.respUtil.IllegalArgumentErrorResp("要删除的货币不存在", ctx)
	}
}

/**
* 修改可用的货币信息，通过request body形式传入新数据
 */
func (c *CurrencyInfoApi) ModifyCurrency(ctx *gin.Context) {
	newCurrency := vo.CurrencyInfoVO{}
	err := ctx.BindJSON(&newCurrency)
	if err != nil {
		c.respUtil.IllegalArgumentErrorResp("传入的新货币信息不正确", ctx)
		return
	}

	var isExist = false
	for currency := currencyList.Front(); currency != nil; currency = currency.Next() {
		item := currency.Value.(*vo.CurrencyInfoVO)
		if strings.EqualFold(newCurrency.CurrencyCode, item.CurrencyCode) {
			isExist = true
			item.CurrencyName = newCurrency.CurrencyName
			item.NominalValue = newCurrency.NominalValue
		}
	}

	if isExist {
		jsonRsult, _ := marshalList(currencyList)
		c.respUtil.SuccessResp(jsonRsult, ctx)
	} else {
		c.respUtil.IllegalArgumentErrorResp("要修改的货币不存在", ctx)
	}
}

/**
* 查询可用的货币信息，通过path variable的形式传入指定的货币代号
 */
func (c *CurrencyInfoApi) QuerySpecifiedCurrency(ctx *gin.Context) {
	currencyCode := ctx.Param("currencyCode")
	var targetCurrency *vo.CurrencyInfoVO
	var isExist = false
	for currency := currencyList.Front(); currency != nil; currency = currency.Next() {
		var item = currency.Value.(*vo.CurrencyInfoVO)
		if strings.EqualFold(currencyCode, item.CurrencyCode) {
			targetCurrency = item
			isExist = true
			break
		}
	}

	if isExist {
		c.respUtil.SuccessResp(targetCurrency, ctx)
	} else {
		c.respUtil.IllegalArgumentErrorResp("要查询的货币不存在", ctx)
	}
}

/**
* 查询可用的货币信息，通过request parameter的形式传入指定的货币代号
 */
func (c *CurrencyInfoApi) QueryAvailableCurrency(ctx *gin.Context) {
	currencyCode := ctx.Query("currencyCode")
	if len(currencyCode) == 0 {
		jsonRsult, _ := marshalList(currencyList)
		c.respUtil.SuccessResp(jsonRsult, ctx)
		return
	}
	var targetCurrencyList = list.New()
	for currency := currencyList.Front(); currency != nil; currency = currency.Next() {
		var item = currency.Value.(*vo.CurrencyInfoVO)
		if strings.EqualFold(currencyCode, item.CurrencyCode) {
			targetCurrencyList.PushBack(item)
		}
	}

	jsonRsult, _ := marshalList(targetCurrencyList)
	c.respUtil.SuccessResp(jsonRsult, ctx)
}

func marshalList(list *list.List) (interface{}, error) {
	buffer := bytes.NewBufferString("[")

	for item := list.Front(); item != nil; item = item.Next() {
		marshalled, err := json.Marshal(item.Value)

		if err != nil {
			return "[]", err
		}

		buffer.WriteString(string(marshalled))

		if item.Next() != nil {
			buffer.WriteRune(',')
		}
	}

	buffer.WriteString("]")

	var jsonRsult interface{}
	json.Unmarshal(buffer.Bytes(), &jsonRsult)
	return jsonRsult, nil
}

package service

import (
	"errors"
	"mojiayi-go-journey/constants"
	"mojiayi-go-journey/dao/domain"
	"mojiayi-go-journey/dao/mapper"
	"mojiayi-go-journey/param"
	"mojiayi-go-journey/setting"
	"mojiayi-go-journey/vo"

	"github.com/shopspring/decimal"
)

type CurrencyInfoService struct {
}

var (
	currencyInfoMapper = *new(mapper.CurrencyInfoMapper)
)

func (c *CurrencyInfoService) AddCurrency(currencyInfoVO vo.CurrencyInfoVO) (int, error) {
	existCurrencyInfo, _ := c.QuerySpecifiedCurrency(currencyInfoVO.CurrencyCode, currencyInfoVO.NominalValue)
	if existCurrencyInfo != nil {
		return constants.INT_ZERO, errors.New("货币" + currencyInfoVO.CurrencyCode + "(" + currencyInfoVO.NominalValue.String() + ")已存在")
	}

	po := *new(domain.CurrencyInfo)
	po.CurrencyCode = currencyInfoVO.CurrencyCode
	po.CurrencyType = currencyInfoVO.CurrencyType
	po.CurrencyName = currencyInfoVO.CurrencyName
	po.NominalValue = currencyInfoVO.NominalValue
	po.WeightInGram = currencyInfoVO.WeightInGram

	id, err := currencyInfoMapper.Insert(po)
	if err != nil || id == constants.INT_ZERO {
		setting.MyLogger.Info("插入货币信息失败,po=" + po.String())
	}

	return id, err
}

func (c *CurrencyInfoService) DeleteCurrency(currencyCode string, nominalValue decimal.Decimal) error {
	existCurrencyInfo, _ := c.QuerySpecifiedCurrency(currencyCode, nominalValue)
	if existCurrencyInfo == nil {
		return errors.New("货币" + currencyCode + "(" + nominalValue.String() + ")不存在")
	}

	effectRow := currencyInfoMapper.DeleteCurrency(existCurrencyInfo.ID)
	if effectRow == constants.INT_ONE {
		return nil
	}
	return errors.New("删除货币" + currencyCode + "(" + nominalValue.String() + ")失败")
}

func (c *CurrencyInfoService) ModifyCurrency(currencyInfoVO vo.CurrencyInfoVO) error {
	existCurrencyInfo, _ := c.QueryCurrencyById(currencyInfoVO.ID)
	if existCurrencyInfo == nil {
		return errors.New("货币" + currencyInfoVO.CurrencyCode + "(" + currencyInfoVO.NominalValue.String() + ")不存在")
	}

	po := *new(domain.CurrencyInfo)
	po.ID = currencyInfoVO.ID
	po.CurrencyType = currencyInfoVO.CurrencyType
	po.CurrencyName = currencyInfoVO.CurrencyName
	po.NominalValue = currencyInfoVO.NominalValue
	po.WeightInGram = currencyInfoVO.WeightInGram

	effectRow := currencyInfoMapper.Modify(po)
	if effectRow == constants.INT_ONE {
		return nil
	}
	return errors.New("修改货币" + currencyInfoVO.CurrencyCode + "(" + currencyInfoVO.NominalValue.String() + ")失败")
}

func (c *CurrencyInfoService) QueryCurrencyById(id int) (*vo.CurrencyInfoVO, error) {
	currencyInfo, err := currencyInfoMapper.SelectById(id)
	if err != nil {
		return nil, err
	}

	return c.po2Vo(currencyInfo), nil
}

func (c *CurrencyInfoService) QuerySpecifiedCurrency(currencyCode string, nominalValue decimal.Decimal) (*vo.CurrencyInfoVO, error) {
	currencyInfo, err := currencyInfoMapper.SelectByCurrencyCode(currencyCode, nominalValue)
	if err != nil {
		return nil, err
	}

	return c.po2Vo(currencyInfo), nil
}

func (c *CurrencyInfoService) QueryAvailableCurrency(param param.QueryCurrencyParam) *vo.BasePageVO {
	pageResult := vo.BasePageVO{}
	pageResult.CurrentPage = param.CurrentPage
	pageResult.PageSize = param.PageSize

	currencyCode := param.CurrencyCode

	total := currencyInfoMapper.CountByCondition(currencyCode)
	pageResult.Total = total
	if total == 0 {
		pageResult.Pages = 0
		pageResult.List = []vo.CurrencyInfoVO{}
		return &pageResult
	}
	recordList, err := currencyInfoMapper.PageByCondition(&pageResult, currencyCode)
	if err != nil {
		pageResult.Pages = 0
		pageResult.List = []vo.CurrencyInfoVO{}
		return &pageResult
	}

	var currencyInfoList = make([]vo.CurrencyInfoVO, len(recordList))

	for index, value := range recordList {
		currencyInfoList[index] = *c.po2Vo(value)
	}

	setting.MyLogger.Info("返回货币信息条数,size=", len(currencyInfoList))

	(&pageResult).List = currencyInfoList
	return &pageResult
}

func (c CurrencyInfoService) po2Vo(po domain.CurrencyInfo) *vo.CurrencyInfoVO {
	vo := *new(vo.CurrencyInfoVO)
	vo.ID = po.ID
	vo.CurrencyCode = po.CurrencyCode
	vo.CurrencyType = po.CurrencyType
	vo.CurrencyName = po.CurrencyName
	vo.NominalValue = po.NominalValue
	vo.WeightInGram = po.WeightInGram
	return &vo
}

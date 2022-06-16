package mapper

import (
	"errors"
	"mojiayi-go-journey/constants"
	"mojiayi-go-journey/dao/domain"
	"mojiayi-go-journey/setting"
	"mojiayi-go-journey/utils"
	"mojiayi-go-journey/vo"
	"strconv"

	"github.com/shopspring/decimal"
)

type CurrencyInfoMapper struct {
}

var (
	paginateUtil utils.PaginateUtil
)

func (c *CurrencyInfoMapper) Insert(currencyInfo domain.CurrencyInfo) (int, error) {
	setting.DB.Create(&currencyInfo)

	if currencyInfo.ID > 0 {
		return currencyInfo.ID, nil
	}
	return constants.INT_ZERO, errors.New("创建货币（" + currencyInfo.CurrencyName + ":" + currencyInfo.NominalValue.String() + "）信息失败")
}

func (c *CurrencyInfoMapper) DeleteCurrency(id int) int {
	wrapper := make(map[string]interface{})
	wrapper["id"] = id
	setting.DB.Model(&domain.CurrencyInfo{}).Delete(wrapper)
	return constants.INT_ONE
}

func (c *CurrencyInfoMapper) Modify(currencyInfo domain.CurrencyInfo) int {
	wrapper := make(map[string]interface{})
	wrapper["id"] = currencyInfo.ID
	setting.DB.Model(&domain.CurrencyInfo{}).Where(wrapper).Updates(currencyInfo)

	return constants.INT_ONE
}

func (c *CurrencyInfoMapper) SelectById(id int) (currencyInfo domain.CurrencyInfo, err error) {
	wrapper := make(map[string]interface{})
	wrapper["id"] = id

	var record domain.CurrencyInfo
	setting.DB.Model(&domain.CurrencyInfo{}).Where(wrapper).Find(&record)

	if record.ID == 0 {
		return record, errors.New("货币(id=" + strconv.Itoa(id) + ")不存在")
	}

	return record, nil
}

func (c *CurrencyInfoMapper) SelectByCurrencyCode(currencyCode string, nominalValue decimal.Decimal) (currencyInfo domain.CurrencyInfo, err error) {
	wrapper := make(map[string]interface{})
	wrapper["currency_code"] = currencyCode
	wrapper["nominal_value"] = nominalValue

	var record domain.CurrencyInfo
	setting.DB.Model(&domain.CurrencyInfo{}).Where(wrapper).Find(&record)

	if record.ID == 0 {
		return record, errors.New("货币" + currencyCode + "(" + nominalValue.String() + ")不存在")
	}

	return record, nil
}

func (c *CurrencyInfoMapper) CountByCondition(currencyCode string) int {
	wrapper := make(map[string]interface{}, 0)
	if currencyCode != "" {
		wrapper["currency_code"] = currencyCode
	}

	var total int64
	setting.DB.Model(&domain.CurrencyInfo{}).Where(wrapper).Count(&total)

	return int(total)
}

func (c *CurrencyInfoMapper) PageByCondition(pageResult *vo.BasePageVO, currencyCode string) (list []domain.CurrencyInfo, err error) {
	wrapper := make(map[string]interface{}, 0)
	if currencyCode != "" {
		wrapper["currency_code"] = currencyCode
	}

	list = []domain.CurrencyInfo{}
	err = setting.DB.Model(&domain.CurrencyInfo{}).Scopes(paginateUtil.Paginate(pageResult)).Where(wrapper).Find(&list).Error
	return list, err
}

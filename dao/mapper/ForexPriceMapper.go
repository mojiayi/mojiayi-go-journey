package mapper

import (
	"errors"
	"mojiayi-go-journey/constants"
	"mojiayi-go-journey/dao/domain"
	"mojiayi-go-journey/setting"
)

type ForexPriceMapper struct{}

func (f *ForexPriceMapper) Insert(forexPrice domain.ForexPrice) (int, error) {
	setting.DB.Create(&forexPrice)
	if forexPrice.ID > 0 {
		return forexPrice.ID, nil
	}
	return constants.INT_ZERO, errors.New("创建外汇牌价（" + forexPrice.SrcCurrencyCode + ":" + forexPrice.DestCurrencyCode + "）信息失败")
}

func (f *ForexPriceMapper) Modify(forexPrice domain.ForexPrice) int {
	wrapper := make(map[string]interface{})
	wrapper["id"] = forexPrice.ID
	setting.DB.Model(&domain.ForexPrice{}).Where(wrapper).Updates(forexPrice)

	return constants.INT_ONE
}

func (f *ForexPriceMapper) SelectByCurrencyCode(srcCurrencyCode string, destCurrencyCode string) (domain.ForexPrice, error) {
	wrapper := make(map[string]interface{})
	wrapper["src_currency_code"] = srcCurrencyCode
	wrapper["dest_currency_code"] = destCurrencyCode

	var forexPrice domain.ForexPrice
	setting.DB.Model(&domain.ForexPrice{}).Where(wrapper).Find(&forexPrice)

	if forexPrice.ID == 0 {
		return forexPrice, errors.New("外汇牌价" + srcCurrencyCode + "(" + destCurrencyCode + ")不存在")
	}

	return forexPrice, nil
}

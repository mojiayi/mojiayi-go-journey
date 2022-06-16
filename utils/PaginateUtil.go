package utils

import (
	"mojiayi-go-journey/constants"
	"mojiayi-go-journey/vo"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PaginateUtil struct {
}

func (p *PaginateUtil) GetCurrentPage(ctx *gin.Context) int {
	tempStr := ctx.Query(constants.CurrentPage)
	if tempStr == "" {
		return constants.DefaultCurrentPage
	}
	currentPage, err := strconv.Atoi(tempStr)
	if err != nil {
		return constants.DefaultCurrentPage
	}
	if currentPage <= 0 {
		return constants.DefaultCurrentPage
	}
	return int(currentPage)
}

func (p *PaginateUtil) GetPageSize(ctx *gin.Context) int {
	tempStr := ctx.Query(constants.PageSize)
	if tempStr == "" {
		return constants.DefaultPageSize
	}
	pageSize, err := strconv.Atoi(tempStr)
	if err != nil {
		return constants.DefaultPageSize
	}
	if pageSize <= 0 || int(pageSize) > constants.MaxPageSize {
		return constants.DefaultPageSize
	}
	return int(pageSize)
}

func (p *PaginateUtil) Paginate(page *vo.BasePageVO) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page.Pages = page.Total / page.PageSize
		if page.Total%page.PageSize != 0 {
			page.Pages++
		}
		if page.CurrentPage > page.Pages {
			page.CurrentPage = page.Pages
		}
		offset := int((page.CurrentPage - 1) * page.PageSize)
		return db.Offset(offset).Limit(int(page.PageSize))
	}
}

package vo

/**
* 分页查询的公共返回字段
 */
type BasePageVO struct {
	CurrentPage int         `json:"currentPage"`
	PageSize    int         `json:"pageSize"`
	Total       int         `json:"total"`
	Pages       int         `json:"pages"`
	List        interface{} `json:"list"`
}

package param

/**
* 分页查询的公共参数
 */
type PageParam struct {
	CurrentPage int `json:"currentPage"`
	PageSize    int `json:"pageSize"`
}

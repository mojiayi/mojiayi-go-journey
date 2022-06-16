package vo

/**
* 分页查询的公共返回字段
 */
type BasePageVO struct {
	CurrentPage int
	PageSize    int
	Total       int
	Pages       int
	List        interface{}
}

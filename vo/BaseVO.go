package vo

/**
* 返回结果公共字段
 */
type BaseVO struct {
	/*
	* 请求状态码，使用 net/http中的标准状态码
	 */
	Code int `json:"code"`
	/*
	* 请求提示信息，当状态码为失败的值时，需要展示给用户
	 */
	Msg string `json:"msg"`
	/*
	* 请求访问时间
	 */
	Timestamp int64 `json:"timestamp"`
	/*
	 * 请求追踪id
	 */
	TraceId string `json:"traceId"`
	/**
	* 返回结果数据，根据接口业务不同，返回不同类型的结果
	 */
	Data interface{} `json:"data"`
}

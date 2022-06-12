package vo

/**
* 返回结果公共字段
 */
type BaseVO struct {
	/*
	* 请求状态码，使用 net/http中的标准状态码
	 */
	Code int
	/*
	* 请求提示信息，当状态码为失败的值时，需要展示给用户
	 */
	Msg string
	/*
	* 请求访问时间
	 */
	Timestamp int64
	/*
	 * 请求追踪id
	 */
	TraceId string
	/**
	* 返回结果数据，根据接口业务不同，返回不同类型的结果
	 */
	Data interface{}
}

func (v *BaseVO) GetCode() int {
	return v.Code
}

func (v *BaseVO) SetCode(code int) *BaseVO {
	v.Code = code
	return v
}

func (v *BaseVO) GetMsg() string {
	return v.Msg
}

func (v *BaseVO) SetMsg(msg string) *BaseVO {
	v.Msg = msg
	return v
}

func (v *BaseVO) GetTimestamp() int64 {
	return v.Timestamp
}

func (v *BaseVO) SetTimestamp(timestamp int64) *BaseVO {
	v.Timestamp = timestamp
	return v
}

func (v *BaseVO) GetTraceId() string {
	return v.TraceId
}

func (v *BaseVO) SetTraceId(traceId string) *BaseVO {
	v.TraceId = traceId
	return v
}

func (v *BaseVO) GetData() interface{} {
	return v.Data
}

func (v *BaseVO) SetData(data interface{}) *BaseVO {
	v.Data = data
	return v
}

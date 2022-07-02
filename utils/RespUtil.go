package utils

import (
	"mojiayi-go-journey/constants"
	"mojiayi-go-journey/vo"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type RespUtil struct {
}

func (r *RespUtil) IllegalArgumentErrorResp(msg string, context *gin.Context) {
	var resp = *new(vo.BaseVO)
	resp.Code = http.StatusBadRequest
	resp.Msg = msg
	resp.Timestamp = time.Now().UnixMilli()
	resp.Data = make(map[string]string, 0)
	resp.TraceId = context.Request.Header.Get(constants.TraceId)
	context.JSON(http.StatusOK, resp)
}

func (r *RespUtil) ErrorResp(code int, msg string, context *gin.Context) {
	var resp = *new(vo.BaseVO)
	resp.Code = code
	resp.Msg = msg
	resp.Timestamp = time.Now().UnixMilli()
	resp.Data = make(map[string]string, 0)
	resp.TraceId = context.Request.Header.Get(constants.TraceId)
	context.JSON(http.StatusOK, resp)
}

func (r *RespUtil) SuccessResp(data interface{}, context *gin.Context) {
	var resp = *new(vo.BaseVO)
	resp.Code = http.StatusOK
	resp.Msg = http.StatusText(http.StatusOK)
	resp.Timestamp = time.Now().UnixMilli()
	resp.Data = data
	resp.TraceId = context.Request.Header.Get(constants.TraceId)
	context.JSON(http.StatusOK, resp)
}

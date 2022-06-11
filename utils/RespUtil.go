package utils

import (
	"mojiayi-go-journey/vo"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type RespUtil struct {
}

func (r *RespUtil) IllegalArgumentErrorResp(msg string, context *gin.Context) {
	var resp = *new(vo.BaseVO)
	resp.SetCode(http.StatusBadRequest)
	resp.SetMsg(msg)
	resp.SetTimestamp(time.Now().UnixMilli())
	resp.SetData(make(map[string]string, 0))
	context.JSON(http.StatusOK, resp)
}

func (r *RespUtil) ErrorResp(code int, msg string, context *gin.Context) {
	var resp = *new(vo.BaseVO)
	resp.SetCode(code)
	resp.SetMsg(msg)
	resp.SetTimestamp(time.Now().UnixMilli())
	resp.SetData(make(map[string]string, 0))
	context.JSON(http.StatusOK, resp)
}

func (r *RespUtil) SuccessResp(data interface{}, context *gin.Context) {
	var resp = vo.BaseVO{}
	resp.SetCode(http.StatusOK)
	resp.SetMsg(http.StatusText(http.StatusOK))
	resp.SetTimestamp(time.Now().UnixMilli())
	resp.SetData(data)
	context.JSON(http.StatusOK, resp)
}

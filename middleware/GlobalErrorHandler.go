package middlewire

import (
	"mojiayi-go-journey/setting"
	"mojiayi-go-journey/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

var respUtil utils.RespUtil

func HandleError(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			setting.MyLogger.Info(r)

			respUtil.ErrorResp(http.StatusForbidden, r.(string), ctx)

			ctx.Abort()
		}
	}()
	ctx.Next()
}

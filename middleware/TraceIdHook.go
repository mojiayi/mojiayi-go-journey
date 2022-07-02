package middleware

import (
	"mojiayi-go-journey/constants"
	"mojiayi-go-journey/setting"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-basic/uuid"
	"github.com/sirupsen/logrus"
)

type TraceIdHook struct {
	TraceId  string
	LocalCtx *gin.Context
}

func EnableTraceIdHook(ctx *gin.Context) {
	var traceId = strings.ReplaceAll(uuid.New(), "-", "")

	hook := newTraceIdHook(traceId, ctx)
	setting.MyLogger.AddHook(hook)
	setting.MetadataLogger.AddHook(hook)
}

func newTraceIdHook(traceId string, localCtx *gin.Context) logrus.Hook {
	hook := TraceIdHook{
		TraceId:  traceId,
		LocalCtx: localCtx,
	}
	return &hook
}

func (hook *TraceIdHook) Fire(entry *logrus.Entry) error {
	entry.Data[constants.TraceId] = hook.TraceId
	entry.Data[constants.Ctx] = hook.LocalCtx
	return nil
}

func (hook *TraceIdHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

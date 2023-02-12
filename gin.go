package hlog

import (
	"errors"
	"github.com/gin-gonic/gin"
)

type GinLoggerConf struct {
}

var NotRegisterMiddleware = errors.New("not register middleware tracing logger")

const (
	hLoggerGinContext = "hlog-gin"
)

func TracingLogger(conf *GormLoggerConf) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// TODO 打点
		traceId := ctx.GetHeader(global.RequestIdHeader)
		if traceId == "" {
			traceId = global.NextId(ctx)
			ctx.Header(global.RequestIdHeader, traceId)
		}
		l := NewLoggerFromRequestId(traceId)
		ctx.Set(hLoggerGinContext, l)
		ctx.Next()

	}
}

func LoggerFromContext(ctx *gin.Context) *Logger {
	logger, ok := ctx.Get(hLoggerGinContext)
	if !ok {
		panic(NotRegisterMiddleware.Error())
	}
	return logger.(*Logger)
}

package hlog

import "context"

type RequestIdGenerator interface {
	NextId(ctx context.Context) string
}

//type RequestContext interface {
//	GetRequestId() string
//}

//type HLogger interface {
//	Debug(ctx RequestContext, Msg string)
//	Info(ctx RequestContext, Msg string)
//	Warn(ctx RequestContext, Msg string)
//	Error(ctx RequestContext, Err error, mayReason string, Input ...any)
//	ErrorWithStack(ctx RequestContext, Err error, mayReason string, Input ...any)
//}

//func Debug(ctx RequestContext, Msg string) {
//	GetLogger().Debug(ctx, Msg)
//}
//
//func Info(ctx RequestContext, Msg string) {
//	GetLogger().Info(ctx, Msg)
//}
//
//func Warn(ctx RequestContext, Msg string) {
//	GetLogger().Warn(ctx, Msg)
//}
//
//func Error(ctx RequestContext, Err error, mayReason string, Input ...any) {
//	GetLogger().Error(ctx, Err, mayReason, Input...)
//}
//
//func ErrorWithStack(ctx RequestContext, Err error, mayReason string, Input ...any) {
//	GetLogger().ErrorWithStack(ctx, Err, mayReason, Input...)
//}

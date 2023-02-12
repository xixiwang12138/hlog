package hlog

import "context"

func StartSpan(ctx context.Context) *Logger {
	//监控打点
	return NewLogger(ctx)
}

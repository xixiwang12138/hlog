package hlog

type RequestContext interface {
	GetRequestId() string
	GetUserFlag() string
}

type HLogger interface {
	Debug(ctx RequestContext, msg string)
	Info(ctx RequestContext, msg string)
	Warn(ctx RequestContext, msg string)
	Err(ctx RequestContext, err error, mayReason string, input ...any)
	ErrWithStack(ctx RequestContext, err error, mayReason string, input ...any)
}

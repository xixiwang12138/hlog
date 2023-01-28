package hlog

import (
	"runtime/debug"
	"time"
)

type OutPutter interface {
	Output(lg *log)
}

type Logger struct {
	putter OutPutter
}

func (l *Logger) Debug(ctx RequestContext, msg string) {
	now := time.Now()
	log := newLog(ctx.GetUserFlag(), ctx.GetRequestId(), Debug, &now, msg)
	l.putter.Output(log)
}

func (l *Logger) Info(ctx RequestContext, msg string) {
	now := time.Now()
	log := newLog(ctx.GetUserFlag(), ctx.GetRequestId(), Info, &now, msg)
	l.putter.Output(log)
}

func (l *Logger) Warn(ctx RequestContext, msg string) {
	now := time.Now()
	log := newLog(ctx.GetUserFlag(), ctx.GetRequestId(), Warn, &now, msg)
	l.putter.Output(log)
}

func (l *Logger) Err(ctx RequestContext, err error, mayReason string, input ...any) {
	now := time.Now()
	log := newLog(ctx.GetUserFlag(), ctx.GetRequestId(), Error, &now, "")
	log.SetError(err)
	if len(input) >= 1 {
		log.SetArgs(input[0])
	}
	log.SetMayCause(mayReason)
	l.putter.Output(log)
}

func (l *Logger) ErrWithStack(ctx RequestContext, err error, mayReason string, input ...any) {
	now := time.Now()
	log := newLog(ctx.GetUserFlag(), ctx.GetRequestId(), Error, &now, "")
	log.SetError(err)
	if len(input) >= 1 {
		log.SetArgs(input[0])
	}
	log.SetMayCause(mayReason)
	log.SetStack(debug.Stack())
	l.putter.Output(log)
}

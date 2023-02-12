package hlog

import (
	"context"
	. "github.com/xixiwang12138/hlog/decode"
	"runtime/debug"
	"time"
)

type OutPutter interface {
	Output(lg *Log)
}

type Logger struct {
	requestId string
	metadata  map[string]any
	putter    OutPutter
}

// Implement context.Context

func (l *Logger) Deadline() (deadline time.Time, ok bool) {
	return
}

func (l *Logger) Done() <-chan struct{} {
	//TODO implement me
	panic("implement me")
}

func (l *Logger) Err() error {
	//TODO implement me
	panic("implement me")
}

func (l *Logger) Value(key any) any {
	if key == global.RequestIdHeader {
		return l.GetRequestId()
	}
	return l.metadata[key.(string)]
}

func (l *Logger) GetRequestId() string {
	return l.requestId
}

// Constructor

func NewLogger(ctx context.Context) *Logger {
	reqId := ctx.Value(global.RequestIdHeader).(string)
	if reqId == "" {
		panic("Cannot Find Tracing Context")
	}
	return &Logger{putter: global.OutPutter, requestId: reqId, metadata: map[string]any{}}
}

func NewLoggerFromRequestId(req string) *Logger {
	return &Logger{putter: global.OutPutter, requestId: req, metadata: map[string]any{}}
}

// Opentracing

func (l *Logger) StartSpan() *Logger { //TODO 集成Opentracing api规范
	return l
}

// Api

func (l *Logger) Debug(msg string) {
	now := time.Now()
	log := NewLog(l.GetRequestId(), LevelDebug, &now, msg)
	l.putter.Output(log)
}

func (l *Logger) Info(msg string) {
	now := time.Now()
	log := NewLog(l.GetRequestId(), LevelInfo, &now, msg)
	l.putter.Output(log)
}

func (l *Logger) Warn(msg string) {
	now := time.Now()
	log := NewLog(l.GetRequestId(), LevelWarn, &now, msg)
	l.putter.Output(log)
}

func (l *Logger) Error(err error, mayReason string, input ...any) {
	now := time.Now()
	log := NewLog(l.GetRequestId(), LevelError, &now, "")
	log.SetError(err)
	log.SetArgs(input)

	log.SetMayCause(mayReason)
	l.putter.Output(log)
}

func (l *Logger) ErrorWithStack(err error, mayReason string, input ...any) {
	now := time.Now()
	log := NewLog(l.GetRequestId(), LevelError, &now, "")
	log.SetError(err)
	log.SetArgs(input)
	log.SetMayCause(mayReason)
	log.SetStack(debug.Stack())
	l.putter.Output(log)
}

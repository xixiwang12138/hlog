package hlog

import (
	"github.com/xixiwang12138/hlog/conf"
	"github.com/xixiwang12138/hlog/internal/sources"
	"runtime/debug"
	"time"
)

var (
	MongoLogCollector = &logRepo{}
	DefaultLogger     = &Logger{putter: MongoLogCollector}
)

func GetLogger() *Logger {
	return DefaultLogger
}

func SetMongoCollector(mongo *conf.MongoDBConfig) {
	sources.MongoSource.Setup(mongo)
	MongoLogCollector.Setup()
}

type OutPutter interface {
	Output(lg *log)
}

type Logger struct {
	putter OutPutter
}

func (l *Logger) Debug(ctx RequestContext, msg string) {
	now := time.Now()
	log := newLog(ctx.GetUserFlag(), ctx.GetRequestId(), LevelDebug, &now, msg)
	l.putter.Output(log)
}

func (l *Logger) Info(ctx RequestContext, msg string) {
	now := time.Now()
	log := newLog(ctx.GetUserFlag(), ctx.GetRequestId(), LevelInfo, &now, msg)
	l.putter.Output(log)
}

func (l *Logger) Warn(ctx RequestContext, msg string) {
	now := time.Now()
	log := newLog(ctx.GetUserFlag(), ctx.GetRequestId(), LevelWarn, &now, msg)
	l.putter.Output(log)
}

func (l *Logger) Error(ctx RequestContext, err error, mayReason string, input ...any) {
	now := time.Now()
	log := newLog(ctx.GetUserFlag(), ctx.GetRequestId(), LevelError, &now, "")
	log.SetError(err)
	if len(input) >= 1 {
		log.SetArgs(input[0])
	}
	log.SetMayCause(mayReason)
	l.putter.Output(log)
}

func (l *Logger) ErrorWithStack(ctx RequestContext, err error, mayReason string, input ...any) {
	now := time.Now()
	log := newLog(ctx.GetUserFlag(), ctx.GetRequestId(), LevelError, &now, "")
	log.SetError(err)
	if len(input) >= 1 {
		log.SetArgs(input[0])
	}
	log.SetMayCause(mayReason)
	log.SetStack(debug.Stack())
	l.putter.Output(log)
}

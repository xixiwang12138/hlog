package hlog

import (
	"context"
	"github.com/xixiwang12138/hlog/conf"
	"github.com/xixiwang12138/hlog/internal/sources"
	"runtime/debug"
	"time"
)

var (
	MongoLogCollector = &logRepo{}
	DefaultCollector  = MongoLogCollector
	DefaultLogger     = &Logger{putter: MongoLogCollector, requestId: "init_context"}
)

func SetMongoCollector(mongo *conf.MongoDBConfig) {
	sources.MongoSource.Setup(mongo)
	MongoLogCollector.Setup()
}

type OutPutter interface {
	Output(lg *log)
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
	return &Logger{putter: DefaultCollector, requestId: reqId, metadata: map[string]any{}}
}

func NewLoggerFromRequestId(req string) *Logger {
	return &Logger{putter: DefaultCollector, requestId: req, metadata: map[string]any{}}
}

// Opentracing

func (l *Logger) StartSpan() *Logger { //TODO 集成Opentracing api规范
	return l
}

// Api

func (l *Logger) Debug(msg string) {
	now := time.Now()
	log := newLog(l.GetRequestId(), LevelDebug, &now, msg)
	l.putter.Output(log)
}

func (l *Logger) Info(msg string) {
	now := time.Now()
	log := newLog(l.GetRequestId(), LevelInfo, &now, msg)
	l.putter.Output(log)
}

func (l *Logger) Warn(msg string) {
	now := time.Now()
	log := newLog(l.GetRequestId(), LevelWarn, &now, msg)
	l.putter.Output(log)
}

func (l *Logger) Error(err error, mayReason string, input ...any) {
	now := time.Now()
	log := newLog(l.GetRequestId(), LevelError, &now, "")
	log.SetError(err)
	if len(input) >= 1 {
		log.SetArgs(input[0])
	}
	log.SetMayCause(mayReason)
	l.putter.Output(log)
}

func (l *Logger) ErrorWithStack(err error, mayReason string, input ...any) {
	now := time.Now()
	log := newLog(l.GetRequestId(), LevelError, &now, "")
	log.SetError(err)
	if len(input) >= 1 {
		log.SetArgs(input[0])
	}
	log.SetMayCause(mayReason)
	log.SetStack(debug.Stack())
	l.putter.Output(log)
}

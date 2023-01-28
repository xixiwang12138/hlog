package hlog

import (
	"runtime"
	"time"
)

type metadata struct {
	count int64 //日志系统中存储的日志总量
}

type Level uint8

const (
	Debug Level = iota
	Info
	Warn
	Error
	Fatal
)

type log struct {
	user       string     //请求用户标识
	requestId  string     //请求链路唯一id
	level      Level      //日志等级
	createTime *time.Time //日志记录创建时间
	funcName   string     //日志创建的函数
	line       int32      //日志产生的行数
	file       string     //日志产生的文件
	msg        string     //日志信息

	input    any    //输入的参数
	err      error  //产生的错误
	mayCause string //产生错误的可能原因
	stack    []byte //堆栈信息
}

const SkipNum = 2

func newLog(user string, id string, level Level, createTime *time.Time, msg string) *log {
	l := &log{user: user, requestId: id, level: level, createTime: createTime, msg: msg}
	pc, file, line, _ := runtime.Caller(SkipNum)
	l.file = file
	l.line = int32(line)
	l.funcName = runtime.FuncForPC(pc).Name()
	return l
}

func (log *log) SetArgs(input any) {
	log.input = input
}

func (log *log) SetError(err error) {
	log.err = err
}

func (log *log) SetMayCause(cause string) {
	log.mayCause = cause
}

func (log *log) SetStack(stack []byte) {
	log.stack = stack
}

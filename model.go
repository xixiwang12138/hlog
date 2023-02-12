package hlog

import (
	"encoding/json"
	"runtime"
	"strconv"
	"time"
)

type Level byte

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
)

var levelList = []string{"Debug", "Info", "Warn", "Error"}

func (level Level) Name() string {
	return levelList[level]
}

type log struct {
	requestId      string //请求链路唯一id
	requestIdBytes []byte
	level          Level //日志等级
	levelBytes     []byte

	createTime *time.Time //日志记录创建时间
	timeBytes  []byte

	funcName   string //日志创建的函数
	line       uint32 //日志产生的行数
	file       string //日志产生的文件
	placeBytes []byte

	msg      string //日志信息
	msgBytes []byte

	input      any //输入的参数
	inputBytes []byte

	err        error //产生的错误
	errBytes   []byte
	mayCause   string //产生错误的可能原因
	causeBytes []byte
	stack      []byte //堆栈信息
}

const SkipNum = 2

func newLog(id string, level Level, createTime *time.Time, msg string) *log {
	l := &log{requestId: id, level: level, createTime: createTime, msg: msg}
	pc, file, line, _ := runtime.Caller(SkipNum)
	l.file = file
	l.line = uint32(line)
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

const (
	lp byte = '{'
	rp byte = '}'
	ls byte = '['
	rs byte = ']'

	rn    byte = '\n'
	rr    byte = '\r'
	comma byte = ','
	colon byte = ':'

	q byte = '"'

	reqIdLen = 32
)

//{"reqId": "32131312", "level": "Warn", "time": "", "place": "", "msg": "", "args": "", "cause": "", "stack": ""}

var (
	req   = []byte("\"reqId\"")
	level = []byte("\"level\"")
	time_ = []byte("\"time\"")
	place = []byte("\"place\"")
	msg   = []byte("\"msg\"")
	args  = []byte("\"args\"")
	err   = []byte("\"err\"")
	cause = []byte("\"cause\"")
	stack = []byte("stack")

	line = []byte("\n")

	fieldsNumWithoutErr = len(req) + len(level) + len(time_) + len(place) + len(msg) + 4 + 10 + 2 + 5
	fieldsNumWithErr    = len(req) + len(level) + len(time_) + len(place) + len(msg) + len(args) + len(cause) + len(stack) + len(err) + 8 + 9 + 18 + 2
)

func (log *log) toByteArray() []byte {
	log.transfer()
	n := log.bytesNum()
	arr := make([]byte, 0, n)
	arr = append(arr, lp)
	arr = append(arr, req...)
	arr = append(arr, colon, q)
	arr = append(arr, log.requestIdBytes...)
	arr = append(arr, q, comma)

	arr = append(arr, level...)
	arr = append(arr, colon, q)
	arr = append(arr, log.levelBytes...)
	arr = append(arr, q, comma)

	arr = append(arr, time_...)
	arr = append(arr, colon, q)
	arr = append(arr, log.timeBytes...)
	arr = append(arr, q, comma)

	arr = append(arr, place...)
	arr = append(arr, colon, q)
	arr = append(arr, log.placeBytes...)
	arr = append(arr, q, comma)

	arr = append(arr, msg...)
	arr = append(arr, colon, q)
	arr = append(arr, log.msgBytes...)

	if log.err == nil {
		arr = append(arr, q, rp)
		return arr
	}
	arr = append(arr, q, comma)

	arr = append(arr, err...)
	arr = append(arr, colon, q)
	arr = append(arr, log.errBytes...)
	arr = append(arr, q, comma)

	arr = append(arr, cause...)
	arr = append(arr, colon, q)
	arr = append(arr, log.causeBytes...)
	arr = append(arr, q, comma)

	arr = append(arr, stack...)
	arr = append(arr, colon, q)
	arr = append(arr, log.stack...)
	arr = append(arr, q, rp)
	arr = append(arr, line...)
	return arr
}

func (log *log) transfer() {
	log.requestIdBytes = []byte(log.requestId)
	log.levelBytes = []byte(levelList[log.level])

	log.placeBytes = []byte(log.file + "." + strconv.Itoa(int(log.line)))

	log.msgBytes = []byte(log.msg)
	log.timeBytes = timeToBytes(*log.createTime)

	if err != nil {
		i, _ := json.Marshal(log.inputBytes)
		log.inputBytes = i

		e, _ := json.Marshal(log.err)
		log.errBytes = e
	}
}

//{"reqId": "32131312", "level": "Warn", "time": "", "place": "",
//"msg": "", "args": "", "cause": "", "stack": ""}

func (log *log) bytesNum() (initNum int) {
	if log.err != nil {
		initNum += fieldsNumWithErr
	} else {
		initNum += fieldsNumWithoutErr
		initNum += len(log.errBytes)
		initNum += len(log.causeBytes)
		initNum += len(log.stack)
	}
	initNum += reqIdLen //reqId 固定长度
	initNum += 5        //level值固定长度
	initNum += len(log.timeBytes)
	initNum += len(log.placeBytes)
	return
}

func timeToBytes(t time.Time) []byte {
	return []byte(t.Format("2006/01/02 15:04:05.000"))
}

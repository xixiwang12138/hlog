package decode

import (
	"fmt"
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

type Log struct {
	RequestId      string //请求链路唯一id
	requestIdBytes []byte
	Level          Level //日志等级
	levelBytes     []byte

	CreateTime *time.Time //日志记录创建时间
	timeBytes  []byte

	Line       uint32 //日志产生的行数
	File       string //日志产生的文件
	placeBytes []byte

	Msg      string //日志信息
	msgBytes []byte

	Input      []any //输入的参数
	inputBytes []byte

	Err        error //产生的错误
	errBytes   []byte
	MayCause   string //产生错误的可能原因
	causeBytes []byte
	Stack      []byte //堆栈信息
}

const SkipNum = 2

func NewLog(id string, level Level, createTime *time.Time, msg string) *Log {
	l := &Log{RequestId: id, Level: level, CreateTime: createTime, Msg: msg}
	_, file, line, _ := runtime.Caller(SkipNum)
	l.File = file
	l.Line = uint32(line)
	return l
}

func (log *Log) SetArgs(input []any) {
	log.Input = input
}

func (log *Log) SetError(err error) {
	log.Err = err
}

func (log *Log) SetMayCause(cause string) {
	log.MayCause = cause
}

func (log *Log) SetStack(stack []byte) {
	log.Stack = stack
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

//{"reqId": "32131312", "Level": "Warn", "time": "", "place": "", "Msg": "", "args": "", "cause": "", "Stack": ""}

var (
	req   = []byte("\"reqId\"")
	level = []byte("\"Level\"")
	time_ = []byte("\"time\"")
	place = []byte("\"place\"")
	msg   = []byte("\"msg\"")
	args  = []byte("\"args\"")
	err_  = []byte("\"err\"")
	cause = []byte("\"cause\"")
	stack = []byte("\"stack\"")

	line = []byte("\n")

	fieldsNumWithoutErr = len(req) + len(level) + len(time_) + len(place) + len(msg) + 4 + 10 + 2 + 5
	fieldsNumWithErr    = len(req) + len(level) + len(time_) + len(place) + len(msg) + len(args) + len(cause) + len(stack) + len(err_) + 8 + 9 + 18 + 2
)

func (log *Log) ToByteArray() []byte {
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

	if log.Err == nil {
		arr = append(arr, q, rp)
		return arr
	}
	arr = append(arr, q, comma)

	arr = append(arr, err_...)
	arr = append(arr, colon, q)
	arr = append(arr, log.errBytes...)
	arr = append(arr, q, comma)

	arr = append(arr, cause...)
	arr = append(arr, colon, q)
	arr = append(arr, log.causeBytes...)
	arr = append(arr, q, comma)

	arr = append(arr, stack...)
	arr = append(arr, colon, q)
	arr = append(arr, log.Stack...)
	arr = append(arr, q, rp)
	arr = append(arr, line...)
	return arr
}

func (log *Log) transfer() {
	log.requestIdBytes = []byte(log.RequestId)
	log.levelBytes = []byte(levelList[log.Level])

	log.placeBytes = []byte(log.File + "." + strconv.Itoa(int(log.Line)))

	log.msgBytes = []byte(log.Msg)
	log.timeBytes = timeToBytes(*log.CreateTime)

	log.inputBytes = fmt.Append([]byte{}, log.Input)

	if log.Err != nil {
		e := []byte(log.Err.Error())
		log.errBytes = e
		log.causeBytes = []byte(log.MayCause)
	}

	fmt.Println()
}

//{"reqId": "32131312", "Level": "Warn", "time": "", "place": "",
//"Msg": "", "args": "", "cause": "", "Stack": ""}

func (log *Log) bytesNum() (initNum int) {
	if log.Err != nil {
		initNum += fieldsNumWithErr
	} else {
		initNum += fieldsNumWithoutErr
		initNum += len(log.errBytes)
		initNum += len(log.causeBytes)
		initNum += len(log.Stack)
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

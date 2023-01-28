package hlog

import (
	"fmt"
	"github.com/xixiwang12138/hlog/internal/repo"
	"github.com/xixiwang12138/hlog/internal/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type MongoModel struct {
	ID primitive.ObjectID `bson:"_id"`
}

type LogRecord struct {
	MongoModel `bson:",inline"`
	User       string     `json:"user,omitempty" bson:"user"`
	RequestId  string     `json:"requestId,omitempty" bson:"requestId"`   //请求链路唯一id
	Level      Level      `json:"level,omitempty" bson:"level"`           //日志等级
	CreateTime *time.Time `json:"createTime,omitempty" bson:"createTime"` //日志记录创建时间
	FuncName   string     `json:"funcName,omitempty" bson:"funcName"`     //日志创建的函数
	Line       int32      `json:"line,omitempty" bson:"line"`             //日志产生的行数
	File       string     `json:"file,omitempty" bson:"file"`             //日志产生的文件
	Msg        string     `json:"msg,omitempty" bson:"msg"`
	Input      string     `json:"input,omitempty" bson:"input"`       //输入的参数
	Err        string     `json:"err,omitempty" bson:"err"`           //产生的错误
	MayCause   string     `json:"mayCause,omitempty" bson:"mayCause"` //产生错误的可能原因
	Stack      string     `json:"stack,omitempty" bson:"stack"`       //堆栈信息
}

type logRepo struct {
	repo.MongoRepo[LogRecord]
}

func (rep *logRepo) Output(lg *log) {
	lr := &LogRecord{
		User:       lg.user,
		RequestId:  lg.requestId,
		Level:      lg.level,
		CreateTime: lg.createTime,
		FuncName:   lg.funcName,
		Line:       lg.line,
		Msg:        lg.msg,
		File:       lg.file,
		Input:      utils.Serialize(lg.input),
		Err:        lg.err.Error(),
		MayCause:   lg.mayCause,
		Stack:      string(lg.stack),
	}
	err := rep.MongoRepo.Create(lr)
	fmt.Println(err) //TODO 设置错误处理
}

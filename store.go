package hlog

import (
	"fmt"
	"github.com/xixiwang12138/hlog/internal/repo"
	"github.com/xixiwang12138/hlog/internal/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type mongoModel struct {
	ID *primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
}

type logRecord struct {
	mongoModel `bson:",inline"`
	RequestId  string     `json:"requestId,omitempty" bson:"requestId"`   //请求链路唯一id
	Level      string     `json:"level,omitempty" bson:"level"`           //日志等级
	CreateTime *time.Time `json:"createTime,omitempty" bson:"createTime"` //日志记录创建时间
	Line       uint32     `json:"line,omitempty" bson:"line"`             //日志产生的行数
	File       string     `json:"file,omitempty" bson:"file"`             //日志产生的文件
	Msg        string     `json:"msg,omitempty" bson:"msg"`
	Input      string     `json:"input,omitempty" bson:"input"`       //输入的参数
	Err        string     `json:"err,omitempty" bson:"err"`           //产生的错误
	MayCause   string     `json:"mayCause,omitempty" bson:"mayCause"` //产生错误的可能原因
	Stack      string     `json:"stack,omitempty" bson:"stack"`       //堆栈信息
}

type logRepo struct {
	repo.MongoRepo[logRecord]
}

func (rep *logRepo) Output(lg *log) {
	lr := &logRecord{
		RequestId:  lg.requestId,
		Level:      lg.level.Name(),
		CreateTime: lg.createTime,
		Line:       lg.line,
		Msg:        lg.msg,
		File:       lg.file,
		MayCause:   lg.mayCause,
	}
	if lg.input != nil {
		lr.Input = utils.Serialize(lg.input)
	}
	if lg.stack != nil {
		lr.Stack = string(lg.stack)
	}
	if lg.err != nil {
		lr.Err = lg.err.Error()
	}
	err := rep.MongoRepo.Create(lr)
	fmt.Println(err) //TODO 设置错误处理
}

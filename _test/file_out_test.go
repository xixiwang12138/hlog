package _test

import (
	"github.com/pkg/errors"
	"github.com/xixiwang12138/hlog"
	"github.com/xixiwang12138/hlog/collector"
	"testing"
)

func TestFileOut(t *testing.T) {
	p := "./log"
	hlog.Custom(hlog.WithFileCollector(&collector.FileConf{
		Folder:     p,
		FilePrefix: "player",
	}))
	xl := hlog.NewLoggerFromRequestId("hello")
	xl.ErrorWithStack(errors.New("数据库连接错误"), "数据库", "123.43.1.2")
	xl.Error(errors.New("网络超时"), "http", "123.43.1.2")
}

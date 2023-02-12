package hlog

import (
	"github.com/xixiwang12138/hlog/collector"
	"github.com/xixiwang12138/hlog/conf"
	"github.com/xixiwang12138/hlog/internal/sources"
)

const (
	DefaultHeader = "X-Request-Id"
)

var global = &Context{RequestIdHeader: DefaultHeader}

type Context struct {
	RequestIdHeader string
	RequestIdGenerator
	OutPutter
}

type Option func()

func Custom(options ...Option) {
	if options != nil {
		for _, option := range options {
			option()
		}
	}
}

func WithRequestHeader(header string) Option {
	return func() {
		global.RequestIdHeader = header
	}
}

func WithRequestIdGenerator(gen RequestIdGenerator) Option {
	return func() {
		global.RequestIdGenerator = gen
	}
}

func WithMongoCollector(mongo *conf.MongoDBConfig) Option {
	return func() {
		sources.MongoSource.Setup(mongo)
		mongoRepo := &collector.LogRepo{}
		mongoRepo.Setup()
		global.OutPutter = mongoRepo
	}
}

func WithStderrCollector(stderrConf *collector.StderrConf) Option {
	return func() {
		global.OutPutter = collector.NewStdWriter(stderrConf)
	}
}

func WithFileCollector(f *collector.FileConf) Option {
	return func() {
		global.OutPutter = collector.NewFileWriter(f)
	}
}

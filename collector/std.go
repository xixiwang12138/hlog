package collector

import (
	"github.com/xixiwang12138/hlog/decode"
	"os"
)

type StderrConf struct {
}

type StdWriter struct {
	cf  *StderrConf
	std *os.File
}

func NewStdWriter(conf *StderrConf) *StdWriter {
	return &StdWriter{conf, os.Stderr}
}

func (s *StdWriter) Output(l *decode.Log) {
	s.std.Write(l.ToByteArray())
}

package hlog

import "os"

var writer = os.Stderr
var stdOupter = &StdWriter{writer}

type StdWriter struct {
	std *os.File
}

func (s *StdWriter) Output(l *log) {
	s.std.Write(l.toByteArray())
}

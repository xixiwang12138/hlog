package hlog

import (
	"testing"
)

func TestToBytes(t *testing.T) {
	logger := &Logger{
		requestId: "ab4324213e",
		putter:    stdOupter,
	}
	logger.Info("***")
}

package hlog

import (
	"fmt"
	"runtime"
	"testing"
)

func A() {
	pc, file, line, ok := runtime.Caller(0)
	name := runtime.FuncForPC(pc).Name()
	fmt.Println(file, line, name, ok)
	B()
}

func B() {
	pc, file, line, ok := runtime.Caller(0)
	name := runtime.FuncForPC(pc).Name()
	fmt.Println(file, line, name, ok)
	C()
}

func C() {
	pc, file, line, ok := runtime.Caller(1)
	name := runtime.FuncForPC(pc).Name()
	fmt.Println(file, line, name, ok)
}

func TestGetCaller(t *testing.T) {
	A()
}

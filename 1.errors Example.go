package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func fn() error{
	e1 := errors.New("error")
	e2 := errors.Wrap(e1, "inner")
	e3 := errors.Wrap(e2, "middle")
	return errors.Wrap(e3, "outer")
}



func main() {
	type stackTracer interface {
		StackTrace() errors.StackTrace
	}

	err, ok := errors.Cause(fn()).(stackTracer)
	if !ok {
		panic("oops, err does not implement stackTracer")
	}

	st := err.StackTrace()  // []Frame数组  st[0]寻找到原始错误位置  st[1]Cause寻找到原因位置
	fmt.Printf("%+v", st[0:2])

	/* OutPut
	main.fn
		E:/Golang Projects/src/Geekgo/2.errors/1.errors Example.go:9
	main.main
		E:/Golang Projects/src/Geekgo/2.errors/1.errors Example.go:22
	*/
}
package main

import (
	//	"fmt"
	"testing"

	libf "libfunc"
)

/*
	闭包:指有权访问另一个函数作用域中变量的函数，创建闭包常见的方式就是在一个函数内部创建另一个函数。
*/
func Test_Closures(t *testing.T) {
	nextInt := intSeq()

	prnlog.LogPrint(libf.LOG_DEBUG, 0, false, false, "", nextInt())
	prnlog.LogPrint(libf.LOG_DEBUG, 0, false, false, "", nextInt())
	prnlog.LogPrint(libf.LOG_DEBUG, 0, false, false, "", nextInt())

	newInts := intSeq()
	prnlog.LogPrint(libf.LOG_DEBUG, 0, false, false, "", newInts())
}

func intSeq() func() int {
	i := 0
	prnlog.LogPrint(libf.LOG_DEBUG, 0, false, false, "", i)
	return func() int {
		i += 1
		return i
	}
}

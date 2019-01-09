package main

import (
	"fmt"
	"testing"
)

/*
  变量
*/
func Test_Variables(t *testing.T) {
	var a string = "a is string"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "这里是声明并初始化变量的简写"
	fmt.Printf(f)
}

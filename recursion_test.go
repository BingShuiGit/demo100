package main

import (
	"fmt"
	"testing"
)

/*
	递归
*/
func Test_Recursion(t *testing.T) {
	res := fact(4)
	fmt.Println("result:", res)
}

//利用递归求阶乘
func fact(n int) int {
	//face 函数在到达 face(0) 前一直调用自身
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

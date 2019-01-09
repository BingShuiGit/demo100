package main

import (
	"fmt"
	"testing"
)

/*
	if/else分支结构
*/
func Test_IfElse(t *testing.T) {
	//判断10以内的奇偶数
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Printf("%d is even!\n", i)
		} else {
			fmt.Printf("%d is odd!\n", i)
		}
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}

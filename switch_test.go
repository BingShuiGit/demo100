package main

import (
	"fmt"
	"testing"
	"time"
)

func Test_Switch(t *testing.T) {

	//表达式switch
	i := 2
	fmt.Print("write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	}

	fmt.Println(time.Now().Weekday())
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	tm := time.Now()
	fmt.Println(tm)
	switch {
	case tm.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	//调用function_test.go中定义的f_fun函数, 内含 "类型switch"
	f_fun(2, "Go", 8, "language", 'a', false, 'A', 3.14)
}

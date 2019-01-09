package main

import (
	"fmt"
	"testing"
	"time"
)

//通道的定义，传输int类型的值，允许缓存量为1
var ch1 = make(chan int, 1)
var ch2 = make(chan int, 1)
var chs = []chan int{ch1, ch2}
var numberss = []int{1, 2, 3, 4, 5}

//select语句会一直等待，直到某个case里的IO操作可以进行
//golang 的 select 就是监听 IO 操作，当 IO 操作发生时，触发相应的动作
//与switch语句可以选择任何可使用相等比较的条件相比，select有比较多的限制，其中最大的一条限制就是每个case语句里必须是一个IO操作，确切的说，应该是一个面向channel的IO操作
func Test_Select(t *testing.T) {
	//	var ch1 = make(chan int, 1)
	//	var ch2 = make(chan int, 1)
	go ff1(ch1)
	go ff2(ch2)
	select {
	case <-ch1:
		fmt.Println("The first case is selected.")
	case <-ch2:
		fmt.Println("The second case is selected.")
	}

	select {
	case getChan(0) <- getNumber(2):
		fmt.Println("1th case is selected.")
	case getChan(1) <- getNumber(3):
		fmt.Println("2th case is selected.")
	default:
		fmt.Println("default!.")
	}
}

func ff1(ch chan int) {
	time.Sleep(time.Second * 5)
	ch <- 1
}
func ff2(ch chan int) {
	time.Sleep(time.Second * 10)
	ch <- 1
}

func getNumber(i int) int {
	fmt.Printf("numbers[%d]\n", i)
	return numberss[i]
}
func getChan(i int) chan int {
	fmt.Printf("chs[%d]\n", i)
	return chs[i]
}

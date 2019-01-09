package main

import (
	"fmt"
	"testing"
	"time"
)

/*
* 协程(Coroutine): 当创建一个Go协程时，创建这个Go协程的语句立即返回。与函数不同，程序流程不会等待Go协程结束再继续执行。程序流程在开启Go协程后立即返回并开始执行下一行代码，
* 忽略Go协程的任何返回值。
* 在主协程存在时才能运行其他协程，主协程终止则程序终止，其他协程也将终止。
* 信道(Channel): 使用信道可以实现协程间通信，信道可以用于阻塞主协程，直到其他协程执行完毕。通过信道发送和接收数据默认是阻塞的,当数据发送给信道后，程序流程在发送语句处阻塞，
* 直到其他协程从该信道中读取数据。
* 同样地，当从信道读取数据时，程序在读取语句处阻塞，直到其他协程发送数据给该信道。
* 使用 make(chan val-type) 创建一个新的通道。通道类型就是他们需要传递值的类型。
* 使用 <-channel 语法从通道中 接收 一个值。
 */
func Test_Goroutines(t *testing.T) {
	f("direct")

	//在函数或方法调用之前加上关键字 go，这样便开启了一个并发的Go协程
	go f("goroutine")

	//msg:going   func(msg string) {fmt.Println(msg)}相当于上式中的f
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	//	go numbers()
	//	go alphabets()
	//	time.Sleep(3000 * time.Millisecond)

	done := make(chan bool)
	go hello(done)

	// 这一行从信道 done 中读取数据，但是没有使用该数据，也没有将它赋值给其他变量，这是完全合法的.现在 main 协程被阻塞，等待从信道 done 中读取数据。
	// hello 协程接受信道 done 作为参数，打印 Hello world goroutine 然后将数据写入信道 done 中。当写入完毕后，main 协程从信道 done 中接收到数据，
	// main 协程将解除阻塞，继续执行下一条语句，也就实现了执行完go hello（）后再继续执行下一条语句
	<-done

	done2 := make(chan bool)
	go alphabets(done2)
	<-done2

	var input string
	fmt.Scanln(&input)

	fmt.Println("done")

	//可缓存通道允许在没有对应接收方的情况下，缓存限定数量的值
	messages := make(chan string, 2)
	messages <- "print1"
	messages <- "print2"

	fmt.Println(<-messages)
	fmt.Println(<-messages)

	//通道选择器
	c1 := make(chan string)
	c2 := make(chan string)
	timeout := make(chan bool, 1)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "One"
	}()

	go func() {
		time.Sleep(time.Second * 3)
		c2 <- "Two"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		timeout <- true
	}()

	for i := 0; i < 2; i++ {
		select {
		// 从c1中读取到数据
		case msg1 := <-c1:
			fmt.Println("received:", msg1)
		case msg2 := <-c2:
			fmt.Println("received:", msg2)
		case msg3 := <-timeout:
			fmt.Println("超时机制: ", msg3)
		}
	}

	//通道方向
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pone(pings, pongs)
	fmt.Println(<-pongs)

}

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}
func alphabets(done2 chan bool) {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
	done2 <- true
}

func hello(done chan bool) {
	fmt.Println("Hello world goroutine")
	//time.Sleep(4 * time.Second)

	//信道同步
	done <- true
}

// ping 函数定义了一个只允许发送数据到该通道的通道。在发送语句中，<-运算符分割channel和要发送的值，只允许写。
// 在接收语句中，<-运算符写在channel对象之前，只允许读。一个不使用接收结果的接收操作也是合法的
func ping(pings chan<- string, msg string) {
	pings <- msg
}

//pong 函数允许通道（pings）来读取数据，另一通道（pongs）来写入数据。
func pone(pings <-chan string, pongs chan<- string) {
	// <-channel 语法从通道中读取一个值
	msg := <-pings
	//使用 channel <- 语法发送一个新的值到写到通道中
	pongs <- msg
}

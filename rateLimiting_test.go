package main

import (
	libf "libfunc"
	"testing"
	"time"
)

//速率限制(英) 是一个重要的控制服务资源利用和质量的途径。Go 通过 Go 协程、通道和打点器优美的支持了速率限制。
func Test_RateLimiting(t *testing.T) {

	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	//这个 limiter 通道将每 200ms 接收一个值。这个是速率限制任务中的管理器。
	//func Tick(d Duration) <-chan Time {} 返回值是一个只允许读取的通道
	limiter := time.Tick(time.Millisecond * 200)
	prnlog.LogPrint(libf.LOG_DEBUG, 0, false, false, "", "limiter: ", <-limiter, time.Now())

	for req := range requests {
		////会循环5次，前面往requests channel中发送了5个值
		<-limiter
		prnlog.LogPrint(libf.LOG_DEBUG, 0, false, false, "", "limiter:", req, <-limiter, time.Now())
		prnlog.LogPrint(libf.LOG_DEBUG, 0, false, false, "", "request", req, time.Now())
	}

	//有时候想临时进行速率限制，并且不影响整体的速率控制,可以通过通道缓冲来实现。这个 burstyLimiter 通道用来进行 3 次临时的脉冲型速率限制。
	burstyLimiter := make(chan time.Time, 3)
	//想将通道填充需要临时改变3次的值，做好准备。
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	//每 200 ms 我们将添加一个新的值到 burstyLimiter中，直到达到 3 个的限制。
	go func() {
		for k := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- k
		}
	}()

	burstyRequest := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequest <- i
	}
	close(burstyRequest)
	for req := range burstyRequest {
		//前三次没有速度限制，会直接打印出后面的println的内容，因为34行以存入3个缓存值
		//继续接收burstyLimiter值，除了前三次，后面的都是time.Tick进行速度限制,200ms打印一次，直到此次循环结束
		<-burstyLimiter
		prnlog.LogPrint(libf.LOG_DEBUG, 0, false, false, "", "request", req, time.Now())
	}
}

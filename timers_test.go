package main

import (
	"fmt"
	"testing"
	"time"
)

/*
* 定时器(Timer)和打点器(Ticker)：
* - 定时器 是当你想要在未来某一刻执行一次时使用的；
* - 打点器 则是当你想要在固定的时间间隔重复执行准备的。
 */
func Test_Timers(t *testing.T) {
	//计时器1
	timer1 := time.NewTimer(time.Second * 2)

	//类似于通道同步：直到这个定时器的通道 C 明确的发送了定时器失效的值之前，将一直阻塞。
	<-timer1.C
	fmt.Println("Timer 1 expired(过期)")

	//计时器2
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	//打点器
	ticker := time.NewTicker(time.Millisecond * 500)

	go func() {
		for i := range ticker.C {
			fmt.Println("Tick at", i)
		}
	}()
	time.Sleep(time.Millisecond * 1600)
	ticker.Stop()
	fmt.Println("Ticker stopped")

}

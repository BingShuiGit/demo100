package main

import (
	libf "libfunc"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

//可以使用一个互斥锁来在 Go 协程间安全的访问数据,锁定后只有一个goroutine在同一时刻访问一个共享变量，解锁后才允许其它goroutine访问该变量
func Test_Mutexex(t *testing.T) {
	var state = make(map[int]int)

	//mutex 将同步对 state 的访问
	var mutex = &sync.Mutex{}

	//为了比较基于互斥锁的处理方式和后面将要看到的其他方式，ops 将记录对 state 的操作次数。
	var ops_Mutexes int64 = 0

	//运行 100 个 Go 协程来重复读取 state
	for r := 0; r < 100; r++ {
		go func() {
			total_Mutexes := 0

			//如果 for 循环的头部没有条件语句，那么就会认为条件永远为 true。因此如果不想造成死循环，循环体内必须有相关的条件判断以确保会在某个时刻退出循环
			//每次循环读取，使用一个键来进行访问，Lock() 这个 mutex 来确保对 state 的独占访问，读取选定的键的值，Unlock() 这个mutex，并且 ops 值加 1
			for {
				key := rand.Intn(5)
				mutex.Lock()
				total_Mutexes += state[key]
				mutex.Unlock()
				atomic.AddInt64(&ops_Mutexes, 1)

				runtime.Gosched()
			}
		}()
	}

	//运行 10 个 Go 协程来模拟写入操作，使用和读取相同的模式。
	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddInt64(&ops_Mutexes, 1)
				runtime.Gosched()
			}
		}()
	}

	time.Sleep(time.Second)

	//获取并输出最终的操作计数
	opsFinal_Mutexes := atomic.LoadInt64(&ops_Mutexes)
	prnlog.LogPrint(libf.LOG_DEBUG, 0, false, false, "", "ops_Mutexes", opsFinal_Mutexes)

	//对 state 使用一个最终的锁，显示它是如何结束的
	mutex.Lock()
	prnlog.LogPrint(libf.LOG_DEBUG, 0, false, false, "", "state", state)
	mutex.Unlock()
}

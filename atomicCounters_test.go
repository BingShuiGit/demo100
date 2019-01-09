package main

import (
	libf "libfunc"
	"runtime"
	"sync/atomic"
	"testing"
	"time"
)

//使用 sync/atomic包在多个 Go 协程中进行 原子计数
func Test_AtomicCounters(t *testing.T) {
	var ops uint64 = 0

	for i := 0; i < 50; i++ {
		go func() {
			for {
				//使用 AddUint64 来让计数器自动增加，使用& 语法来给出 ops 的内存地址。atomic.AddUint64:原子.增/减
				atomic.AddUint64(&ops, 1)
				//允许其它 Go 协程的执行
				runtime.Gosched()
			}
		}()
	}
	time.Sleep(time.Second)

	//atomic.LoadUint64:原子.载入
	//为了在计数器还在被其它 Go 协程更新时，安全的使用它，我们通过 LoadUint64 将当前值的拷贝提取到 opsFinal中
	opsFinal := atomic.LoadUint64(&ops)
	prnlog.LogPrint(libf.LOG_DEBUG, 0, false, false, "", "ops: ", opsFinal)
}

package main

import (
	"fmt"
	"testing"
	"time"
)

/*
* 工作池
* 使用 Go 协程和通道实现一个工作池
 */
func Test_Workpools(t *testing.T) {
	jobs := make(chan int, 100)
	result := make(chan int, 100)

	for w := 1; w < 3; w++ {
		go worker(w, jobs, result)
	}

	for i := 1; i <= 9; i++ {
		jobs <- i
	}
	close(jobs)

	for a := 1; a <= 9; a++ {
		<-result
	}
}

func worker(id int, jobs <-chan int, result chan<- int) {
	for j := range jobs {
		fmt.Println("Work", id, "processing job", j)
		time.Sleep(time.Second)
		result <- j * 2
	}
}

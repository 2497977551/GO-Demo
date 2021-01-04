package main

import (
	"fmt"
	"time"
)

func worker(id int, job <-chan int, res chan<- int) {
	for i := range job {
		fmt.Printf("workerID %d, start job%b\n", id, i)
		time.Sleep(time.Second)
		fmt.Printf("worker:%d end job:%d\n", id, i)
		res <- i * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	//	开启三个goroutine
	for i := 0; i < 3; i++ {
		go worker(i, jobs, results)
	}

	// 发送五个任务
	for i := 0; i < 5; i++ {
		jobs <- i
	}
	close(jobs)

	// 输出结果
	for i := 0; i < 5; i++ {
		res := <-results
		fmt.Println(res)
	}
}

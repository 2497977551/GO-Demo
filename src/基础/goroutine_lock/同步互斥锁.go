package main

import (
	"fmt"
	"sync"
)

var (
	x    int
	wg   sync.WaitGroup
	lack sync.Mutex // 互斥锁
)

func add() {
	for i := 0; i < 5000; i++ {
		lack.Lock() // 上锁,上锁之后后续的goroutine必须等待解锁才能使用
		x = x + 1
		lack.Unlock() // 解锁

	}
	wg.Done() // 每次执行将计数器减1
}
func main() {
	wg.Add(2) // 设置计数器为2
	go add()
	go add()
	wg.Wait() // 等待计数器为0
	fmt.Println(x)
}

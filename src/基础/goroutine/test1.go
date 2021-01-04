package main

import (
	"fmt"

	"sync"
)

var wg sync.WaitGroup

//func hello(i int){
//	fmt.Println("hello golang",i)
//	wg.Done() // 计数器-1，相当于return函数内部执行到这里会退出
//}
func main() {

	//runtime.GOMAXPROCS(2) 自定义启用x个cpu内核，
	// 定义一个计数器,计数器有多少个就代表开启多少个并发线程
	wg.Add(10000)
	//for i := 0; i < 10000; i++ {
	//	go hello(i)
	//}
	// 匿名函数
	for i := 0; i < 10000; i++ {
		go func(i int) {
			fmt.Println(i)
			wg.Done() // 计数器减一
		}(i)
	}
	wg.Wait() // 等待goroutine执行完毕，计数器为0时进行下一步，如果还有数会阻塞
}

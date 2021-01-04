package main

import "fmt"

func main() {
	/*
		channel是一种类型，一种引用类型
		引用类型必须要初始化才能使用
	*/

	//	1. var 关键字声明
	//	var ch chan int
	//	ch = make(chan int,1)

	// 2. 简短声明
	ch := make(chan int, 1)
	ch <- 10  // 发送一个值给ch
	x := <-ch // 从ch中接收值
	fmt.Println(x)
	close(ch) // 关闭channel
}

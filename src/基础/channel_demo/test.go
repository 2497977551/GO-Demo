package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

/*
	1.在chan后面添加<-代表只能发送不能接收: chan<-
	2.在chan前面添加<-代表只能接收不能发送: <-chan
*/

// 循环发送10个值给ch1
func f1(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	// 循环结束关闭channel
	close(c)
	wg.Done()
}

// 循环接收并计算平方
func f2(c1 <-chan int, c2 chan<- int) {
	for {
		/*
			接收channel会返回两个值
			1.具体的值
			2.状态：true代表有值 false代表没有值
		*/
		v, ok := <-c1
		// 判断是否为false，为false代表没有值跳出循环
		if !ok {
			break
		}
		// 每次循环计算一次平方并发送给c2也就是ch2
		c2 <- v * v
	}
	// 循环结束关闭ch2这个chanel
	close(c2)
	wg.Done()
}
func main() {
	wg.Add(2)
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)
	go f1(ch1)
	go f2(ch1, ch2)

	// 通过循环获取ch2这个channel的值并输出
	for i := range ch2 {
		fmt.Println(i)
	}
	wg.Wait()
}

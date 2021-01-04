package main

import "fmt"

func main() {
	a()
	b()
	c()
}
func a() {
	fmt.Println("我是a函数")
}

/*
	panic 代表异常，代码碰到panic会崩溃退出
	所以需要使用recover收集异常并跳过panic继续执行接下来的代码
	recover必须配合defer使用，而defer必须声明在panic之前
	defer就是延迟执行，即代码最后执行
*/

func b() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("我是recover")
		}
	}()
	panic("我是panic")
}

func c() {
	fmt.Println("我是c函数")
}

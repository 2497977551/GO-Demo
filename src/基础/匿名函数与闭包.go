package main

import (
	"fmt"
	"strings"
)

//	闭包即内部函数调用外部变量，将函数当做返回值
func bibao() func() {
	name := "闭包函数"
	return func() {
		fmt.Println("我是", name)
	}
}

// 闭包练习:
func index1(str string) func(string) string {
	return func(name string) string {
		// 判断name的后缀是否是str，适合用于判断文件后缀
		if !strings.HasSuffix(name, str) {
			return name + str
		}
		return name
	}
}

// 闭包练习2
func index2(num int) (func(int) int, func(int) int) {
	add := func(num1 int) int {
		num += num1
		return num
	}
	sub := func(num2 int) int {
		num -= num2
		return num
	}
	return add, sub
}
func main() {
	//	匿名函数就是没有函数名的函数
	f := func() {
		fmt.Println("我是匿名函数1")
	}
	f()

	//	通过在函数后面添加圆括号可以立即执行该函数
	func() {
		fmt.Println("我是匿名函数2")
	}()

	//调用闭包
	r := bibao()
	r()
	//调用闭包练习1:
	i := index1(".exe")
	i2 := i("joshua")
	fmt.Println(i2)

	//	调用闭包练习2：x是add()函数,y是sub()函数
	x, y := index2(666)
	// x1接收的是 add()函数的返回值
	x1 := x(200)
	// y1接收的是 sub()函数的返回值
	y1 := y(300)
	fmt.Println(x1, y1)
}

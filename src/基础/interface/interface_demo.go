package main

import "fmt"

/*
	什么是接口？
	接口是一种类型，是一种抽象的类型
*/
// 定义一个猫的结构体，并添加了say的方法
type cat struct {
	name string
}

func (c *cat) say() {
	fmt.Printf("%s被打的喵喵叫\n", c.name)
}

// 定义一个狗的结构体，并添加了say的方法
type dog struct {
	name string
}

func (d *dog) say() {
	fmt.Printf("%s被打的汪汪叫\n", d.name)
}

// 定义接口
type sayer interface {
	say()
}

// 定义打的函数，不管传进来什么都会调用say()这个方法
func hit(a sayer) {
	a.say()
}
func main() {
	// 因为sayer接口是一种类型，所以可以在变量声明的时候使用
	var i sayer
	// 因为在方法中使用的是指针类型所以需要传入结构体地址
	c := &cat{
		name: "小橘",
	}
	d := &dog{
		name: "大黄",
	}
	i = d
	d.say()
	hit(c)
	fmt.Println(i)

}

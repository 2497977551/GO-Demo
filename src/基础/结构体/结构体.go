package main

import "fmt"

// 定义结构体
type cat struct {
	name  string
	color string
	age   int
}

func main() {
	// 实例化结构体
	var c cat
	c.name = "橘猫"
	c.color = "orange"
	c.age = 2
	fmt.Println(c)

	//	匿名结构体
	var user struct {
		// 同类型可以在同一行声明
		name, sex string
		age       int
	}
	fmt.Println(user.name)

	//	创建指针类型结构体
	var u = new(cat)
	fmt.Printf("%T\n", u)
	fmt.Printf("p2=%#v\n", u)

	//	取结构体的地址实例化
	//	使用&对结构体进行取地址操作相当于对该结构体类型进行了一次new实例化操作。
	var u2 = &cat{}
	u2.name = "波斯猫"
	fmt.Println(u2.name)
}

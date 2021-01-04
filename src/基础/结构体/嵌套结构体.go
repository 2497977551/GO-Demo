package main

import "fmt"

type address struct {
	city string
}
type hobby struct {
	game string
}

// 匿名结构体,匿名结构体就是结构体字段没有名称
type person3 struct {
	string
	int
	//	嵌套结构体它的类型是address
	address address
	//	当字段与类型同名时可以省略字段名，也就是匿名嵌套结构体
	hobby
}

func main() {

	p := person3{
		"joshua",
		20,
		address{"广州"},
		hobby{"英雄联盟"},
	}
	fmt.Printf("我的名字是%s，我今年%d，我家在%s，我最喜欢的游戏是%s",
		p.string, p.int, p.address.city, p.hobby.game)
}

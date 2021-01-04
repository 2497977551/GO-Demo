package main

import "fmt"

type residence struct {
	house string
}

func (r *residence) sleep() {
	fmt.Printf("它在%s里面睡觉\n", r.house)
}

//	此时dog结构体继承了residence结构体内的字段以及方法
type dog struct {
	name      string
	residence *residence
}

func (d *dog) cry() {
	fmt.Printf("%s会吼叫\n", d.name)
}
func main() {
	d := &dog{
		name:      "大黄",
		residence: &residence{house: "木屋"},
	}
	d.cry()
	d.residence.sleep()

}

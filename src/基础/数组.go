package main

import "fmt"

// 创建并初始化数组
var i = [...]int64{1, 3, 4, 5, 6}
var k = [...][5]int{
	{1, 2, 3, 4, 5},
	{6, 7, 8, 9, 10},
	{11, 12, 13, 14, 15},
	{16, 17, 18, 19, 20},
	{21, 22, 23, 24, 25},
	{26, 27, 28, 29, 30},
}

func main() {
	a := i
	fmt.Println(a[1])
	// 二维数组遍历1
	for index := 0; index < len(k); index++ {
		for index2 := 0; index2 < len(k[index]); index2++ {
			fmt.Print(k[index][index2], ",")
		}
	}
	//	二位数组遍历2
	for _, i3 := range k {
		for _, i5 := range i3 {
			fmt.Print(i5, ",")
		}
	}
}

package main

import "fmt"

func index(a, b, c int) int {
	if a > 0 {
		c := b + c
		fmt.Println(c)
	} else {
		fmt.Println("错误")
	}
	return c
}

func main() {
	index(1, 2, 3)

}

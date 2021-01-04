package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const name int = 20
	var number = rand.Intn(100)
	if name == number {
		fmt.Println("猜对了", number)
	} else if name > number {
		fmt.Println("太大了", number)
	} else {
		fmt.Println("太小了", number)
	}

}

package main

import "fmt"

func main() {
	c := make(chan bool, 1)
	go func() {
		fmt.Println("GO GO GO")
		c <- true
		close(c)
	}()
	for k := range c {
		fmt.Println(k)
	}

}

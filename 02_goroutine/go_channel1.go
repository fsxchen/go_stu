package main

import "fmt"

var b chan int

func main() {
	fmt.Println(b)

	b = make(chan int, 1)
	b <- 10
	fmt.Println(b)
}
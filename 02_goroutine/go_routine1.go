package main

import (
	"fmt"
	"time"
)


func hello() {
	fmt.Println("hello")
}

func main() {
	for i:=1; i< 1000; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	
	time.Sleep(time.Second)
	fmt.Println("main")
}
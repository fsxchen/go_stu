
package main

import "fmt"

func deferDemo() {
	fmt.Println("start")
	defer fmt.Println("end")
	fmt.Println("do something")
}

func main() {
	deferDemo()
}
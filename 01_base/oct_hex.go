package main

import "fmt"

func main() {
	var i1 = 101
	fmt.Printf("%d\n", i1)
	fmt.Printf("%b\n", i1)		//二进制输出
	fmt.Printf("%o\n", i1)		//八进制输出
	fmt.Printf("%x\n", i1)		//十六进制输出

	var i2 = 077
	fmt.Printf("%d\n", i2)

	i3 := 0x1234567
	fmt.Printf("%d\n", i3)
}
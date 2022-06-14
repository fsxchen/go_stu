/*
 * @Author: yangxingchen
 * @Date: 2022-06-14 16:49:24
 * @LastEditors: yangxingchen
 * @LastEditTime: 2022-06-14 16:59:31
 * @Description: 
 */
package main

import "fmt"


// 定义变量
var (
	name string
	age int
	isOk bool 
)

const (
	pi = 3.14159265358
)

const (
	n1 = iota
	n2
	n3
)

func main() {
	isOk = true
	name = "arron"
	age = 29

	// 类型推导
	var c = 18;
	
	// 简短申明，只能在函数中用
	s2 := "hell"

	fmt.Println(isOk)
	fmt.Printf("%s\n", name)
	fmt.Println(age)

	fmt.Println(pi)
	fmt.Println(c)
	fmt.Println(s2)

	fmt.Println(n3)
}

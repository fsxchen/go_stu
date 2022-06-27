/*
 * @Author: yangxingchen
 * @Date: 2022-06-20 15:58:14
 * @LastEditors: yangxingchen
 * @LastEditTime: 2022-06-20 16:04:33
 * @Description: 
 */
package main

import "fmt"

func f1(f func()) {
	fmt.Println("this is f1")
	f()
}

func f2(x, y int) {
	fmt.Println("this is f2")
	fmt.Println(x + y)
}

// 定义一个包装f2的函数

func f3(f func(int, int)) {
	return func() {
		f()
	}
}

func main() {
	f1(f2)
}
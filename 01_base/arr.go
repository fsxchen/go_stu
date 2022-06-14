/*
 * @Author: yangxingchen
 * @Date: 2022-06-14 17:32:46
 * @LastEditors: yangxingchen
 * @LastEditTime: 2022-06-14 17:40:28
 * @Description: 
 */
package main

import "fmt"

func main() {

	// 申明
	var a1 [3]bool		//默认是false【0值】
	var a2 [4]bool

	// 初始化
	// v1
	a1 = [3]bool{true, true, true}
	fmt.Println(a1)
	// v2 自动推导长度
	a10 := [...]int{0, 1, 2, 3, 4, 5}
	// v3 指定索引
	a3 := [5]int{0: 1, 4: 2}    //后面依然是默认的情况

	fmt.Println("a1: %T, a2: %T", a1, a2)
	fmt.Println(a10)
	fmt.Println(a3)
}
/*
 * @Author: yangxingchen
 * @Date: 2022-06-21 14:08:46
 * @LastEditors: yangxingchen
 * @LastEditTime: 2022-06-21 14:10:00
 * @Description: 
 */

package main

import "fmt"

type myInt int

func (m myInt) hello() {
	fmt.Println("I am an int")
}

func main() {
	m := myInt(180)
	m.hello()
}
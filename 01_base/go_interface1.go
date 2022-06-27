/*
 * @Author: yangxingchen
 * @Date: 2022-06-21 15:11:09
 * @LastEditors: yangxingchen
 * @LastEditTime: 2022-06-21 15:20:42
 * @Description: 
 */

package main

import "fmt"

type speaker interface {
	speak()
}

type cat struct{}

type dog struct {}

func (c cat) speak() {
	fmt.Println("mmmmmm")
}

func (d dog) speak() {
	fmt.Println("wwwwwww")
}

func da(x speaker) {
	x.speak()
}

func main()  {
	var c1 cat
	da(c1)
}
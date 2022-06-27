/*
 * @Author: yangxingchen
 * @Date: 2022-06-21 15:51:14
 * @LastEditors: yangxingchen
 * @LastEditTime: 2022-06-21 15:55:40
 * @Description: 
 */

package main

import "fmt"

type animal interface {
	move()
	eat(string)
}

type cat struct {
	name string
	feet int8
}

func (c *cat) move() {
	fmt.Println("move.....")
}

func (c cat) eat(food string) {
	fmt.Printf("cat eat %s\n", food)
}

func main() {
	var a1 animal
	// 下面两种都是可以
	c1 := cat{"tom", 4}
	c2 := &cat{"jialaolian", 4}

	a1 = c1
	a1 = c2
	fmt.Println(a1)
}

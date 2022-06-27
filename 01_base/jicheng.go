/*
 * @Author: yangxingchen
 * @Date: 2022-06-21 14:53:55
 * @LastEditors: yangxingchen
 * @LastEditTime: 2022-06-21 14:56:35
 * @Description: 
 */

package main

import "fmt"

type animal struct {
	name string
}

func (a animal) move() {
fmt.Printf("%s会动！", a.name)
}

type dog struct {
	feet uint8
	animal
}

func (d dog) wang() {
	fmt.Println("汪汪汪...", d.name)
}

func main() {
	d1 := dog {
		animal: animal{name: "aaa"},
		feet:	4,
	}

	fmt.Println(d1)
	d1.wang()
	d1.move()
}

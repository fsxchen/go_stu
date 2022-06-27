
package main

import "fmt"

type car interface {
	run()
}

type falali struct {
	brand string
}

func (f falali)run() {
	fmt.Printf("%s速度70\n", f.brand)
}

type baoshijie struct {
	brand string
}

func (f baoshijie)run() {
	fmt.Printf("%s速度700\n", f.brand)
}

func drive(c car) {
	c.run()
}

func main() {
	var f1 = falali{
		brand: "falali",
	}

	var b1 = baoshijie {
		"baoshijie",
	}

	drive(f1)
	drive(b1)
}
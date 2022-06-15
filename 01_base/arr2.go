package main

import "fmt"

func main() {
	a1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := 0

	for _, v := range a1 {
		sum += v
	}

	fmt.Println(sum)

	// 找出和为8的元素

	for i := 0; i < len(a1); i++ {
		for j := i+1; j < len(a1); j++ {
			if a1[i] + a1[j] == 8 {
				fmt.Printf("(%d, %d)\n", i, j)
			}
		}
	}
}
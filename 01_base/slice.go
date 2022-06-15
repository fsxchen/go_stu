package main

import "fmt"

func main() {
	var s1 []int                  // 定义一个存放int类型元素的切片
	var s2 []string				// string 类型的切片

	fmt.Println(s1, s2)

	s1 = []int{1,  2, 3}
	s2 = []string{"a", "b"}

	fmt.Println(s1, s2)

	// 长度和容量
	fmt.Printf("%d %d \n", len(s1), cap(s1))

	// 2 由数组定义切片
	a1 := [...]int{1, 2, 3, 4, 5, 6, 7}

	s3 := a1[0:4]

	s4 := a1[3:]
	s5 := a1[:4]
	s6 := a1[:]

	s7 := s6[3:]

	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(s5)

	a1[4] = 100
	fmt.Println(s6)
	fmt.Println(s7)
	fmt.Printf("%d %d\n", len(s7), cap(s7))
}
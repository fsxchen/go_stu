package main

import "fmt"

func main() {
	s1 := make([]int, 5, 10)
	fmt.Printf("s1=%v len(s1)=%d cpa(s1)=%d\n", s1, len(s1), cap(s1))

	s2 := make([]int, 0, 10)
	fmt.Printf("s2=%v len(s2)=%d cpa(s2)=%d\n", s2, len(s2), cap(s2))

	s3 := []int{1, 3, 5}
	s4 := s3

	fmt.Println(s3, s4)
	s3[0] = 100
	fmt.Println(s3, s4)

	for i:=0; i < len(s3); i++ {
		fmt.Println(s3[i])
	}

	for i, v := range(s3) {
		fmt.Println(i, v)
	}

	s3 = append(s3, 5, 6, 7)
	// s3 = s3.append()

	for i, v := range(s3) {
		fmt.Println(i, v)
	}

}
package main

import "fmt"

func assign(a interface{}) {
	fmt.Printf("%T\n", a)
	// str, ok := a.(string)
	// if !ok {
	// 	fmt.Println("错了")
	// } else {
	// 	fmt.Println(str)
	// }
	switch t := a.(type) {
	case string:
		fmt.Println("字符串", t)
	case int:
		fmt.Println("int", t)
	case int64:
		fmt.Println("int64", t)
	}
}

func main() {
	assign(100)
}
/*
 * @Author: yangxingchen
 * @Date: 2022-06-20 15:38:44
 * @LastEditors: yangxingchen
 * @LastEditTime: 2022-06-20 15:58:23
 * @Description: 
 */
package main
import "fmt"

func f1() {
	fmt.Println("Hello world")
}

func f2() int {
	return 1
}

// 函数作为参数
func f3(x func() int) {
	ret := x()
	fmt.Println(ret)
}

func main() {
	a := f1
	fmt.Printf("%T\n", a)

	b := f2
	fmt.Printf("%T\n", b)
}

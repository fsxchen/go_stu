/*
 * @Author: yangxingchen
 * @Date: 2022-06-20 17:45:17
 * @LastEditors: yangxingchen
 * @LastEditTime: 2022-06-20 17:51:43
 * @Description: 
 */

package main
import "fmt"

type dog struct {
	name string
}

func newDog(name string) *dog {
	return &dog {
		name: name,
	}
}

// 方法是作用于特定类型的函数
// 接受者表示的是调用该方法的具体类型变量。多用类型名首字母小写表示
func (d dog)wang() {
	fmt.Printf("%s 汪汪汪\n", d.name)
}

func main() {
	d1 := newDog("nam")
	d1.wang()
}
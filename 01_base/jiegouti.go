
package main
import "fmt"
type person struct {
	 name string
	 age int
	 hobby []string
	 gender string
}

func f(x person) {
	// 由于结构体是值类型的，所以不会对传过来的参数有所影响
	x.gender = "nv"
}

func f2(x *person) {
	(*x).gender = "女"
}

func main() {
	var z person
	z.name = "hello"
	z.age = 90
	z.gender = "男"
	z.hobby = []string{"a", "b"}

	fmt.Println(z)
	fmt.Printf("%T\n", z)

	var s struct {
		name string
		age int	
	}

	s.name = "hello"
	s.age = 80

	fmt.Println(s)

	f(z)

	fmt.Println(z)

	f2(&z)
	fmt.Println(z)

	// 结构体占用连续的内存
	type x struct{
		a int8
		b int8
		c int8
	}

	m := x{
		a: int8(10),
		b: int8(20),
		c: int8(30),
	}

	fmt.Println("%p\n", &(m.a))
	fmt.Println("%p\n", &(m.b))
	fmt.Println("%p\n", &(m.c))
}
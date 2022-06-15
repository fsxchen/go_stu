package main

import "fmt"

func main() {
	// 定义，未初始化，内存中没有空间
	var m1 map[string]int
	// 初始化
	m1 = make(map[string]int, 10)
	m1["aaa"] = 100
	m1["bbb"] = 200

	fmt.Println(m1)
	// 如果不存在，返回对应的0值
	fmt.Println(m1["dd"])
	v, ok := m1["cc"]
	if !ok {
		fmt.Println("no")
	} else {
		fmt.Println(v)
	}

	for k, v := range m1 {
		fmt.Println(k, v)
	}

	for k := range m1 {
		fmt.Println(k)
	}

	// 删除操作
	delete(m1, "bbb")
	fmt.Println(m1)
}
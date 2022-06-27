/*
 * @Author: yangxingchen
 * @Date: 2022-06-21 15:00:07
 * @LastEditors: yangxingchen
 * @LastEditTime: 2022-06-21 15:08:41
 * @Description: 
 */
package main

import "encoding/json"
import "fmt"

// 使用``告诉外部的包读取的时候的操作
type person struct {
	Name string `json:"name", db:"name"`
	Age int `json:"age"`
}

func main() {
	p1 := person{
		Name: "test",
		Age: 10,
	}

	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	fmt.Println(string(b))

	str := `{"name": "zzz", "age": 18}`

	var p2 person
	json.Unmarshal([]byte(str), &p2)

	fmt.Printf("%#v\n", p2)
}
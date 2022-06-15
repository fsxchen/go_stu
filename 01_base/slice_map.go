package main

import (
	"fmt"
	"math/rand"
	"time"
	"sort"
)



func main() {
	rand.Seed(time.Now().UnixNano())
	var socreMap = make(map[string]int, 200)

	for i:=0; i< 100; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100)
		socreMap[key] = value
	}

	fmt.Println(socreMap)

	var keys = make([]string, 0, 200)
	for key := range socreMap {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		fmt.Println(key, socreMap[key])
	}
}
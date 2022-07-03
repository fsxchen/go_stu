package main

import (
	"fmt"
	"time"
	"sync"
)

var wg sync.WaitGroup

// 通过全局变量进行控制
var notify bool

// 通过channel来控制

var exitChan chan bool

func f() {
	defer wg.Done()
LOOP:
	for {
		fmt.Println("Hello")
		time.Sleep(time.Millisecond * 500)
		// if notify {
		// 	break
		// }
		select {
		case <- exitChan:
			// break		// 这个break只能推出select，解决方法
			break LOOP
		default:

		}
	}
}

func main() {
	wg.Add(1)
	go f()
	// 如何让这个子goroutine推出呢？
	time.Sleep(time.Second * 5)
	// notify = true
	exitChan = make(chan bool)
	exitChan <- true

	wg.Wait()
}
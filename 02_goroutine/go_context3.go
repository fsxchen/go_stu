package main

import (
	"fmt"
	"time"
	"sync"
	"context"
)

var wg sync.WaitGroup


func f(ctx context.Context) {
	defer wg.Done()
LOOP:
	for {
		fmt.Println("Hello")
		time.Sleep(time.Millisecond * 500)
		// if notify {
		// 	break
		// }
		select {
		case <- ctx.Done():
			// break		// 这个break只能推出select，解决方法
			break LOOP
		default:

		}
	}
}

func main() {
	ctx, cancel := context.WaitCancel(context.Background())
	wg.Add(1)
	go f(ctx)
	// 如何让这个子goroutine推出呢？
	time.Sleep(time.Second * 5)
	// notify = true
	// 
	cancel()

	wg.Wait()
}
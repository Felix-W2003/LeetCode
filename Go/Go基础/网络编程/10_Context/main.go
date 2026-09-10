package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	ctx = context.WithoutCancel(ctx)
	go performTask(ctx)
	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}

func performTask(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("程序执行完毕")
			return
		default:
			fmt.Println("程序正在执行。。。。")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

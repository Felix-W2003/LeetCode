package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	ctx := context.WithValue(context.Background(), "userId", "123")
	cw, cancel := context.WithCancel(ctx)
	cancel()
	cw.Done()
	ctx2 := context.WithValue(ctx, "uuid", "2345")
	wg.Add(1)
	go performTask1(ctx)
	wg.Add(1)
	go performTask2(ctx2)
	time.Sleep(time.Second)
	fmt.Println("阻塞ing。。。。。。。")
	wg.Wait()
	fmt.Println("程序运行完毕")

}

func performTask1(ctx context.Context) {
	value1 := ctx.Value("userId")
	value2 := ctx.Value("uuid")
	fmt.Println(value1, value2)
	wg.Done()
}

func performTask2(ctx context.Context) {
	value1 := ctx.Value("userId")
	value2 := ctx.Value("uuid")
	fmt.Println(value1, value2)
	wg.Done()
}

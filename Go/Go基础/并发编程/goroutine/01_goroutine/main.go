package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func test1() {
	for i := 0; i < 100; i++ {
		fmt.Println("nihao test1", i, runtime.NumCPU())
		time.Sleep(time.Millisecond * 50)
	}
	wg.Done() //协程计数器-1
}

func test2() {
	for i := 0; i < 150; i++ {
		fmt.Println("nihao test2", i, runtime.NumCPU())
		time.Sleep(time.Millisecond * 50)
	}
	wg.Done() //协程计数器-1
}
func main() {
	wg.Add(1) //协程计数器+1
	go test1()

	wg.Add(1) //协程计数器+1
	go test2()
	for i := 0; i < 10; i++ {
		fmt.Println("nihao golang", i, runtime.NumCPU())
		time.Sleep(time.Millisecond * 50)
	}

	defer func() {
		wg.Wait()
		fmt.Println("defer退出")
	}()
	fmt.Println("主线程退出")
}

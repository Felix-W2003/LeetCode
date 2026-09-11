package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}

	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello" + fmt.Sprint(i)
	}

	// 使用select实现channel的多路复用的时候不需要关闭管道
	for {
		select {
		case v := <-intChan:
			fmt.Println(v)

		case v := <-stringChan:
			fmt.Println(v)

		default:
			fmt.Println("Failed to read channel")
			return
		}
	}
}

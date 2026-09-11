package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// 双向管道
	ch1 := make(chan int, 2)
	ch1 <- 10
	ch1 <- 12
	m1 := <-ch1
	m2 := <-ch1
	fmt.Println(m1, m2)

	// 声明管道为只写
	ch2 := make(chan<- int, 2)
	ch2 <- 10
	ch2 <- 12

	// 声明管道为只读
	//ch3 := make(<-chan int, 2)
}

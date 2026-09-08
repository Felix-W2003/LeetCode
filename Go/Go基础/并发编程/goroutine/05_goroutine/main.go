package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func fn1(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		time.Sleep(time.Millisecond * 50)
		ch <- i
		fmt.Printf("write %v to ch\n", i)
	}
	close(ch)
	wg.Done()
}

func fn2(ch <-chan int) {
	for num := range ch {
		time.Sleep(time.Millisecond * 50)
		fmt.Println("success to read", num)
	}
	wg.Done()
}
func main() {
	ch := make(chan int, 10)
	wg.Add(1)
	go fn1(ch)
	wg.Add(1)
	go fn2(ch)

	wg.Wait()
	fmt.Println("END")
}

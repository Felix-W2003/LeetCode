package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func addNum(ch chan int) {
	for i := 2; i <= 100; i++ {
		ch <- i
	}
	close(ch)
	wg.Done()
}

func isPrime(numPipe chan int, primeList chan int, exitChan chan bool) {
	for num := range numPipe {
		var flag bool = true
		for i := 2; i*i <= num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			primeList <- num
		}
	}

	exitChan <- true
	wg.Done()
}

func printPrime(primeList chan int) {
	for v := range primeList {
		fmt.Println(v)
	}
	wg.Done()
}
func main() {
	now := time.Now()
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 1000)
	exitChan := make(chan bool, 8)

	wg.Add(1)
	go addNum(intChan)

	for i := 1; i <= 8; i++ {
		wg.Add(1)
		go isPrime(intChan, primeChan, exitChan)
	}

	wg.Add(1)
	go printPrime(primeChan)

	wg.Add(1)
	go func() {
		for i := 0; i < 8; i++ {
			<-exitChan
		}

		close(primeChan)
		wg.Done()
	}()

	wg.Wait()

	fmt.Println("time is ", time.Since(now))
}

/*
问题：为什么没有重复打印？
答：管道(Channel)的消费特性，numPipe中的数字只能被一个
协程取出一次，一旦被取出，其他协程就再也拿不到这个数字了
这是Go管道的基本保证（FIFO队列）
*/

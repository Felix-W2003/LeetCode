package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func test(num int) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {

		fmt.Printf("携程（%v）打印的第%v条数据\n", num, i)
	}
}

func main() {
	for i := 1; i <= 7; i++ {
		wg.Add(1)
		go test(i)
	}
	wg.Wait()
	fmt.Println("关闭主线程")
}

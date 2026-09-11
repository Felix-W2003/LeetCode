package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go Task(ctx)
	time.Sleep(3 * time.Second)
	fmt.Println("Main goroutine : Done")

}

func Task(ctx context.Context) {

	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Task: Finished")
	case <-ctx.Done():
		fmt.Println("Task:Context cancelled or time out ", ctx.Err())
	}
}

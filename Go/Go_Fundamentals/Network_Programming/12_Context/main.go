package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))
	defer cancel()
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Task completed or deadline exceeded", ctx.Err())
				return
			default:
				time.Sleep(500 * time.Millisecond)
				fmt.Println("Task is running...")
			}
		}
	}(ctx)
	time.Sleep(3 * time.Second)

}

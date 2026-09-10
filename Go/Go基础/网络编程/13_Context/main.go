package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	stop := context.AfterFunc(ctx, func() { fmt.Println("After Func is running") })
	go func(ctx context.Context, stop func() bool) {
		select {
		case <-ctx.Done():
			fmt.Println("Context canceled", ctx.Err())
			fmt.Println("stop", stop())
		}
	}(ctx, stop)

	time.Sleep(3 * time.Second)

}

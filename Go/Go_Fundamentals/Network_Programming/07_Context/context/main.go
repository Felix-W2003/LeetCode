package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	urls := []string{
		"http://127.0.0.1/users",
		"http://127.0.0.1/products",
		"http://127.0.0.1/orders",
	}
	results := make(chan string, 10)
	for _, url := range urls {
		go fetchAPI(ctx, url, results)

	}

	for range urls {
		fmt.Println(<-results)
	}

}
func fetchAPI(ctx context.Context, url string, results chan string) {
	req, err := http.NewRequestWithContext(ctx, "Get", url, nil)
	if err != nil {
		results <- fmt.Sprintf("create request err:%v", err)
		return
	}
	client := http.DefaultClient
	response, err := client.Do(req)
	if err != nil {
		results <- fmt.Sprintln("fetch api err :%v", err)
		return
	}
	defer response.Body.Close()
	results <- fmt.Sprintf("fetch api success:%s,status :%d", url, response.StatusCode)
}

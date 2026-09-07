package main

import (
	"context"
	"log"
	"time"

	"github.com/elastic/go-elasticsearch/v9"
)

// InitClient 只负责创建客户端并测试连接
func InitClient() (*elasticsearch.TypedClient, error) {
	client, err := elasticsearch.NewTyped(
		elasticsearch.WithAddresses("http://localhost:9200"),
		elasticsearch.WithBasicAuth("elastic", "620WFwf0111"),
	)
	if err != nil {
		return nil, err
	}

	// 测试连接
	res, err := client.Info().Do(context.Background())
	if err != nil {
		return nil, err
	}
	log.Printf("Connected to Elasticsearch: %+v", res)

	return client, nil
}

func main() {
	client, err := InitClient()
	if err != nil {
		log.Fatalf("Failed to initialize client: %s", err)
	}
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := client.Close(ctx); err != nil {
			log.Printf("Error closing client: %s", err) // 仅记录，不终止程序
		}
	}()

	// 使用 client 进行后续操作...
}

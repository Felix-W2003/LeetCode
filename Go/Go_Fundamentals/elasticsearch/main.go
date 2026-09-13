package main

import (
	"Go/Go_Fundamentals/elasticsearch/initialize"
	"context"
	"log"
	"time"
)

func main() {
	es := initialize.InitEs()
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := es.Close(ctx); err != nil {
			log.Fatal("Error closing the client:", err)
		}
	}()

	res, err := es.Info().Do(context.Background())
	if err != nil {
		log.Fatal("Error getting info:", err)
	}
	log.Println(res)
}

package initialize

import (
	"log"

	"github.com/elastic/go-elasticsearch/v9"
)

func InitEs() *elasticsearch.TypedClient {
	es, err := elasticsearch.NewTyped(
		elasticsearch.WithAddresses(
			"http://localhost:9200"),
		elasticsearch.WithBasicAuth("elastic", "620WFwf0111"),
	)
	if err != nil {
		log.Fatal("Can not connect es:", err)
	}
	return es
}

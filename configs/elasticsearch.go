package configs

import (
	"log"

	"github.com/olivere/elastic/v7"
)

// var ESClient *elastic.Client

func ConnectElasticsearch(config *Config) *elastic.Client {
	client, err := elastic.NewClient(elastic.SetURL(config.ESURL), elastic.SetSniff(false))
	if err != nil {
		log.Fatal("Failed to connect to Elasticsearch:", err)
	}

	// ESClient = client
	log.Println("Connected to Elasticsearch")

	return client
}

package main

import (
	server "elastic-logger-app/api"
	"elastic-logger-app/configs"
	"log"
)

func main() {
	config := configs.LoadConfig()
	mongodbClient := configs.ConnectDB(config)
	elasticSearchClient := configs.ConnectElasticsearch(config)

	server := server.InitServer(":"+config.ServerPort, mongodbClient, elasticSearchClient)
	if err_run_server := server.RunApp(); err_run_server != nil {
		log.Fatal("Cannot run app: ", err_run_server)
	}
}

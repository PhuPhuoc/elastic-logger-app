package main

import (
	"context"
	server "elastic-logger-app/api"
	"elastic-logger-app/configs"
	"log"
)

func main() {
	// This context is used for all initialization steps (DB connection, services...)
	// and will be cancelled automatically when main() exits.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Load environment configuration
	config := configs.LoadConfig()

	// MySQL connection will live until the application exits.
	mysqlClient := configs.ConnectMysql(config)

	// ConnectMongodb should accept context so timeout/cancel is controlled by main.
	mongodbClient := configs.ConnectMongodb(ctx, config)
	// Defer MongoDB disconnection for clean shutdown.
	defer func() {
		if err := mongodbClient.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	// Connect to Elasticsearch
	elasticSearchClient := configs.ConnectElasticsearch(config)

	// Connect to RabbitMQ
	rabbitConn := configs.ConnectRabbitMQ(config)
	defer rabbitConn.Close()

	// Initialize HTTP server
	server := server.InitServer(":"+config.APP_PORT, mysqlClient, mongodbClient, elasticSearchClient)

	// Run HTTP server (blocking call)
	if err_run_server := server.RunApp(); err_run_server != nil {
		log.Fatal("Cannot run app: ", err_run_server)
	}
}

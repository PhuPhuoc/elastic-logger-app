package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	APP_PORT    string
	ELASTIC_URL string

	MONGODB_URI      string
	MONGODB_DATABASE string

	MYSQL_HOST     string
	MYSQL_PORT     string
	MYSQL_USER     string
	MYSQL_PASSWORD string
	MYSQL_DATABASE string

	RABBITMQ_URL string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		APP_PORT: getEnv("PORT", "8080"),

		// MongoDB
		MONGODB_URI:      getEnv("MONGODB_URI", "mongodb://localhost:27017"),
		MONGODB_DATABASE: getEnv("MONGODB_DATABASE", "myapp"),
		// Elastic
		ELASTIC_URL: getEnv("ELASTICSEARCH_URL", "http://localhost:9200"),

		// MySQL
		MYSQL_HOST:     getEnv("MYSQL_HOST", "localhost"),
		MYSQL_PORT:     getEnv("MYSQL_PORT", "3306"),
		MYSQL_USER:     getEnv("MYSQL_USER", "root"),
		MYSQL_PASSWORD: getEnv("MYSQL_PASSWORD", "rootpwd"),
		MYSQL_DATABASE: getEnv("MYSQL_DATABASE", "cmd_elastic_app"),

		// RabbitMQ
		RABBITMQ_URL: getEnv("RABBITMQ_URL", "amqp://guest:guest@localhost:5672"),
		// RABBITMQ_URL:  getEnv("RABBITMQ_URL", "amqp://guest:guest@localhost:5672/")
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

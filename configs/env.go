package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBURI      string
	DBName     string
	ESURL      string
	ServerPort string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		DBURI:      getEnv("MONGODB_URI", "mongodb://localhost:27017"),
		DBName:     getEnv("MONGODB_DATABASE", "myapp"),
		ESURL:      getEnv("ELASTICSEARCH_URL", "http://localhost:9200"),
		ServerPort: getEnv("PORT", "8080"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

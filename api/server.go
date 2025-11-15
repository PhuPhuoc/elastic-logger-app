package server

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"go.mongodb.org/mongo-driver/mongo"
)

type server struct {
	port    string
	mongo   *mongo.Client
	elastic *elastic.Client
}

func InitServer(port string, mongo *mongo.Client, elastic *elastic.Client) *server {
	return &server{
		port:    port,
		mongo:   mongo,
		elastic: elastic,
	}
}

func (server *server) RunApp() error {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	configcors := cors.DefaultConfig()
	configcors.AllowAllOrigins = true
	configcors.AllowMethods = []string{"POST", "GET", "PUT", "DELETE", "PATCH", "OPTIONS"}
	configcors.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept", "User-Agent", "Cache-Control", "Pragma"}
	configcors.ExposeHeaders = []string{"Content-Length"}
	configcors.AllowCredentials = true
	configcors.MaxAge = 12 * time.Hour

	router.Use(cors.New(configcors))
	router.GET("/ping", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "elastic-logger-app response: pong"}) })

	log.Println("server start listening at port: ", server.port)
	return router.Run(server.port)
}

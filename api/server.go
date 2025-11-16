package server

import (
	"database/sql"
	"elastic-logger-app/builder"
	accounthttp "elastic-logger-app/modules/account/infras/http"
	accountcommands "elastic-logger-app/modules/account/usecase/commands"
	accountqueries "elastic-logger-app/modules/account/usecase/queries"
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
	mysql   *sql.DB
	mongo   *mongo.Client
	elastic *elastic.Client
}

func InitServer(port string, mysql *sql.DB, mongo *mongo.Client, elastic *elastic.Client) *server {
	return &server{
		port:    port,
		mysql:   mysql,
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

	account_builder := builder.NewAccountBuilder(server.mysql, server.mongo)
	acc_cmd_builder := accountcommands.NewAccountCmdWithBuilder(account_builder)
	acc_query_builder := accountqueries.NewAccountQueryWithBuilder(account_builder)

	api := router.Group("/api/v1")
	{
		accounthttp.NewAccountHTTP(acc_cmd_builder, acc_query_builder).Routes(api)
	}

	log.Println("server start listening at port: ", server.port)
	return router.Run(server.port)
}

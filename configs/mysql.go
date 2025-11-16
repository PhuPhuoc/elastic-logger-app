package configs

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectMysql(config *Config) *sql.DB {
	dbHost := config.MYSQL_HOST
	dbPort := config.MYSQL_PORT
	dbUser := config.MYSQL_USER
	dbPassword := config.MYSQL_PASSWORD
	dbName := config.MYSQL_DATABASE

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=True&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
	log.Println("Connecting to MySQL with DSN:", dsn)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Open DSN failed: ", err)
	}

	// Ping to verify connection works
	if err := db.Ping(); err != nil {
		log.Fatal("Ping MySQL failed: ", err)
	}

	return db
}

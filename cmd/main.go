package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/agrz4/rest-api-postgre/cmd/api"
	"github.com/agrz4/rest-api-postgre/config"
	"github.com/agrz4/rest-api-postgre/db"
)

func main() {
	// init db
	dbConn, err := db.NewPostgresSQL(fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Envs.DBHost,
		config.Envs.DBPORT,
		config.Envs.DBUser,
		config.Envs.DBPassword,
		config.Envs.DBName,
	))
	if err != nil {
		log.Fatal("error connecting to postgres")
	}

	if err := initDB(dbConn); err != nil {
		log.Fatal("connection with db error : ", err)
	}
	// start api server
	apiServer := api.NewAPIServer(":8080", dbConn)
	if err := apiServer.Run(); err != nil {
		log.Fatal("error running api server")
	}
}

func initDB(db *sql.DB) error {
	err := db.Ping()
	if err != nil {
		return err
	}

	log.Println("Connected to database : ", config.Envs.DBName)

	return nil
}

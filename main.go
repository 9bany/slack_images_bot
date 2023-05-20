package main

import (
	"database/sql"
	"log"
	"os"

	server "github.com/9bany/bot_workflows/src"
	db "github.com/9bany/bot_workflows/src/db/sqlc"
	"github.com/9bany/bot_workflows/src/utils"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	envPathFile := os.Getenv("ENV_PATH")
	config := utils.LodConfig(envPathFile)

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Println("got at error: ", err)
	}

	driver, err := postgres.WithInstance(conn, &postgres.Config{})
	if err != nil {
		log.Println("can not create driver postgres: ", err)
	}

	log.Println("config.DBName: ", config.DBName)
	m, err := migrate.NewWithDatabaseInstance(
		"file://./src/db/migration",
		config.DBName,
		driver,
	)
	if err != nil {
		log.Println("can not create postgres migration: ", err)
	}
	m.Up()

	queries := db.New(conn)

	server, err := server.NewBot(config, queries)
	if err != nil {
		panic("Can not initialization server")
	}

	err = server.Start()
	if err != nil {
		panic("Can not start server")
	}

	os.Exit(0)
}

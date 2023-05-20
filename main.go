package main

import (
	"database/sql"
	"log"
	"os"

	server "github.com/9bany/bot_workflows/src"
	db "github.com/9bany/bot_workflows/src/db/sqlc"
	"github.com/9bany/bot_workflows/src/utils"

	_ "github.com/lib/pq"
)

func main() {
	mode := os.Getenv("MODE")
	log.Println("Running with mode: ", mode)

	var config utils.Config

	if mode == "dev" {
		config = utils.LodConfig(".", "dev")
	} else {
		config = utils.LodConfig(".", "app")
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Println("got at error: ", err)
	}

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

package main

import (
	"context"
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"

	db "github.com/9bany/bot_workflows/src/db/sqlc"
	"github.com/9bany/bot_workflows/src/utils"
)

func main() {
	config := utils.LodConfig(".", "dev")
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Println("got at error: ", err)
	}

	querieer := db.New(conn)

	files, err := utils.FilePathWalkDir("./data/images")
	if err != nil {
		log.Println("can not file path walk")
	}

	for index, element := range files {
		file_data, err := os.ReadFile(element)
		if err != nil {
			log.Println("can not open file")
		}

		_, err = querieer.CreateImage(context.Background(), file_data)
		if err != nil {
			log.Println("can not insert image to db")
		}

		log.Println("Inserted a image to db", index)
	}

}

package main

import (
	"os"

	server "github.com/9bany/bot_workflows/src"
	"github.com/9bany/bot_workflows/src/utils"
)

func main() {
	configuration := utils.LodConfig(".")
	server, err := server.NewBot(configuration)
	if err != nil {
		panic("Can not initialization server")
	}

	err = server.Start()
	if err != nil {
		panic("Can not start server")
	}

	os.Exit(0)
}

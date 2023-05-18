package main

import (
	"os"

	"github.com/9bany/bot_workflows/config"
	"github.com/9bany/bot_workflows/server"
)

func main() {
	configuration := config.LodConfig(".")
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

package server

import (
	"context"
	"log"

	"github.com/9bany/bot_workflows/src/cmds"
	db "github.com/9bany/bot_workflows/src/db/sqlc"
	"github.com/9bany/bot_workflows/src/utils"
	"github.com/shomali11/slacker"
)

type Server struct {
	slackBot *slacker.Slacker
	db       *db.Queries
}

func NewBot(config utils.Config, db *db.Queries) (*Server, error) {

	bot := slacker.NewClient(config.BotToken, config.AppToken)

	return &Server{
		slackBot: bot,
		db:       db,
	}, nil

}

func (server *Server) Start() error {

	server.initDefault()
	server.initCommands()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := server.slackBot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func (server *Server) initDefault() {
	server.slackBot.Init(func() {
		log.Println("Connected!")
	})

	server.slackBot.Err(func(err string) {
		log.Println(err)
	})

	server.slackBot.DefaultEvent(func(event interface{}) {
		log.Println(event)
	})

}

func (server *Server) initCommands() {
	relaxCommand := cmds.NewRelaxCommand(server.db)
	// commands
	server.slackBot.Command(cmds.PingCommandDefinition())
	server.slackBot.Command(relaxCommand.CommandDefinition())
	// help
	server.slackBot.Help(cmds.HelpCommandDefinition())
}

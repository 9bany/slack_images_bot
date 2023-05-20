package cmds

import (
	"bytes"
	"fmt"
	"log"

	db "github.com/9bany/bot_workflows/src/db/sqlc"
	"github.com/shomali11/slacker"
	"github.com/slack-go/slack"
)

const relaxCommandString = "relax"

type RelaxCommand struct {
	db *db.Queries
}

func NewRelaxCommand(db *db.Queries) *RelaxCommand {
	return &RelaxCommand{
		db: db,
	}
}

func (command *RelaxCommand) CommandDefinition() (string, *slacker.CommandDefinition) {
	definition := &slacker.CommandDefinition{
		Description: "Upload a sentence!",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {

			apiClient := botCtx.APIClient()
			event := botCtx.Event()

			image, err := command.db.GetRanDomImage(botCtx.Context())
			if err != nil {
				log.Println("can not get images from db")
			}

			log.Println("successfully get images fom db")

			if event.ChannelID != "" {
				_, err := apiClient.UploadFile(slack.FileUploadParameters{
					Reader:          bytes.NewBuffer(image.Photo),
					Title:           "hehe",
					Filename:        "images.jpg",
					Channels:        []string{event.ChannelID},
					ThreadTimestamp: event.ThreadTimeStamp,
				})
				if err != nil {
					fmt.Printf("Error encountered when uploading file: %+v\n", err)
				}
			}
		},
	}
	return relaxCommandString, definition
}

package cmds

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/9bany/bot_workflows/server/utils"
	"github.com/shomali11/slacker"
	"github.com/slack-go/slack"
)

var relaxCommandString = "relax"

func RelaxCommandDefinition() (string, *slacker.CommandDefinition) {
	definition := &slacker.CommandDefinition{
		Description: "Upload a sentence!",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {

			files, err := utils.FilePathWalkDir("./data/images")
			if err != nil {
				log.Println("Not any file in bot data warehouse")
			}

			index := rand.Intn(len(files))
			filePath := files[index]
			log.Println(filePath)
			apiClient := botCtx.APIClient()
			event := botCtx.Event()
			if event.ChannelID != "" {
				_, err := apiClient.UploadFile(slack.FileUploadParameters{
					File:            filePath,
					Title:           "Mlem mlem",
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

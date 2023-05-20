package cmds

import (
	"errors"
	"log"
	"strconv"

	"github.com/peterhellberg/hn"
	"github.com/shomali11/slacker"
	"github.com/slack-go/slack"
)

const newsCommandString = "news {count}"

type NewsCommand struct{}

func NewNewsCommand() *NewsCommand {
	return &NewsCommand{}
}

func (command *NewsCommand) newFeeds(count int) (list []*hn.Item) {
	hn := hn.DefaultClient

	ids, err := hn.TopStories()
	if err != nil {
		panic(err)
	}

	for _, id := range ids[:count] {
		item, err := hn.Item(id)
		if err != nil {
			log.Println(err)
			continue
		}
		list = append(list, item)
	}
	return
}
func (command *NewsCommand) CommandDefinition() (string, *slacker.CommandDefinition) {
	definition := &slacker.CommandDefinition{
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			count := request.Param("count")
			var countInt int
			var err error

			if count == "" {
				countInt = 10
			} else {
				// string to int
				countInt, err = strconv.Atoi(count)
				if err != nil {
					response.ReportError(errors.New("oops, {count} must be integer"))
					return
				}
			}

			if countInt < 0 {
				response.ReportError(errors.New("oops, {count} must be than 0"))
			}

			listFeeds := command.newFeeds(countInt)

			attachments := []slack.Attachment{}
			for _, elemet := range listFeeds {
				attachments = append(attachments, slack.Attachment{
					Color:      "#ffaa00",
					AuthorName: elemet.Title,
					AuthorLink: elemet.URL,
					Title:      elemet.By,
					Text:       elemet.Type,
				})
			}
			response.Reply("news report",
				slacker.WithAttachments(attachments),
				slacker.WithThreadReply(true),
			)
		},
	}
	return newsCommandString, definition
}

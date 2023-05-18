package cmds

import "github.com/shomali11/slacker"

var relaxCommandString = "relax"

func RelaxCommandDefinition() (string, *slacker.CommandDefinition) {
	definition := &slacker.CommandDefinition{
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			response.Reply("Are your sure ?")
		},
	}
	return relaxCommandString, definition
}

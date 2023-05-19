package cmds

import "github.com/shomali11/slacker"

func HelpCommandDefinition() *slacker.CommandDefinition {
	definition := &slacker.CommandDefinition{
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			response.Reply("Your own help function...")
		},
	}
	return definition
}

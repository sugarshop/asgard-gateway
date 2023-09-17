package commander

import (
	"log"

	"github.com/sashabaranov/go-openai"
	"github.com/sugarshop/asgard-gateway/discord/bot"
	"github.com/sugarshop/asgard-gateway/discord/commands"
	"github.com/sugarshop/env"
)

type DalleCommander struct {
	Cmd *bot.Command
}

func NewDalleCommander(dmPermission bool) *DalleCommander {
	openaiKey, ok := env.GlobalEnv().Get("OPENAIAPIKEY")
	if !ok {
		log.Println("no OPENAIAPIKEY env set")
	}
	openaiClient := openai.NewClient(openaiKey)
	return &DalleCommander{
		Cmd: commands.ImageCommand(&commands.ImageCommandParams{
			OpenAIClient: openaiClient,
			DMPermission: dmPermission,
		}),
	}
}

func (h *DalleCommander) Register(bot *bot.Bot) {
	bot.Router.Register(h.Cmd)
}

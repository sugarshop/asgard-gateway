package commander

import (
	"log"

	"github.com/sashabaranov/go-openai"
	"github.com/sugarshop/asgard-gateway/discord/bot"
	"github.com/sugarshop/asgard-gateway/discord/commands"
	"github.com/sugarshop/asgard-gateway/discord/commands/gpt"
	"github.com/sugarshop/asgard-gateway/discord/constants"
	"github.com/sugarshop/env"
)

type GPTCommander struct {
	Cmd *bot.Command
}

func NewGPTCommander(dmPermission bool) *GPTCommander {
	openaiKey, ok := env.GlobalEnv().Get("OPENAIAPIKEY")
	if !ok {
		log.Println("no OPENAIAPIKEY env set")
	}
	openaiClient := openai.NewClient(openaiKey)
	ignoredChannelsCache := &gpt.IgnoredChannelsCache{}
	// chat command dm supported
	gptMessagesCache, err := gpt.NewMessagesCache(constants.DiscordThreadsCacheSize)
	if err != nil {
		log.Fatalf("Error initializing GPTMessagesCache: %v", err)
	}
	return &GPTCommander{
		Cmd: commands.ChatCommand(&commands.ChatCommandParams{
			OpenAIClient: openaiClient,
			OpenAICompletionModels: []string{
				openai.GPT4,
				openai.GPT3Dot5Turbo,
				openai.GPT40314,
				openai.GPT3Dot5Turbo0301,
				"gpt-3.5-turbo-16k-0613",
				"gpt-3.5-turbo-16k",
			},
			GPTMessagesCache:     gptMessagesCache,
			IgnoredChannelsCache: ignoredChannelsCache,
			DMPermission:         dmPermission,
		}),
	}
}

func (h *GPTCommander) Register(bot *bot.Bot) {
	bot.Router.Register(h.Cmd)
}

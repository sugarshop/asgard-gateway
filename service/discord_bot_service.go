package service

import (
	"log"
	"sync"

	"github.com/sashabaranov/go-openai"
	"github.com/sugarshop/asgard-gateway/pkg/bot"
	"github.com/sugarshop/asgard-gateway/pkg/commands"
	"github.com/sugarshop/asgard-gateway/pkg/commands/gpt"
	"github.com/sugarshop/asgard-gateway/pkg/constants"
	"github.com/sugarshop/env"
)

// DiscordBotService a service running a discord bot.
type DiscordBotService struct {
	DiscordBot   *bot.Bot
	OpenaiClient *openai.Client

	GptMessagesCache     *gpt.MessagesCache
	IgnoredChannelsCache *gpt.IgnoredChannelsCache
}

var (
	discordBotService *DiscordBotService
	discordBotOnce    sync.Once
)

// DiscordBotServiceInstance initialize a discord bot instance
func DiscordBotServiceInstance() *DiscordBotService {
	openaiKey, ok := env.GlobalEnv().Get("OPENAIAPIKEY")
	if !ok {
		log.Println("no OPENAIAPIKEY env set")
	}
	discordBotToken, ok := env.GlobalEnv().Get("DISCORDBOTTOKEN")
	if !ok {
		log.Println("no DISCORDBOTTOKEN env set")
	}
	gptMessagesCache, err := gpt.NewMessagesCache(constants.DiscordThreadsCacheSize)
	if err != nil {
		log.Fatalf("Error initializing GPTMessagesCache: %v", err)
	}
	discordBot, err := bot.NewBot(discordBotToken)
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}
	openaiClient := openai.NewClient(openaiKey)
	ignoredChannelsCache := &gpt.IgnoredChannelsCache{}
	discordBot.Router.Register(commands.ChatCommand(&commands.ChatCommandParams{
		OpenAIClient: openaiClient,
		OpenAICompletionModels: []string{
			"gpt-4",
			"gpt-3.5-turbo",
			"gpt-4-0314",
			"gpt-3.5-turbo-0301",
		},
		GPTMessagesCache:     gptMessagesCache,
		IgnoredChannelsCache: ignoredChannelsCache,
	}))

	discordBot.Router.Register(commands.ImageCommand(openaiClient))

	// Run the bot
	discordBot.Run("", true)

	discordBotOnce.Do(func() {
		discordBotService = &DiscordBotService{
			DiscordBot:           discordBot,
			OpenaiClient:         openaiClient,
			GptMessagesCache:     gptMessagesCache,
			IgnoredChannelsCache: ignoredChannelsCache,
		}
	})
	return discordBotService
}

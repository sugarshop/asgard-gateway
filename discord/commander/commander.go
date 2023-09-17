package commander

import (
	"github.com/sugarshop/asgard-gateway/discord/bot"
)

// Commander Commander
type Commander interface {
	Register(*bot.Bot)
}

func commanders() []Commander {
	return []Commander{
		NewGPTCommander(true),  // GPT bot with DM permission
		NewGPTCommander(false), // GPT bot without permission
		NewDalleCommander(true),
	}
}

// Register Register all Command endpoints
func Register(bot *bot.Bot) {
	for _, co := range commanders() {
		co.Register(bot)
	}
}

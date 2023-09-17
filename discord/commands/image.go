package commands

import (
	discord "github.com/bwmarrin/discordgo"
	"github.com/sashabaranov/go-openai"
	"github.com/sugarshop/asgard-gateway/discord/bot"
	"github.com/sugarshop/asgard-gateway/discord/commands/dalle"
)

const imageCommandName = "image"

type ImageCommandParams struct {
	OpenAIClient *openai.Client
	DMPermission bool
}

func ImageCommand(params *ImageCommandParams) *bot.Command {
	return &bot.Command{
		Name:                     imageCommandName,
		Description:              "Generate creative images from textual descriptions",
		DMPermission:             params.DMPermission,
		DefaultMemberPermissions: discord.PermissionViewChannel,
		SubCommands: bot.NewRouter([]*bot.Command{
			dalle.Command(params.OpenAIClient),
		}),
	}
}

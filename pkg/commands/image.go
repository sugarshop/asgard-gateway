package commands

import (
	discord "github.com/bwmarrin/discordgo"
	"github.com/sashabaranov/go-openai"
	"github.com/sugarshop/asgard-gateway/pkg/bot"
	"github.com/sugarshop/asgard-gateway/pkg/commands/dalle"
)

const imageCommandName = "image"

func ImageCommand(client *openai.Client) *bot.Command {
	return &bot.Command{
		Name:                     imageCommandName,
		Description:              "Generate creative images from textual descriptions",
		DMPermission:             false,
		DefaultMemberPermissions: discord.PermissionViewChannel,
		SubCommands: bot.NewRouter([]*bot.Command{
			dalle.Command(client),
		}),
	}
}

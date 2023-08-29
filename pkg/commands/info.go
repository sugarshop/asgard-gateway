package commands

import (
	discord "github.com/bwmarrin/discordgo"
	"github.com/sugarshop/asgard-gateway/pkg/bot"
	"github.com/sugarshop/asgard-gateway/pkg/constants"
)

const (
	infoCommandName = "info"
)

func infoHandler(ctx *bot.Context) {
	ctx.Respond(&discord.InteractionResponse{
		Type: discord.InteractionResponseChannelMessageWithSource,
		Data: &discord.InteractionResponseData{
			// Note: only visible to the user who invoked the command
			Flags: discord.MessageFlagsEphemeral,
			// Content: "Surprise!",
			Components: []discord.MessageComponent{
				discord.ActionsRow{
					Components: []discord.MessageComponent{
						&discord.Button{
							Label: "Source code",
							Style: discord.LinkButton,
							URL:   "https://github.com/sugarshop/go-remai-bot-discord",
						},
					},
				},
			},
			Embeds: []*discord.MessageEmbed{
				{
					Title:       "Bot Version",
					Description: "Version: " + constants.Version,
					Color:       0x00bfff,
				},
			},
		},
	})
}

func InfoCommand() *bot.Command {
	return &bot.Command{
		Name:                     infoCommandName,
		Description:              "Show information about current version of Rem AI",
		DMPermission:             true,
		DefaultMemberPermissions: discord.PermissionViewChannel,
		Handler:                  bot.HandlerFunc(infoHandler),
	}
}

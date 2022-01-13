package commandHandlers

import (
	"github.com/bwmarrin/discordgo"
)

func HandlePingCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Author:      &discordgo.MessageEmbedAuthor{},
					Color:       0x00ff00,
					Description: "This is a description",
				},
			},
		},
	})
}

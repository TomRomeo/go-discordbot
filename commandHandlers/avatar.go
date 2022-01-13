package commandHandlers

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func HandleAvatarCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {

	var user *discordgo.User
	// check if option submitted
	if len(i.ApplicationCommandData().Options) >= 1 {
		user = i.ApplicationCommandData().Options[0].UserValue(s)
	} else {
		user = i.Member.User
	}

	avatarUrl := user.AvatarURL("128")

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Author:      &discordgo.MessageEmbedAuthor{},
					Color:       0x5865F2,
					Description: fmt.Sprintf("Avatar of %s", user.Mention()),
					Image: &discordgo.MessageEmbedImage{
						URL: avatarUrl,
					},
				},
			},
		},
	})
}

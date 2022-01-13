package commandHandlers

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"math/rand"
	"os"
	"time"
)

func HandleCoinflipCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {

	decisionText := [2]string{"Head", "Tails"}
	decisionsImagePaths := [2]string{"./assets/head.png", "./assets/tails.png"}

	rand.Seed(time.Now().Unix())
	randint := rand.Intn(2)

	// open a random image
	image, err := os.Open(decisionsImagePaths[randint])
	if err != nil {
		log.Fatalf("Couldn't open image file: %v", err)
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Author: &discordgo.MessageEmbedAuthor{},
					Color:  0x00ff00,
					Title:  fmt.Sprintf("You rolled %s!", decisionText[randint]),
					Image: &discordgo.MessageEmbedImage{
						URL:    "attachment://decision.png",
						Width:  128,
						Height: 128,
					},
				},
			},
			Files: []*discordgo.File{
				{
					Name:   "decision.png",
					Reader: image,
				},
			},
		},
	})
}

package main

import (
	"awesomeProject1/commandHandlers"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not parse the .env file: ", err)
	}

	dg, err := discordgo.New("Bot " + os.Getenv("BOT_KEY"))
	if err != nil {
		log.Fatalf("An Error occured while creating the bot object: ", err)
		return
	}

	err = dg.Open()
	if err != nil {
		log.Fatalf("Something went wrong while trying to connect to Discord...", err)
		return
	}

	var commands = []*discordgo.ApplicationCommand{
		{
			Name:        "ping",
			Description: "Test ping command",
		},
		{
			Name:        "coinflip",
			Description: "Flip a coin!",
		},
		{
			Name:        "avatar",
			Description: "Display the avatar of a user",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionUser,
					Name:        "username",
					Description: "The username to show the avatar for",
					Required:    false,
				},
			},
		},
	}

	// register the commands
	for _, v := range commands {
		_, err := dg.ApplicationCommandCreate(dg.State.User.ID, "", v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
	}
	log.Println("Registered all Commands successfully...")
	log.Println("Bot is running. Press CTRL-C to exit.")

	dg.AddHandler(commandHandler)

	// graceful exit logic
	gracefulExit(dg)
}

func gracefulExit(dg *discordgo.Session) {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	log.Println("Gracefully shutting down the discord bot")
	dg.Close()

}

func commandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {

	if i.Type != discordgo.InteractionApplicationCommand {
		return
	}
	data := i.ApplicationCommandData()

	switch data.Name {
	case "ping":
		commandHandlers.HandlePingCommand(s, i)

	case "avatar":
		commandHandlers.HandleAvatarCommand(s, i)

	case "coinflip":
		commandHandlers.HandleCoinflipCommand(s, i)
	}
}

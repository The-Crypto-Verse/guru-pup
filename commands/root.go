package commands

import (
	"log"
	"os"
	"strings"

	_ "github.com/joho/godotenv/autoload"

	"github.com/TheBoringDude/minidis"
	"github.com/bwmarrin/discordgo"
)

func Manager() *minidis.Minidis {
	token := os.Getenv("TOKEN")
	guilds := strings.Split(os.Getenv("GUILD"), ",")

	// check if token is not empty
	if token == "" {
		log.Fatal("No TOKEN provided!")
	}

	bot := minidis.New(token)

	// set intents
	bot.SetIntents(discordgo.IntentsGuilds | discordgo.IntentsGuildMessages)

	// sync only to specific servers
	bot.SyncToGuilds(guilds...)

	bot.OnReady(func(s *discordgo.Session, i *discordgo.Ready) {
		log.Println("Bot is ready!")
	})

	bot.OnBeforeStart(func(s *discordgo.Session) {
		// try to remove the commands first to prevent duplicate commands
		if err := bot.ClearCommands(); err != nil {
			log.Fatal(err)
		}
	})

	bot.OnClose(func(s *discordgo.Session) {
		log.Println("Closing...")
	})

	// register commands in here
	bot.RegisterCommands(
		RegisterCommand,
		ProfileCommand,
		HelpCommand,
	)

	return bot
}

package commands

import (
	"log"

	"github.com/The-Crypto-Verse/guru-pup/lib"
	"github.com/TheBoringDude/minidis"
	"github.com/bwmarrin/discordgo"
)

var Bot *minidis.Minidis

func init() {
	// check if token is not empty
	if lib.TOKEN == "" {
		log.Fatal("No TOKEN provided!")
	}

	Bot = minidis.New(lib.TOKEN)

	// sync only to specific servers
	Bot.SyncToGuilds(lib.GUILDS...)

	Bot.OnReady(func(s *discordgo.Session, i *discordgo.Ready) {
		log.Println("Bot is ready!")
	})

	Bot.OnBeforeStart(func(s *discordgo.Session) {
		// try to remove the commands first to prevent duplicate commands
		if err := Bot.ClearCommands(); err != nil {
			log.Fatal(err)
		}
	})

	Bot.OnClose(func(s *discordgo.Session) {
		log.Println("Closing...")
	})

	// register commands in here
	Bot.RegisterCommands(HelpCommand, RegisterCommand, ProfileCommand)
}

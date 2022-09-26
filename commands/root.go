package commands

import (
	"log"

	"github.com/The-Crypto-Verse/guru-pup/lib"
	"github.com/TheBoringDude/minidis"
	"github.com/bwmarrin/discordgo"
)

func Manager() *minidis.Minidis {
	// check if token is not empty
	if lib.TOKEN == "" {
		log.Fatal("No TOKEN provided!")
	}

	bot := minidis.New(lib.TOKEN)

	// set intents
	bot.SetIntents(discordgo.IntentsGuilds | discordgo.IntentsGuildMessages)

	// sync only to specific servers
	bot.SyncToGuilds(lib.GUILDS...)

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

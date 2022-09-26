package main

import (
	_ "github.com/joho/godotenv/autoload"

	"log"

	"github.com/The-Crypto-Verse/guru-pup/commands"
)

func main() {
	bot := commands.Manager()

	if err := bot.Run(); err != nil {
		log.Fatal(err)
	}
}

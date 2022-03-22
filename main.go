package main

import (
	"log"

	"github.com/The-Crypto-Verse/guru-pup/commands"
)

func main() {
	bot := commands.Manager()

	if err := bot.Run(); err != nil {
		log.Fatal(err)
	}
}

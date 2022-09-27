package main

import (
	_ "github.com/joho/godotenv/autoload"

	"log"

	"github.com/The-Crypto-Verse/guru-pup/commands"
	"github.com/TheBoringDude/minidis"
)

func main() {
	if err := minidis.Execute(commands.Bot); err != nil {
		log.Fatal(err)
	}
}

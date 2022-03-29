package commands

import (
	"time"

	"github.com/TheBoringDude/minidis"
	"github.com/bwmarrin/discordgo"
)

var HelpCommand = &minidis.SlashCommandProps{
	Command:     "help",
	Description: "Show help and more info about Guru Pup",
	Execute: func(c *minidis.SlashContext) error {
		embed := &discordgo.MessageEmbed{
			Title: "Help | Guru Pup",
			Author: &discordgo.MessageEmbedAuthor{
				IconURL: c.Bot.AvatarURL(""),
				Name:    c.Bot.Username,
			},
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: c.Bot.AvatarURL(""),
			},
			Description: "Hi I am Guru Pup, a discord bot helper for The CryptoVerse.",
			Fields: []*discordgo.MessageEmbedField{
				{
					Name: ":robot: Commands",
					Value: `The following commands are available:
	**/register**
	- Link you wax wallet to your discord profile
	**/profile**
	- Check your profile, show your WAX and KASU balance and total assets owned for the collection
	**/help**
	- Show this help message.
					`,
				},
				{
					Name:  "\u200b",
					Value: "\u200b",
				},
				{
					Name:  ":globe_with_meridians: Website",
					Value: "https://thecryptopups.io/",
				},
			},
			Footer: &discordgo.MessageEmbedFooter{
				Text: "Guru Pup | CC 2022",
			},
			Timestamp: time.Now().Format(time.RFC3339),
		}

		return c.ReplyC(minidis.ReplyProps{
			Embeds: []*discordgo.MessageEmbed{
				embed,
			},
			IsEphemeral: false,
		})
	},
}

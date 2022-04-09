package commands

import (
	"fmt"
	"strconv"
	"time"

	"github.com/The-Crypto-Verse/guru-pup/lib"
	"github.com/The-Crypto-Verse/guru-pup/types"
	"github.com/TheBoringDude/minidis"
	"github.com/bwmarrin/discordgo"
)

var ProfileCommand = &minidis.SlashCommandProps{
	Command:     "profile",
	Description: "View your profile",
	Execute: func(c *minidis.SlashContext) error {
		c.DeferReply(true)

		base, err := lib.UsersBase()
		if err != nil {
			_, err := c.Followup("DB Connection error, please contact the developer regarding this problem.")
			return err
		}

		var currentUser = types.UserProps{
			Wallet: "",
			ID:     "",
		}
		if err := base.Get(c.Author.ID, &currentUser); err != nil {
			_, err := c.Followup("You are currently not registered yet.")
			return err
		}

		bals, err := lib.GetTokenBalance(currentUser.Wallet)
		if err != nil {
			_, err := c.Followup("Failed to retrieve your wallet's balance.")
			return err
		}

		assets, err := lib.CountAssetStats(currentUser.Wallet)
		if err != nil {
			_, err := c.Followup("Failed to retrieve your assets stats.")
			return err
		}

		embed := &discordgo.MessageEmbed{
			Title: fmt.Sprintf("Profile | %s", c.Author.String()),
			Author: &discordgo.MessageEmbedAuthor{
				IconURL: c.Author.AvatarURL(""),
				Name:    c.Author.Username,
			},
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: c.Author.AvatarURL(""),
			},
			Timestamp:   time.Now().Format(time.RFC3339),
			Description: "Your current profile (TBU)",
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:  "Wallet",
					Value: currentUser.Wallet,
				},
				{
					Name:   "WAX Balance",
					Value:  bals.Wax,
					Inline: true,
				},
				{
					Name:   "KASU Balance",
					Value:  bals.Kasu,
					Inline: true,
				},
				{
					Name:   "Total Assets",
					Value:  strconv.Itoa(assets),
					Inline: true,
				},
			},
			Footer: &discordgo.MessageEmbedFooter{
				Text: "Guru Pup | CC 2022",
			},
		}

		return c.EditC(minidis.EditProps{
			Content: "",
			Embeds:  []*discordgo.MessageEmbed{embed},
		})
	},
}

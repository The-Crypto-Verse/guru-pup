package commands

import (
	"github.com/The-Crypto-Verse/guru-pup/lib"
	"github.com/The-Crypto-Verse/guru-pup/types"
	"github.com/TheBoringDude/minidis"
	"github.com/bwmarrin/discordgo"
)

var RegisterCommand = &minidis.SlashCommandProps{
	Name:        "register",
	Description: "Register your wallet to the bot.",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionString,
			Name:        "wallet",
			Description: "Your wax wallet address,",
			Required:    true,
		},
	},
	Execute: func(c *minidis.SlashContext) error {
		wallet := c.Options["wallet"].StringValue()
		id := c.Author.ID

		c.DeferReply(true)

		// get users deta base
		base, err := lib.UsersBase()
		if err != nil {
			_, err := c.Followup("DB Connection error, please contact the developer regarding this problem.")
			return err
		}

		var currentUser = types.UserProps{
			Wallet: "",
			ID:     "",
		}
		base.Get(id, &currentUser)

		// check if currentUser exists
		if currentUser.ID != "" {
			_, err := c.Followup("You are already registered!")
			return err
		}

		stats, err := lib.GetAssetStats(wallet)
		if err != nil {
			return c.Edit("Failed to get wallet stats. Please try again later.")
		}

		roles := lib.GetRoles(stats)
		for _, v := range roles {
			if err = c.Session.GuildMemberRoleAdd(c.GuildId, id, v); err != nil {
				return c.Edit("Failed to give roles to user. Please report this error to the admins.")
			}
		}

		// put data into the db
		if _, err = base.Put(&types.UserProps{
			Key:    id,
			ID:     id,
			Wallet: wallet,
		}); err != nil {
			return c.Edit("I failed to register your wallet. Please try again later.")
		}

		return c.Edit("You have successfully registered your wallet! You can now try the other commands.")
	},
}

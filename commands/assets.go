package commands

import (
	"fmt"

	"github.com/The-Crypto-Verse/guru-pup/lib"
	"github.com/The-Crypto-Verse/guru-pup/types"
	"github.com/TheBoringDude/minidis"
)

var AssetsCommand = &minidis.SlashCommandProps{
	Command:     "assets",
	Description: "View your total assets overview",
	Execute: func(c *minidis.SlashContext) error {
		c.DeferReply(true)

		base, err := lib.UsersBase()
		if err != nil {
			_, err := c.Followup("DB Connection error, please contact the developer regarding this problem.")
			return err
		}

		var currentUser types.UserProps
		if err := base.Get(c.Author.ID, &currentUser); err != nil {
			_, err := c.Followup("You are currently not registered yet.")
			return err
		}

		assetStats, err := lib.GetAssetStats(currentUser.Wallet)

		for _, j := range assetStats.Data.Templates {
			// should get the template in here
			fmt.Println(j.TemplateID)
		}

		return c.ReplyString("this is assets content")
	},
}

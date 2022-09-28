package admin

import (
	"strings"

	"github.com/The-Crypto-Verse/guru-pup/lib"
	"github.com/TheBoringDude/minidis"
	"github.com/bwmarrin/discordgo"
)

func hasRole(roles []string, role string) bool {
	for _, v := range roles {
		if v == role {
			return true
		}
	}

	return false
}

var GetWalletsCommand = &minidis.SlashCommandProps{
	Name:        "get-wallets",
	Description: "Get the wallets of registered users. (admin-only)",
	Execute: func(c *minidis.SlashContext) error {
		c.DeferReply(true)

		// https://github.com/bwmarrin/discordgo/issues/1024
		perms, err := c.Session.UserChannelPermissions(c.Author.ID, c.ChannelId)
		if err != nil {
			_, err := c.Followup("There was a problem getting the user's permissions.")
			return err
		}
		if perms&discordgo.PermissionAdministrator == 0 {
			// not admin
			return c.Edit("You do not have permission to perform such actions!")
		}

		users, err := lib.GetAllUser()
		if err != nil {
			return c.Edit("Failed to get all users.")
		}

		shrimpyPups := []string{}
		crabbyPups := []string{}
		octoPups := []string{}
		fishPups := []string{}
		dolphinPups := []string{}
		sharkPups := []string{}
		whalePups := []string{}

		for _, v := range users {
			member, err := c.Session.GuildMember(c.GuildId, v.ID)
			if err != nil {
				continue
			}

			if hasRole(member.Roles, lib.Shrimpy) {
				shrimpyPups = append(shrimpyPups, v.Wallet)
			}
			if hasRole(member.Roles, lib.Crabby) {
				crabbyPups = append(crabbyPups, v.Wallet)
			}
			if hasRole(member.Roles, lib.Octo) {
				octoPups = append(octoPups, v.Wallet)
			}
			if hasRole(member.Roles, lib.Fish) {
				fishPups = append(fishPups, v.Wallet)
			}
			if hasRole(member.Roles, lib.Dolphin) {
				dolphinPups = append(dolphinPups, v.Wallet)
			}
			if hasRole(member.Roles, lib.Shark) {
				sharkPups = append(sharkPups, v.Wallet)
			}
			if hasRole(member.Roles, lib.Whale) {
				whalePups = append(whalePups, v.Wallet)
			}

		}

		return c.EditC(minidis.EditProps{
			Content: "The following are the registered wallets to me...",
			Attachments: []*discordgo.File{
				{
					ContentType: "text/plain",
					Name:        "Shrimpy Pups.txt",
					Reader:      strings.NewReader("-- Shrimpy Pups -- \n\n" + strings.Join(shrimpyPups, "\n")),
				},
				{
					ContentType: "text/plain",
					Name:        "Crabby Pups.txt",
					Reader:      strings.NewReader("-- Crabby Pups -- \n\n" + strings.Join(crabbyPups, "\n")),
				},
				{
					ContentType: "text/plain",
					Name:        "Octo Pups.txt",
					Reader:      strings.NewReader("-- Octo Pups -- \n\n" + strings.Join(octoPups, "\n")),
				},
				{
					ContentType: "text/plain",
					Name:        "Fish Pups.txt",
					Reader:      strings.NewReader("-- Fish Pups -- \n\n" + strings.Join(fishPups, "\n")),
				},
				{
					ContentType: "text/plain",
					Name:        "Dolphin Pups.txt",
					Reader:      strings.NewReader("-- Dolphin Pups -- \n\n" + strings.Join(dolphinPups, "\n")),
				},
				{
					ContentType: "text/plain",
					Name:        "Sharko Pups.txt",
					Reader:      strings.NewReader("-- Sharko Pups -- \n\n" + strings.Join(sharkPups, "\n")),
				},
				{
					ContentType: "text/plain",
					Name:        "Whale Pups.txt",
					Reader:      strings.NewReader("-- Whale Pups -- \n\n" + strings.Join(whalePups, "\n")),
				},
			},
		})
	},
}

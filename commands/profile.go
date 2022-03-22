package commands

import "github.com/TheBoringDude/minidis"

var ProfileCommand = &minidis.SlashCommandProps{
	Command:     "profile",
	Description: "View your profile",
	Execute: func(c *minidis.SlashContext) error {
		return c.Reply("/your profile should be in here")
	},
}

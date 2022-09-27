package worker

import (
	"fmt"
	"log"
	"time"

	"github.com/The-Crypto-Verse/guru-pup/lib"
	"github.com/bwmarrin/discordgo"
)

func Start(session *discordgo.Session, guildId string) {
	for {
		log.Println("- starting worker -")

		// get all users
		allUsers, err := lib.GetAllUser()
		if err != nil {
			log.Println(err)
		}

		for _, v := range allUsers {
			fmt.Printf("Handling user: %s (%s) \n", v.ID, v.Wallet)

			// check if member is in server
			_, err := session.GuildMember(guildId, v.ID)
			if err != nil {
				fmt.Printf("Member: %s(%s) is not in server \n", v.ID, v.Wallet)
				continue
			}

			stats, err := lib.GetAssetStats(v.Wallet)
			if err != nil {
				log.Println(err)
				continue
			}

			roles := lib.GetRoles(stats)
			for _, x := range roles {
				// give roles
				if err = session.GuildMemberRoleAdd(guildId, v.ID, x); err != nil {
					fmt.Printf("Failed to give role to %s (err: %v)\n", v.Wallet, err)
				}
			}
		}

		// sleep for 1 minute
		time.Sleep(time.Duration(1) * time.Minute)
	}
}

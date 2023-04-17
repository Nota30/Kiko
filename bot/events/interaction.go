package events

import (
	modules "github.com/Nota30/Kiko/modules/Information"
	"github.com/bwmarrin/discordgo"
)

func InteractionHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Member.User.Bot {
		return
	}

	switch i.ApplicationCommandData().Name {
	case "ping":
		modules.Ping(s, i)
	case "interactionroles":
		modules.InteractionRoles(s, i)
	case "register":
		modules.Register(s, i)
	}

}

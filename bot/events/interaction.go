package events

import (
	modules "github.com/Nota30/Kiko/modules/Information"
	"github.com/bwmarrin/discordgo"
)

func InteractionHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Member.User.Bot {
		return
	}

	if i.Data.Type() != discordgo.InteractionApplicationCommand {
		return
	}

	switch i.ApplicationCommandData().Name {
	case "ping":
		modules.Ping(s, i)
	case "register":
		modules.Register(s, i)
	}

}

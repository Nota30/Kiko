package events

import (
	info "github.com/Nota30/Kiko/modules/Information"
	tower "github.com/Nota30/Kiko/modules/Tower"
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
		info.Ping(s, i)
	case "register":
		info.Register(s, i)
	case "profile":
		info.Profile(s, i)
	case "move":
		tower.Move(s, i)
	case "explore":
		tower.Explore(s, i)
	}
}

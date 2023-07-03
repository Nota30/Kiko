package events

import (
	information "github.com/Nota30/Kiko/modules/Information"
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
		information.PingCMD(s, i)
	case "register":
		information.RegisterCMD(s, i)
	case "profile":
		information.ProfileCMD(s, i)
	case "move":
		tower.MoveCMD(s, i)
	case "explore":
		tower.ExploreCMD(s, i)
	}
}

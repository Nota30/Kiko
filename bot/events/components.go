package events

import (
	modules "github.com/Nota30/Kiko/modules/Information"
	"github.com/bwmarrin/discordgo"
)

func ComponentHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Member.User.Bot {
		return
	}

	if i.Data.Type() != discordgo.InteractionMessageComponent {
		return
	}

	// 3 is select menu type
	if i.MessageComponentData().ComponentType == 3 {
		switch i.MessageComponentData().CustomID {
			case "select_class":
				modules.RegisterSelector(s, i)
		}
	}
}
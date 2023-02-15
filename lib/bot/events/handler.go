package events

import (
	"fmt"

	"github.com/Nota30/Kiko/tools"
	"github.com/bwmarrin/discordgo"
)

func EventHandler(s *discordgo.Session, evt interface{}) {
	switch evt.(type) {
	case *discordgo.MessageCreate:
		event, ok := evt.(*discordgo.MessageCreate)
		if ok {
			MSGHandler(s, event)
		}
	case *discordgo.InteractionCreate:
		fmt.Println("A Interaction happened!")
	}
}

func interactionCreate(s *discordgo.Session,  m *discordgo.InteractionCreate) {
	tools.Respond(s, m, "Pong!")
}
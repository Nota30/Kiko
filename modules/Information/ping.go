package modules

import (
	"github.com/Nota30/Kiko/config"
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

func Ping(s *discordgo.Session, i *discordgo.InteractionCreate) {
	logrus.Info(config.Classes)
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Pong!",
		},
	})
}

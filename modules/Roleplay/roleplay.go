package roleplay

import (
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

func Actions (s *discordgo.Session, m *discordgo.MessageCreate, action string) {
	logrus.Info(action)
}


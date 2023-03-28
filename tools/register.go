package tools

import (
	"strconv"

	"github.com/Nota30/Kiko/config"
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)


func RegisterCommands(s *discordgo.Session, GuildId string) {
	commands := config.Commands
	for i, v := range commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, GuildId, v)

		if err != nil {
			logrus.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		
		logrus.Info("Added Command: "+ cmd.Name + " [" + strconv.Itoa(i) + "]")
	}
}
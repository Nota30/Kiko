package events

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

func MSGCreateHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if m.Author.Bot {
		return
	}
	channel, err := s.Channel(m.ChannelID)
	if err != nil {
		return
	}
	if channel.Type == 1 {
		return
	}

	commands := []string{"help"}
	prefix := "k."

	hasPrefix := strings.HasPrefix(m.Content, prefix)
	if !hasPrefix {
		return
	}

	args := getCommand(m.Content, len(prefix))
	cmd := strings.Split(strings.ToLower(args), " ")
	exists := stringInSlice(cmd[0], commands)

	if !exists {
		return
	}

	logrus.Info(exists)
}

func getCommand(str string, n int) string {
	if len(str) >= n {
		return str[n:]
	}
	return str[:0]
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

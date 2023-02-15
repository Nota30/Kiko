package events

import (
	"strings"

	roleplay "github.com/Nota30/Kiko/modules/Roleplay"
	"github.com/bwmarrin/discordgo"
)

func MSGHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if m.Author.Bot {
		return
	}

	commands := []string{"bully", "cuddle", "hug", "kiss", "lick", "pat", "bonk", "yeet", "highfive", "handhold", "bite", "slap", "kill", "kick", "poke"}
	prefix := "k."

	hasPrefix := strings.HasPrefix(m.Content, prefix)
	if !hasPrefix {
		return
	}

	args := getCommand(m.Content, 2)
	cmd := strings.Split(strings.ToLower(args), " ")
	exists := stringInSlice(cmd[0], commands)

	if !exists {
		return
	}

	roleplay.Actions(s, m, cmd[0])
}

func getCommand(str string, n int) string {
	m := 0
    for i := range str {
        if m >= n {
            return str[i:]
        }
        m++
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
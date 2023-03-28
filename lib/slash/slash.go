package slash

import (
	"fmt"
	"strconv"

	"github.com/Nota30/Kiko/config"
	"github.com/Nota30/Kiko/tools"
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

func UpdateSlashCMDS() {
	s, err := discordgo.New(tools.GetEnv("DISCORD_TOKEN"))
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	err = s.Open()
	if err != nil {
		logrus.Fatalf("Cannot open the session: %v", err)
	}

	logrus.Info("Registering Commands...")

	commands := config.Commands
	for i, v := range commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, "983931249456455720", v)

		if err != nil {
			logrus.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}

		logrus.Info("Added Command: " + cmd.Name + " [" + strconv.Itoa(i) + "]")
	}

	defer s.Close()
}

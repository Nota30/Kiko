package main

import (
	"fmt"

	"github.com/Nota30/Kiko/config"
	"github.com/Nota30/Kiko/tools"
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

func main() {
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

	tools.RegisterCommands(s, config.DevGuild)

	defer s.Close()
}
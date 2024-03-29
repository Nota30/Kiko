package bot

import (
	"os"

	"github.com/Nota30/Kiko/bot/events"
	"github.com/Nota30/Kiko/config"
	"github.com/Nota30/Kiko/tools"
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

var (
	Dg *discordgo.Session
)

func DiscordConnect() {
	logrus.Info("Starting connection to Discord.")
	var err error
	Dg, err = discordgo.New(os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		logrus.Error("There was an error creating the discord session,", err)
		return
	}

	Dg.AddHandler(events.MSGCreateHandler)
	Dg.AddHandler(events.InteractionHandler)
	Dg.AddHandler(events.ComponentHandler)
	Dg.AddHandler(events.ActivityHandler)

	Dg.Identify.Intents |= discordgo.IntentMessageContent
	Dg.Identify.Intents |= discordgo.IntentsGuildMessages
	Dg.Identify.Intents |= discordgo.IntentsGuilds
	Dg.Identify.Intents |= discordgo.IntentsGuildMembers

	Dg.StateEnabled = false
	Dg.LogLevel = discordgo.LogInformational
	Dg.SyncEvents = true

	err = Dg.Open()
	if err != nil {
		logrus.Error("There was an error in opening the discord session,", err)
	}

	logrus.Info("Connected to Discord!")

	apps, _ := Dg.ApplicationCommands(Dg.State.Application.ID, "")
	config.ApplicationCommands = apps

	for _, val := range apps {
		if val.Name == "register" {
			config.RegisterCommand = "</register:" + val.ID + ">"
		}
	}

	env := os.Getenv("env")
	if env == "development" {
		logrus.Info("Registering Slash Commands in dev mode...")
		tools.RegisterCommands(Dg, config.DevGuild)
	} else if env == "production" {
		logrus.Info("Registering Slash Commands in prod mode...")
		tools.RegisterCommands(Dg, "")
	}
}

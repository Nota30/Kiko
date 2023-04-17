package bot

import (
	"github.com/Nota30/Kiko/bot/events"
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
	Dg, err = discordgo.New(tools.GetEnv("DISCORD_TOKEN"))
	if err != nil {
		logrus.Error("There was an error creating the discord session,", err)
		return
	}
	
	Dg.AddHandler(events.MSGCreateHandler)
	Dg.AddHandler(events.InteractionHandler)

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
	// logrus.Info("Registering Slash Commands...")
	// tools.RegisterCommands(Dg, config.DevGuild)
}

// var (
// 	Sharder *dshardmanager.Manager
// )

// var (
// 	shardCount int
// )

// func DiscordConnect() {
// 	shardManger()
// 	logrus.Info("Starding Sharder")
// 	Sharder.Init()
// 	go Sharder.Start()
// }

// func shardManger() {
// 	Sharder = dshardmanager.New(tools.GetEnv("DISCORD_TOKEN"))
// 	Sharder.Name = "Kiko"
// 	Sharder.LogChannel = tools.GetEnv("LOG_CHANNEL")
// 	Sharder.StatusMessageChannel = tools.GetEnv("STATUS_CHANNEL")

// 	recommended, err := Sharder.GetRecommendedCount()
// 	if err != nil {
// 		logrus.Error("Failed getting recommended shard count")
// 	}

// 	logrus.Info("ShardCount:", recommended)

// 	Sharder.SessionFunc = func(token string) (session *discordgo.Session, err error) {
// 		session, err = discordgo.New(token)
// 		if err !=nil {
// 			logrus.Error("[ShardManager]", "error opening connection", err)
// 			return
// 		}

// 		session.StateEnabled = false
// 		session.LogLevel = discordgo.LogInformational
// 		session.SyncEvents = true
// 		session.Identify.Intents |= discordgo.IntentMessageContent

// 		return
// 	}

// 	Sharder.AddHandler(events.EventHandler)
// }

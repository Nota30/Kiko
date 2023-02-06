package bot

import (
	"github.com/Nota30/Kiko/tools"
	"github.com/bwmarrin/discordgo"
	"github.com/jonas747/dshardmanager"
	"github.com/sirupsen/logrus"
)

var (
	Sharder *dshardmanager.Manager
)

var (
	shardCount int
)

func DiscordConnect () {
	shardManger()
	logrus.Info("Starding Sharder")
	Sharder.Init()
	go Sharder.Start()
}

func shardManger() {
	Sharder = dshardmanager.New(tools.GetEnv("DISCORD_TOKEN"))
	Sharder.Name = "Kiko"
	Sharder.LogChannel = "986122070607949834"
	Sharder.StatusMessageChannel = "986122070607949834"

	recommended, err := Sharder.GetRecommendedCount()
	if err != nil {
		logrus.Error("Failed getting recommended shard count")
	}

	logrus.Info("ShardCount:", recommended)

	Sharder.SessionFunc = func(token string) (session *discordgo.Session, err error) {
		session, err = discordgo.New(token)
		if err !=nil {
			logrus.Error("[ShardManager]", "error opening connection", err)
			return
		}

		session.StateEnabled = false
		session.LogLevel = discordgo.LogInformational
		session.SyncEvents = true
		session.Identify.Intents |= discordgo.IntentMessageContent

		return
	}

	Sharder.AddHandler(interactionCreate)
}

func interactionCreate(s *discordgo.Session,  m *discordgo.InteractionCreate) {
	tools.Respond(s, m, "Pong!")
}
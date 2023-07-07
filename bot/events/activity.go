package events

import (
	"time"

	"github.com/Nota30/Kiko/config"
	"github.com/bwmarrin/discordgo"
)

type ChannelActivity struct {
	MessageCount int
	LastUpdate   time.Time
	LastAuthor   string
	LastSpawn    time.Time
}

var (
	channelActivityMap = make(map[string]*ChannelActivity)
	spawnCooldown      = 1 * time.Minute
)

const (
	activityThreshold  = 30
	activityTimeWindow = 1 * time.Hour
)

func ActivityHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.Bot || message.Author.ID == session.State.User.ID {
		return
	}

	channelID := message.ChannelID
	activity, exists := channelActivityMap[channelID]

	if !exists {
		activity = &ChannelActivity{
			MessageCount: 1,
			LastUpdate:   time.Now(),
			LastAuthor:   session.State.User.ID,
			LastSpawn:    time.Now(),
		}
		channelActivityMap[channelID] = activity
		return
	}

	if time.Since(activity.LastUpdate) > activityTimeWindow {
		activity.MessageCount = 1
		activity.LastUpdate = time.Now()
		activity.LastAuthor = session.State.User.ID
		return
	}
	if session.State.User.ID != activity.LastAuthor {
		activity.MessageCount++
		activity.LastUpdate = time.Now()
		activity.LastAuthor = session.State.User.ID
	}

	if activity.MessageCount >= activityThreshold {
		if time.Since(activity.LastSpawn) > spawnCooldown {
			activity.LastSpawn = time.Now()
			enemy := config.GetRandomEnemy()
			enemyEmbed := config.GetEnemyEmbed(enemy)
			session.ChannelMessageSendEmbed(channelID, enemyEmbed)
		}
	}
}

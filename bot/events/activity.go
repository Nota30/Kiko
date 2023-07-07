package events

import (
	"fmt"
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

// Mess with the variables and constants to adjust intervals of spawning and activity thresholds
var (
	channelActivityMap = make(map[string]*ChannelActivity)
	spawnCooldown      = 1 * time.Minute
)

const (
	activityThreshold  = 3
	activityTimeWindow = 2 * time.Minute
)

func ActivityHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.Bot || message.Author.ID == session.State.User.ID {
		return
	}

	channelID := message.ChannelID
	activity, exists := channelActivityMap[channelID]

	if !exists {
		fmt.Printf("Created channel activity map!!!\n") // Debug line
		activity = &ChannelActivity{
			MessageCount: 1,
			LastUpdate:   time.Now(),
			LastAuthor:   message.Author.ID,
			LastSpawn:    time.Now(),
		}
		channelActivityMap[channelID] = activity
		return
	}

	if time.Since(activity.LastUpdate) > activityTimeWindow {
		fmt.Printf("Time exceeded, reset activity\n") // Debug line
		activity.MessageCount = 1
		activity.LastUpdate = time.Now()
		activity.LastAuthor = message.Author.ID
		return
	}

	// Comment this if statement if you want to test the system alone (otherwise it'll think you are spamming and won't work)
	//if session.State.User.ID != activity.LastAuthor {
	activity.MessageCount++
	activity.LastUpdate = time.Now()
	activity.LastAuthor = message.Author.ID
	//}

	// Debug lines
	fmt.Printf("Message count: " + fmt.Sprint(activity.MessageCount) + "\n")
	fmt.Printf("Last Update: " + fmt.Sprint(activity.LastUpdate) + "\n")
	fmt.Printf("Last Author: " + fmt.Sprint(activity.LastAuthor) + "\n")
	fmt.Printf("Last Spawn: " + fmt.Sprint(activity.LastSpawn) + "\n")
	// Debug lines

	if activity.MessageCount >= activityThreshold {
		fmt.Printf("Activity threshold reached!" + "\n") // Debug line
		if time.Since(activity.LastSpawn) > spawnCooldown {
			fmt.Printf("Spawning enemy..." + "\n") // Debug line
			enemy := config.GetRandomEnemy()
			if enemy != nil {
				activity.LastSpawn = time.Now()
				enemyEmbed := config.GetEnemyEmbed(enemy)
				session.ChannelMessageSendEmbed(channelID, enemyEmbed)
			}
		}
	}
}

package roleplay

import (
	"time"

	"github.com/Nota30/Kiko/lib/cache"
	database "github.com/Nota30/Kiko/lib/db"
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

func Actions (s *discordgo.Session, m *discordgo.MessageCreate, action string) {
	logrus.Info(action)
	cache.Client.SetEx(cache.Ctx, "hello", "bye", time.Hour)
	database.Db.Raw("INSERT INTO ROLEPLAY_COUNT (discord_id, type, send, received) VALUES (?, ?, ?, ?)", m.Author.ID, "hug", 1, 2)
}

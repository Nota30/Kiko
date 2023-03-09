package roleplay

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Nota30/Kiko/config"
	"github.com/Nota30/Kiko/lib/cache"
	database "github.com/Nota30/Kiko/lib/db"
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm/clause"
)


func generateEmbed(action string, header string, footer string, author discordgo.User, mention discordgo.User) (*discordgo.MessageEmbed) {
	resp, err := http.Get("https://api.waifu.pics/sfw/" + action)
	if err != nil {
   		logrus.Println(err)
		return nil
	}
	defer resp.Body.Close()
	var result map[string]string
	json.NewDecoder(resp.Body).Decode(&result)

	sent, recieved := count(action, author, mention)

	title1 := strings.Replace(header, "{0}", author.Mention(), -1)
	title2 := strings.Replace(title1, "{1}", mention.Mention(), -1)

	footer1 := strings.Replace(footer, "{0}", author.Username, -1)
	footer2 := strings.Replace(footer1, "{1}", strconv.Itoa(*recieved), -1)
	footer3 := strings.Replace(footer2, "{2}", mention.Username, -1)
	footer4 := strings.Replace(footer3, "{3}", strconv.Itoa(*sent), -1)

	embed := &discordgo.MessageEmbed{
		Color: config.Color.Default,
		Description: title2,
		Footer: &discordgo.MessageEmbedFooter{
			Text: footer4,
		},
		Image: &discordgo.MessageEmbedImage{
			URL:  result["url"],
		},
	}

	return embed
}

func count(action string, author discordgo.User, mention discordgo.User) (*int, *int) {
	var result []database.RoleplayCount
	database.Db.Raw("SELECT * FROM ROLEPLAY_COUNTS WHERE (discord_id = ? AND type = ?) OR (discord_id = ? AND type = ?)", author.ID, action, mention.ID, action).Scan(&result)
	u1 := findRPData(author.ID, action, result)

	if author.ID == mention.ID {
		return &u1.Sent, &u1.Received
	}

	u2 := findRPData(mention.ID, action, result)

	u1.Sent ++
	u2.Received ++

	database.Db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "discord_id",}, {Name: "type"}},
		DoUpdates: clause.AssignmentColumns([]string{"sent", "received"}),
	}).Create(&u1)

	database.Db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "discord_id",}, {Name: "type"}},
		DoUpdates: clause.AssignmentColumns([]string{"sent", "received"}),
	}).Create(&u2)

	return &u1.Sent, &u2.Received
}

func findRPData(discordId string, action string, list []database.RoleplayCount) database.RoleplayCount {
    for _, b := range list {
        if b.DiscordId == discordId {
            return b
        }
    }

	return database.RoleplayCount{
		DiscordId: discordId,
		Type: action,
		Sent: 0,
		Received: 0,
	}
}

func inCD(discordId string, guildId string, action string) time.Duration {
	req := cache.Client.TTL(cache.Ctx, "rpcd:" + discordId + ":" + guildId + ":" + action)
	rm, err := req.Result()

	if rm == time.Nanosecond*-2 {
		cache.Client.SetEx(cache.Ctx, "rpcd:" + discordId + ":" + guildId + ":" + action, 0, time.Minute*2)
		return 0
	}

	if err != nil {
		return 0
	}

	return rm
}
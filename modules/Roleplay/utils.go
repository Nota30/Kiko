package roleplay

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	database "github.com/Nota30/Kiko/lib/db"
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

type rp struct {
  Id   int64
  Discord_id string
  Type string
  Send int
  Received int
}

func generateEmbed (action string, header string, footer string, author discordgo.User, mention discordgo.User) (*discordgo.MessageEmbed) {
	resp, err := http.Get("https://api.waifu.pics/sfw/" + action)
	if err != nil {
   		logrus.Println(err)
		return nil
	}
	defer resp.Body.Close()
	var result map[string]string
	json.NewDecoder(resp.Body).Decode(&result)

	sent, recieved := getCount(action, author, mention)

	title1 := strings.Replace(header, "{0}", author.Mention(), -1)
	title2 := strings.Replace(title1, "{1}", mention.Mention(), -1)

	footer1 := strings.Replace(footer, "{0}", author.Username, -1)
	footer2 := strings.Replace(footer1, "{1}", strconv.Itoa(recieved), -1)
	footer3 := strings.Replace(footer2, "{2}", mention.Username, -1)
	footer4 := strings.Replace(footer3, "{3}", strconv.Itoa(sent), -1)

	embed := &discordgo.MessageEmbed{
		Color: 0x38d0fa,
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

func getCount (action string, author discordgo.User, mention discordgo.User) (int, int) {
	var result []rp
	database.Db.Raw("SELECT * FROM ROLEPLAY_COUNT WHERE (discord_id = ? AND type = ?) OR (discord_id = ? AND type = ?)", author.ID, action, mention.ID, action).Scan(&result)

	adata := findData(author.ID, result)
	bdata := findData(mention.ID, result)

	if adata == nil {
		adata = &rp{
			Id: 0,
			Discord_id: author.ID,
			Type: action,
			Send: 0,
			Received: 0,
		}

		database.Db.Exec("INSERT INTO ROLEPLAY_COUNT (discord_id, type, send, received) VALUES (?, ?, ?, ?)", author.ID, action, 1, 0)
	} else {
		database.Db.Exec("UPDATE ROLEPLAY_COUNT SET send = ? WHERE discord_id = ? AND type = ?", adata.Send + 1, mention.ID, action)
	}

	if bdata == nil {
		bdata = &rp{
			Id: 0,
			Discord_id: mention.ID,
			Type: action,
			Send: 0,
			Received: 0,
		}

		database.Db.Exec("INSERT INTO ROLEPLAY_COUNT (discord_id, type, send, received) VALUES (?, ?, ?, ?)", mention.ID, action, 0, 1)
	} else {
		database.Db.Exec("UPDATE ROLEPLAY_COUNT SET received = ? WHERE discord_id = ? AND type = ?", bdata.Received + 1, mention.ID, action)
	}

	return adata.Send + 1, bdata.Received + 1
}

func findData(a string, list []rp) *rp {
    for _, b := range list {
        if b.Discord_id == a {
            return &b
        }
    }

	return nil
}
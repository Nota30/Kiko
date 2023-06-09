package tower

import (
	"fmt"

	database "github.com/Nota30/Kiko/db"
	"github.com/bwmarrin/discordgo"
)

func Move(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options
	var data []database.Floor
	database.Db.Where("discord_id = ?", i.Member.User.ID).Order("floor desc").Find(&data)

	if len(data) == 0 {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
			Content: "You can't move floors yet! You haven't explored any floors yet. (Tip: You can explore floors with the `/explore` command)",
			},
		})
	} else if int(options[0].IntValue()) > data[0].Floor {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
			Content: "You can't move to that floor yet! You haven't reached that floor yet. You can ascend to higher floors by exploring and defeating enemies in the tower",
			},
		})
	} else {
		database.Db.Model(&database.Floor{}).Where("discord_id = ?", i.Member.User.ID).Where("active = ?", true).Update("active", false)
		curr, found := getData(data, int(options[0].IntValue()))
		if found {
			curr.Active = true
			database.Db.Save(&curr)
		} else {
			database.Db.Save(&database.Floor{
				DiscordId: i.Member.User.ID,
				Active: true,
				Floor: int(options[0].IntValue()),
			})
		}

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("You have moved to Floor **%d**", int(options[0].IntValue())),
			},
		})
	}
}

func getData(data []database.Floor, floor int) (user database.Floor, found bool) {
	for _, val := range data {
		if val.Floor == floor {
			return val, true
		}
	}

	return database.Floor{}, false
}
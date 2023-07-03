package tower

import (
	"fmt"

	"github.com/Nota30/Kiko/database"
	"github.com/Nota30/Kiko/models"
	"github.com/Nota30/Kiko/tools"
	"github.com/bwmarrin/discordgo"
)

func MoveCMD(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
	})
	options := i.ApplicationCommandData().Options
	embed := tools.NewEmbed(i.Member, "Move Floors")
	embed.Description = fmt.Sprintf("You have moved to Floor **%d**", int(options[0].IntValue()))
	floor, err := models.Floor.FindFloor(i.Member.User.ID, int(options[0].IntValue()))
	if err != nil {
		embed.Description = "You have not reached that floor yet! You can only move to floors you have explored or unlocked. (You can unlock floors by using the `/explore` command)"
	} else {
		if !floor.Active {
			err := database.Db.Model(&models.TFloor{}).Where("active = ?", true).Where("discord_id", i.Member.User.ID).Update("active", false).Error
			if err != nil {
				tools.SendError(s, i, "Couldn't update data. Please try again later.")
				return
			}
			floor.Active = true
			_, err = models.Floor.SaveFloor(floor)
			if err != nil {
				tools.SendError(s, i, "Couldn't update data. Please try again later.")
				return
			}
		}
	}

	s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
		Embeds: &[]*discordgo.MessageEmbed{&embed},
	})
}
package information

import (
	"fmt"

	"github.com/Nota30/Kiko/config"
	"github.com/Nota30/Kiko/models"
	"github.com/Nota30/Kiko/tools"
	"github.com/bwmarrin/discordgo"
)

func ProfileCMD(s *discordgo.Session, i *discordgo.InteractionCreate) {
	tools.AwaitInteraction(s, i)
	embed := tools.NewEmbed(i.Member, "Kiko Profile")

	user, err := models.User.FindUser(i.Member.User.ID)
	if err != nil {
		embed.Description  = "You aren't registered yet! You can register with the " + config.RegisterCommand + " command."
	} else {
		floor, _ := models.Floor.FindActiveFloor(i.Member.User.ID)
		general := fmt.Sprintf("```Class: %s \nSubClass: %s\nLevel: %d\nCoins: %d\nTower Floor: %d\nCurrent Floor Exploration: %v%%```", user.Class, user.Subclass, user.Level, user.Coins, floor.Floor, floor.Exploration)
		stats := fmt.Sprintf("```arm\nStrength: %d (+↑w-i%%) (tot. 10)\nAgility: %d (+↑w-i%%) (tot. 10)\nMana: %d (+↑w-i%%) (tot. 10)\nHealth: %d (+↑w-i%%) (tot. 10)\nDefence: %d (+↑w-i%%) (tot. 10)\nLuck: %d (+↑w-i%%) (tot. 10)```", user.Strength, user.Agility, user.Mana, user.Health, user.Defence, user.Luck)
		embed.Description = general + "\n**Stats**" + stats
	}

	_, err = s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
		Embeds: &[]*discordgo.MessageEmbed{&embed},
	})
	if err != nil {
		tools.SendError(s, i, "An error occured while responding.")
	}
}
package information

import (
	"fmt"

	"github.com/Nota30/Kiko/config"
	database "github.com/Nota30/Kiko/db"
	"github.com/bwmarrin/discordgo"
)

func Profile(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
	})
	var user database.User
	result := database.Db.Where("discord_id = ?", i.Member.User.ID).First(&user)
	var floor database.Floor
	database.Db.Where(&database.Floor{
		DiscordId: i.Member.User.ID,
		Active: true,
	}).First(floor)

	embed := discordgo.MessageEmbed{
		Color: config.Color.Default,
		Author: &discordgo.MessageEmbedAuthor{
			Name: "Kiko Profile",
			IconURL: i.Member.AvatarURL(""),
		},
		Description: "",
	}
	if result.Error != nil {
		embed.Description  = "You aren't registered yet! You can register with the " + config.RegisterCommand + " command."
	} else {
		general := fmt.Sprintf("```Class: %s \nSubClass: %s\nLevel: %d\nCoins: %d\nTower Floor: %d\nCurrent Floor Exploration: %v%%```", user.Class, user.Subclass, user.Level, user.Coins, floor.Floor, floor.Exploration)
		stats := fmt.Sprintf("```arm\nStrength: %d (+↑w-i%%) (tot. 10)\nAgility: %d (+↑w-i%%) (tot. 10)\nMana: %d (+↑w-i%%) (tot. 10)\nHealth: %d (+↑w-i%%) (tot. 10)\nDefence: %d (+↑w-i%%) (tot. 10)\nLuck: %d (+↑w-i%%) (tot. 10)```", user.Strength, user.Agility, user.Mana, user.Health, user.Defence, user.Luck)
		embed.Description = general + "\n**Stats**" + stats
	}

	s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
		Embeds: &[]*discordgo.MessageEmbed{&embed},
	})
}
package tower

import (
	"math/rand"
	"time"

	"github.com/Nota30/Kiko/config"
	"github.com/bwmarrin/discordgo"
)

func Explore(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
	})
	actions := make([]int, 0)
	actions = append(actions, 0, 0, 1, 1, 1, 1, 1, 2, 2, 2)
	rand.Seed(time.Now().Unix())
	action := actions[rand.Intn(len(actions))]

	embed := discordgo.MessageEmbed{
		Color: config.Color.Default,
		Author: &discordgo.MessageEmbedAuthor{
			Name: "Explore",
			IconURL: i.Member.AvatarURL(""),
		},
	}
	
	if action == 0 {
		embed.Description = "Dun dun dun!! You found nothing while exploring. Try exploring more! <:kiko_nervous:1007538837783859230>"
	} else if action == 1 {
		embed.Description = "Battle embed"
	} else if action == 2 {
		embed.Description = "Mineral embed"
	}
	s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
		Embeds: &[]*discordgo.MessageEmbed{&embed},
	})
}
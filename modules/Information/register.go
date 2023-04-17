package modules

import (
	"github.com/bwmarrin/discordgo"
)

func Register(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: options[0].Name + ": " + options[0].StringValue(),
		},
	})
}

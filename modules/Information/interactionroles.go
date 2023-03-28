package modules

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func InteractionRoles(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options
	rolesEmbed := discordgo.MessageEmbed{
		Fields: []*discordgo.MessageEmbedField{},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Role Selector by Kiko",
		},
	}
	content := ""

	switch options[0].Name {
	case "create":
		content = "Created role selection embed"
		rolesEmbed.Title = options[0].Options[1].StringValue()
		if len(options[0].Options) >= 3 {
			rolesEmbed.Description = options[0].Options[2].StringValue()
		}
		s.ChannelMessageSendEmbed(options[0].Options[0].ChannelValue(s).ID, &rolesEmbed)
	case "add":
		content = "Added role to role selector"
		message, err := s.ChannelMessage(i.ChannelID, options[0].Options[0].StringValue())
		if err != nil {
			log.Println(err)
		}
		message.Embeds[0].Fields = append(message.Embeds[0].Fields, &discordgo.MessageEmbedField{
			Name:  options[0].Options[2].StringValue() + " " + options[0].Options[1].RoleValue(s, i.GuildID).Name,
			Value: "",
		})
		s.ChannelMessageEditEmbed(i.ChannelID, message.ID, message.Embeds[0])
	case "remove":
		content = "Removed role from role selector"
		message, err := s.ChannelMessage(i.ChannelID, options[0].Options[0].StringValue())
		if err != nil {
			log.Println(err)
		}
		for n, v := range message.Embeds[0].Fields {
			if strings.Contains(v.Name, " "+options[0].Options[1].RoleValue(s, i.GuildID).Name) {
				if n == len(message.Embeds[0].Fields)-1 {
					message.Embeds[0].Fields = message.Embeds[0].Fields[0:n]
				} else {
					message.Embeds[0].Fields = append(message.Embeds[0].Fields[0:n], message.Embeds[0].Fields[n+1:]...)
				}
				break
			}
		}
		s.ChannelMessageEditEmbed(i.ChannelID, message.ID, message.Embeds[0])

	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: content,
		},
	})
}

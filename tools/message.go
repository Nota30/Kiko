package tools

import "github.com/bwmarrin/discordgo"

func Respond(s *discordgo.Session, i *discordgo.InteractionCreate, content string) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: content,
				},
			})
}

func NewSelectMenu(id string, placeholder string, options []discordgo.SelectMenuOption) discordgo.ActionsRow {
	return discordgo.ActionsRow{
		Components: []discordgo.MessageComponent{
			discordgo.SelectMenu{
				CustomID: id,
				Placeholder: placeholder,
				Options: options,
			},
		},
	}
}

func SendError(s *discordgo.Session, i *discordgo.InteractionCreate, description string) {
	embed := NewErrorEmbed(i.Member, description)
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: &discordgo.InteractionResponseData{
			Content: "",
			Embeds: []*discordgo.MessageEmbed{&embed},
			Components: nil,
		},
	})
}
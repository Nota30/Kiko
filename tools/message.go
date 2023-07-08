package tools

import "github.com/bwmarrin/discordgo"

func Respond(s *discordgo.Session, i *discordgo.InteractionCreate, content string) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: content,
				},
			})
			
	if err != nil {
		SendError(s, i, "An error occured while responding.")
	}
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

func AwaitInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
	})

	if err != nil {
		SendError(s, i, "An error occured while responding.")
	}
}

func SendError(s *discordgo.Session, i *discordgo.InteractionCreate, description string) {
	embed := NewErrorEmbed(i.Member, description)
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: &discordgo.InteractionResponseData{
			Content: "",
			Embeds: []*discordgo.MessageEmbed{&embed},
			Components: nil,
		},
	})

	if err != nil {
		_, err := s.ChannelMessageSend(i.ChannelID, description)
		if err != nil {
			return
		}
	}
}

func ResponseEdit(s *discordgo.Session, i *discordgo.InteractionCreate, messageData *MessageData) {
	var embeds *[]*discordgo.MessageEmbed
	if messageData.Embed != nil {
		embeds = &[]*discordgo.MessageEmbed{messageData.Embed}
	}

	if messageData.Embeds != nil {
		embeds = messageData.Embeds
	}

	_, err := s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
		Content: &messageData.Content,
		Embeds: embeds,
		Components: messageData.Components,
	})

	if err != nil {
		SendError(s, i, "An error occured while responding.")
	}
}

type MessageData struct {
	Content string
	Embed *discordgo.MessageEmbed
	Embeds *[]*discordgo.MessageEmbed
	Components *[]discordgo.MessageComponent
	Files []*discordgo.File
}
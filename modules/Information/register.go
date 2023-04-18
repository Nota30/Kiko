package modules

import (
	"time"

	"github.com/Nota30/Kiko/config"
	"github.com/bwmarrin/discordgo"
)

func Register(s *discordgo.Session, i *discordgo.InteractionCreate) {
	classes := *config.Classes
	choices := []discordgo.SelectMenuOption{}

	for tableName := range classes {
		choice := discordgo.SelectMenuOption{
			Label: tableName,
			Value: classes[tableName]["1"][0],
			Emoji: discordgo.ComponentEmoji{
				Name: classes[tableName]["1"][1],
			},
			Default: false,
		}
		choices = append(choices, choice)
	}

	embed := &discordgo.MessageEmbed{
		Color: config.Color.Default,
		Author: &discordgo.MessageEmbedAuthor{
			Name: "Kiko Registration",
			IconURL: i.Member.AvatarURL(""),
		},
		Description: 
		"Welcome to Kiko! Here you will register into the game and choose a class." +
		"It is also recommended to check out the `/tutorial` command before you proceed with the game.\n" +
		"**Please note that you cannot change classes further on in the game.**\n\n" + 
		"> There are 5 classes, which are `Warrior`, `Mage`, `Archer`, `Assassin` and `Martial Artist`\n\n" +
		"Each class has several sub-classes which you will evolve into as you further progress into the game. " + 
		"You will start out with the base class and recieve the default weapon for that class. " +
		"For more information on classes please check `/classes`.\n",
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Kiko is a QT",
			IconURL: s.State.User.AvatarURL(""),
		},
		Timestamp: time.Now().Format(time.RFC3339),
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embed},
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.SelectMenu{
							CustomID: "select_class",
							Placeholder: "Choose your class 👇",
							Options: choices,
						},
					},
				},
			},
		},
	})

	// Disable the selector after 3 minutes to prevent this event from being sent
	time.AfterFunc(3*time.Minute, func() {
		s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
			Components: &[]discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.SelectMenu{
							CustomID: "select_class",
							Placeholder: "Choose your class 👇",
							Disabled: true,
							Options: choices,
						},
					},
				},
			},
		})
	})
}

func RegisterSelector() {}
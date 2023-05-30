package modules

import (
	"time"

	"github.com/Nota30/Kiko/config"
	"github.com/Nota30/Kiko/config/store/weapons"
	"github.com/Nota30/Kiko/tools"
	"github.com/Nota30/Kiko/types"
	"github.com/bwmarrin/discordgo"
)

func Register(s *discordgo.Session, i *discordgo.InteractionCreate) {
	classes := config.Classes
	choices := []discordgo.SelectMenuOption{}

	for class, value := range classes {
		choice := discordgo.SelectMenuOption{
			Label: class,
			Value: class,
			Emoji: discordgo.ComponentEmoji{
				Name: value.Emote,
			},
			Default: false,
		}
		choices = append(choices, choice)
	}

	embed := &discordgo.MessageEmbed{
		Color: config.Color.Default,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    "Kiko Registration",
			IconURL: i.Member.AvatarURL(""),
		},
		Description: "Welcome to Kiko! Here you will register into the game and choose a class." +
			"It is also recommended to check out the `/tutorial` command before you proceed with the game.\n" +
			"**Please note that you cannot change classes further on in the game.**\n\n" +
			"> There are 5 classes, which are `Warrior`, `Mage`, `Archer`, `Assassin` and `Martial Artist`\n\n" +
			"Each class has several sub-classes which you will evolve into as you further progress into the game. " +
			"You will start out with the base class and recieve the default weapon for that class. " +
			"For more information on classes please check `/classes`.\n",
		Footer: &discordgo.MessageEmbedFooter{
			Text:    "Kiko is a QT",
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
							CustomID:    "select_class",
							Placeholder: "Choose your class 👇",
							Options:     choices,
						},
					},
				},
			},
		},
	})
}

func RegisterSelector(s *discordgo.Session, i *discordgo.InteractionCreate) {
	isExpired := time.Since(tools.ConvertToTime(i.Message.ID))

	choices := []discordgo.SelectMenuOption{}
	choice := discordgo.SelectMenuOption{
		Label: "Disabled",
		Value: "disabled",
		Default: false,
	}
	choices = append(choices, choice)

	if isExpired.Minutes() > 3 {
		return
	}

	var weapon types.Weapon
	class := i.MessageComponentData().Values[0]
	subclass := config.Classes[class].AdvanceClasses.One.Name

	switch class {
	case "Warrior":
		weapon = weapons.Common_Sword
	case "Mage":
		weapon = weapons.Common_Staff
	case "Martial Artist":
		weapon = weapons.Common_Gauntlet
	case "Assassin":
		weapon = weapons.Common_Dagger
	case "Archer":
		weapon = weapons.Common_Bow
	}

	embed := &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			Name:    "Class Selection",
			IconURL: i.Member.AvatarURL(""),
		},
		Color: config.Color.Default,
		Description: "The `" + i.MessageComponentData().Values[0] + "` class has been selected. " +
			"You have now been given the base sub-class of **`" + subclass + "`.**\n" +
			"Your randomly selected weapon is **`" + weapon.Name + "`**\n" +
			"I hope you have fun and enjoy your time on Kiko!",
		Timestamp: time.Now().Format(time.RFC3339),
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Kiko's Epic World",
		},
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embed},
		},
	})

	s.ChannelMessageEditComplex(&discordgo.MessageEdit{
		ID: i.Message.ID,
		Channel: i.Message.ChannelID,
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.SelectMenu{
						CustomID:    "select_class",
						Placeholder: "Choose your class 👇",
						Disabled:    true,
						Options:     choices,
					},
				},
			},
		},
	})
}
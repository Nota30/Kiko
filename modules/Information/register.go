package modules

import (
	"errors"
	"time"

	"github.com/Nota30/Kiko/config"
	"github.com/Nota30/Kiko/config/store/weapons"
	database "github.com/Nota30/Kiko/db"
	"github.com/Nota30/Kiko/tools"
	"github.com/Nota30/Kiko/types"
	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"
)

func Register(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var user database.User
	var embed discordgo.MessageEmbed
	var components []discordgo.MessageComponent
	classes := config.Classes
	choices := []discordgo.SelectMenuOption{}

	result := database.Db.Where("discord_id = ?", i.Member.User.ID).First(&user)
	
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		embed = discordgo.MessageEmbed{
			Color: config.Color.Default,
			Author: &discordgo.MessageEmbedAuthor{
				Name:    "Kiko Registration",
				IconURL: i.Member.AvatarURL(""),
			},
			Description: "Welcome to Kiko! Here you will register into the game and choose a class." +
				"It is also recommended to check out the `/tutorial` command before you proceed with the game.\n" +
				"**Please note that you cannot change classes further on in the game.**\n\n" +
				"There are 5 classes, which are:\n ```md\n# Warrior\n# Mage\n# Archer\n# Assassin\n# Martial Artist```\n" +
				"Each class has several sub-classes which you will evolve into as you further progress into the game. " +
				"You will start out with the base class and recieve the default weapon for that class. " +
				"For more information on classes please check `/classes`.\n",
		}

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

		components = []discordgo.MessageComponent{
			tools.NewSelectMenu("select_class", "Choose your class 👇", choices),
		}
	} else {
		embed = discordgo.MessageEmbed{
			Color: config.Color.Default,
			Author: &discordgo.MessageEmbedAuthor{
				Name:    "Kiko Registration",
				IconURL: i.Member.AvatarURL(""),
			},
			Description: "We love your enthusiasm for Kiko, however it seems you already have an active account!",
		}
		components = nil
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds:     []*discordgo.MessageEmbed{&embed},
			Components: components,
		},
	})
}

func RegisterSelector(s *discordgo.Session, i *discordgo.InteractionCreate) {
	isExpired := time.Since(tools.ConvertToTime(i.Message.ID))
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
	
	embed := discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			Name:    "Class Selection",
			IconURL: i.Member.AvatarURL(""),
		},
		Color: config.Color.Default,
		Description: "The `" + i.MessageComponentData().Values[0] + "` class has been selected. " +
			"You have now been given the base subclass: **```arm\n" + subclass + "```**\n" +
			"Your starter weapon is now: **```fix\n" + weapon.Name + "```**\n" +
			"I hope you have fun and enjoy your time on Kiko!",
	}

	user := database.User{
		DiscordId: i.Member.User.ID, 
		Class: i.MessageComponentData().Values[0], 
		Subclass: subclass,
		ClassAscension: "1",
	}

	item := database.Inventory{
		DiscordId: i.Member.User.ID,
		ItemName: weapon.Name,
		ItemType: "weapon",
		Active: true,
		Durability: weapon.Stats.Durability,
	}
	
	result := database.Db.Create(&user)
	if result.Error != nil {
		embed.Description = "There was a problem with registering your account, please try again later."
	} else {
		result = database.Db.Create(&item)
		if result.Error != nil {
			embed.Description = "There was a problem with registering your account, please try again later."
		}
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: &discordgo.InteractionResponseData{
			Embeds:     []*discordgo.MessageEmbed{&embed},
			Components: []discordgo.MessageComponent{},
		},
	})
}

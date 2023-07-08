package information

import (
	"errors"
	"time"

	"github.com/Nota30/Kiko/config"
	"github.com/Nota30/Kiko/config/store/weapons"
	"github.com/Nota30/Kiko/models"
	"github.com/Nota30/Kiko/tools"
	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"
)

func RegisterCMD(s *discordgo.Session, i *discordgo.InteractionCreate) {
	tools.AwaitInteraction(s, i)
	var components []discordgo.MessageComponent
	embed := tools.NewEmbed(i.Member, "Kiko Registration")
	classes := config.Classes
	choices := []discordgo.SelectMenuOption{}

	_, err := models.User.FindUser(i.Member.User.ID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		embed.Description = "Welcome to Kiko! Here you will register into the game and choose a class." +
		"It is also recommended to check out the `/tutorial` command before you proceed with the game.\n" +
		"**Please note that you cannot change classes further on in the game.**\n\n" +
		"There are 5 classes, which are:\n ```md\n# Warrior\n# Mage\n# Archer\n# Assassin\n# Martial Artist```\n" +
		"Each class has several sub-classes which you will evolve into as you further progress into the game. " +
		"You will start out with the base class and recieve the default weapon for that class. " +
		"For more information on classes please check `/classes`.\n"
		
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
		embed.Description = "We love your enthusiasm for Kiko, however it seems you already have an active account!"
		components = nil
	}

	tools.ResponseEdit(s, i, &tools.MessageData{
		Embed: &embed,
		Components: &components,
	})
}

// Handle the Selector component for the register command
func RegisterSelector(s *discordgo.Session, i *discordgo.InteractionCreate) {
	isExpired := time.Since(tools.ConvertToTime(i.Message.ID))
	if isExpired.Minutes() > 3 {
		return
	}

	var weapon weapons.TWeapon
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

	_, err := models.User.CreateUser(&models.TUser{
		DiscordId: i.Member.User.ID,
		Class: i.MessageComponentData().Values[0],
		Subclass: subclass,
		ClassAscension: "1",
	})
	if err != nil {
		tools.SendError(s, i, "There was a problem with registering your account, please try again later.")
		return
	} 

	_, err = models.Floor.CreateFloor(&models.TFloor{
		DiscordId: i.Member.User.ID,
		Active: true,
	})
	if err != nil {
		tools.SendError(s, i, "There was a problem with registering your account, please try again later.")
		return
	}

	_, err = models.Inventory.CreateInventory(&models.TInventory{
		DiscordId: i.Member.User.ID,
		ItemName: weapon.Name,
		ItemType: "weapon",
		Active: true,
		Durability: weapon.Stats.Durability,
	})
	if err != nil {
		tools.SendError(s, i, "There was a problem with registering your account, please try again later.")
		return
	}

	embed := tools.NewEmbed(i.Member, "Class Selection")
	embed.Description = "The `" + i.MessageComponentData().Values[0] + "` class has been selected. " +
	"You have now been given the base subclass: **```arm\n" + subclass + "```**\n" +
	"Your starter weapon is now: **```fix\n" + weapon.Name + "```**\n" +
	"I hope you have fun and enjoy your time on Kiko!"

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: &discordgo.InteractionResponseData{
			Embeds:     []*discordgo.MessageEmbed{&embed},
			Components: []discordgo.MessageComponent{},
		},
	})

	if err != nil {
		tools.SendError(s, i, "An error occured while responding.")
	}
}

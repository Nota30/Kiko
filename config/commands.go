package config

import (
	"github.com/bwmarrin/discordgo"
)

var dmPermission = false

var Commands = []*discordgo.ApplicationCommand{
	{
		Name:        "ping",
		Description: "Ping Pong.",
	},
	{
		Name:        "register",
		Description: "Register for the game",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "class",
				Description: "Pick your character class",
				Required:    true,
				Choices: func() []*discordgo.ApplicationCommandOptionChoice {
					choices := []*discordgo.ApplicationCommandOptionChoice{}
					SetupClasses()

					for tableName := range Data {
						for _, value := range Data[tableName]["1"] {
							choice := &discordgo.ApplicationCommandOptionChoice{
								Name:  tableName,
								Value: value,
							}
							choices = append(choices, choice)
						}
					}

					return choices
				}(),
			},
		},
	},
}

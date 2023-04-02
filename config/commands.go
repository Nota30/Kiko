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
		Name:        "interactionroles",
		Description: "Interaction Roles.",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Name:        "create",
				Description: "Create the role selector",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionChannel,
						Name:        "channel",
						Description: "The channel in which to add the role selector",
						Required:    true,
					},
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "title",
						Description: "The title of your role selector",
						Required:    true,
					},
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "description",
						Description: "The description of your role selector",
						Required:    false,
					},
				},
			},
			{
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Name:        "add",
				Description: "Add a role to the role selector",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "message-id",
						Description: "The message id of the role selector",
						Required:    true,
					},
					{
						Type:        discordgo.ApplicationCommandOptionRole,
						Name:        "role",
						Description: "The role to add",
						Required:    true,
					},
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "icon",
						Description: "The icon of the role",
						Required:    true,
					},
				},
			},
			{
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Name:        "remove",
				Description: "Remove a role from the role selector",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "message-id",
						Description: "The message id of the role selector",
						Required:    true,
					},
					{
						Type:        discordgo.ApplicationCommandOptionRole,
						Name:        "role",
						Description: "The role to remove",
						Required:    true,
					},
				},
			},
		},
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
					SetupClasses()
					choices := []*discordgo.ApplicationCommandOptionChoice{}

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

package config

import (
	"github.com/bwmarrin/discordgo"
)

var dmPermission = false
var towerMinFloor = 1.0

var Commands = []*discordgo.ApplicationCommand{
	{
		Name:        "ping",
		Description: "Ping Pong.",
		DMPermission: &dmPermission,
	},
	{
		Name:        "register",
		Description: "Register into Kiko",
		DMPermission: &dmPermission,
	},
	{
		Name:        "profile",
		Description: "View your profile on Kiko",
		DMPermission: &dmPermission,
	},
	{
		Name:        "explore",
		Description: "Explore the floor you are on",
		DMPermission: &dmPermission,
	},
	{
		Name:        "inventory",
		Description: "View the items you have",
		DMPermission: &dmPermission,
	},
	{
		Name:        "move",
		Description: "Move to a different floor in the tower",
		DMPermission: &dmPermission,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name: "floor",
				Type: discordgo.ApplicationCommandOptionInteger,
				MinValue: &towerMinFloor,
				Description: "The floor number",
				Required: true,
			},
		},
	},
	{
		Name:        "battle",
		Description: "Start a battle using a battle card",
		DMPermission: &dmPermission,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name: "id",
				Type: discordgo.ApplicationCommandOptionInteger,
				MinValue: &towerMinFloor,
				Description: "The id of the battle card (You can find the id using the /inventory command)",
				Required: true,
			},
		},
	},
}

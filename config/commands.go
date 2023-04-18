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
		Description: "Register into the best RPG Game!",
	},
}

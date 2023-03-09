package config

import "github.com/bwmarrin/discordgo"

var dmPermission = false

var Commands = []*discordgo.ApplicationCommand{
	{
		Name: "ping",
		Description: "Ping Pong.",
	},
}
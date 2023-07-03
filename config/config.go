package config

import (
	"github.com/Nota30/Kiko/config/store"
	"github.com/bwmarrin/discordgo"
)

type color struct {
	Default int
	Blue 	int
	Error   int
}

var ApplicationCommands []*discordgo.ApplicationCommand
var RegisterCommand string
var Color color = color{Default: 0xf62dcd, Error: 0xf32917, Blue: 0x17caf3}
var DevGuild = "983931249456455720"
var Classes = store.Classes
var Weapons = store.Weapons

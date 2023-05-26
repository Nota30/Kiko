package config

import "github.com/Nota30/Kiko/config/store"

type color struct {
	Default int
	Error   int
}

var Color color = color{Default: 0xf62dcd, Error: 0xf32917}
var DevGuild = "983931249456455720"
var Classes = store.Classes
var Weapons = store.Weapons

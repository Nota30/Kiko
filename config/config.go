package config

type color struct {
	Default int
	Error   int
}

var Color color = color{Default: 0xf62dcd, Error: 0xf32917}
var DevGuild = "972433897319178340"
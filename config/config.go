package config

type color struct {
	Default int
	Error   int
}

var Color color = color{Default: 0x38d0fa, Error: 0xf32917}
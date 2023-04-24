package config

func LoadConfigs() {
	SetupClasses()
	SetupWeapons()
}

type color struct {
	Default int
	Error   int
}

var Color color = color{Default: 0xf62dcd, Error: 0xf32917}
var DevGuild = "983931249456455720"
var Classes = &ClassData
var Swords = &SwordsData
var Staffs = &StaffsData
var Gauntlets = &GauntletsData
var Daggers = &DaggersData
var Bows = &BowsData

package classes

import "github.com/Nota30/Kiko/types"

// Martial Artist Class
var BaseMartialArtist = types.Class{
	Name: "Martial Artist",
	Emote: "🥋",
	AdvanceClasses: types.Ascension{
		One: Outlaw,
		Two: Barbarian,
		Three: Striker,
		Four: types.AdvAscension{
			First: Berserker, 
			Second: Brawler,
		},
		Five: types.AdvAscension{
			First: Warlord,
			Second: Grandmaster,
		},
	},
}

var Outlaw = types.AdvanceClass{
	Name: "Outlaw",
	Type: "Martial Artist",
	Ascension: 1,
	Stats: types.Stats{
		Strength: 4,
		Agility: 2,
		Mana: 2,
		Health: 0,
		Defence: 0,
		Luck: 0,
	},
}

var Barbarian = types.AdvanceClass{
	Name: "Barbarian",
	Type: "Martial Artist",
	Ascension: 2,
	Stats: types.Stats{
		Strength: 4,
		Agility: 2,
		Mana: 2,
		Health: 0,
		Defence: 0,
		Luck: 0,
	},
}

var Striker = types.AdvanceClass{
	Name: "Striker",
	Type: "Martial Artist",
	Ascension: 3,
	Stats: types.Stats{
		Strength: 4,
		Agility: 2,
		Mana: 2,
		Health: 0,
		Defence: 0,
		Luck: 0,
	},
}

var Berserker = types.AdvanceClass{
	Name: "Berserker",
	Type: "Martial Artist",
	Ascension: 4,
	Stats: types.Stats{
		Strength: 4,
		Agility: 2,
		Mana: 2,
		Health: 0,
		Defence: 0,
		Luck: 0,
	},
}

var Brawler = types.AdvanceClass{
	Name: "Brawler",
	Type: "Martial Artist",
	Ascension: 4,
	Stats: types.Stats{
		Strength: 4,
		Agility: 2,
		Mana: 2,
		Health: 0,
		Defence: 0,
		Luck: 0,
	},
}

var Warlord = types.AdvanceClass{
	Name: "Warlord",
	Type: "Martial Artist",
	Ascension: 5,
	Stats: types.Stats{
		Strength: 4,
		Agility: 2,
		Mana: 2,
		Health: 0,
		Defence: 0,
		Luck: 0,
	},
}

var Grandmaster = types.AdvanceClass{
	Name: "Grandmaster",
	Type: "Martial Artist",
	Ascension: 5,
	Stats: types.Stats{
		Strength: 4,
		Agility: 2,
		Mana: 2,
		Health: 0,
		Defence: 0,
		Luck: 0,
	},
}
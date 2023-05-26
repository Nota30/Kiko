package classes

import "github.com/Nota30/Kiko/types"

// Assasin Class
var BaseAssassin = types.Class{
	Name: "Assassin",
	Emote: "🗡️",
	AdvanceClasses: types.Ascension{
		One: Thief, 
		Two: Myrmidon, 
		Three: Hitman, 
		Four: types.AdvAscension{
			First: Rogue, 
			Second: Executioner,
		},
		Five: types.AdvAscension{
			First: Slayer,
			Second: Reaper,
		},
	},
}

var Thief = types.AdvanceClass{
	Name: "Thief",
	Type: "Assassin",
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

var Myrmidon = types.AdvanceClass{
	Name: "Myrmidon",
	Type: "Assassin",
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

var Hitman = types.AdvanceClass{
	Name: "Hitman",
	Type: "Assassin",
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

var Rogue = types.AdvanceClass{
	Name: "Rogue",
	Type: "Assassin",
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

var Executioner = types.AdvanceClass{
	Name: "Executioner",
	Type: "Assassin",
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

var Slayer = types.AdvanceClass{
	Name: "Slayer",
	Type: "Assassin",
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

var Reaper = types.AdvanceClass{
	Name: "Reaper",
	Type: "Assassin",
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
package classes

import "github.com/Nota30/Kiko/types"

// Mage Class
var BaseMage = types.Class{
	Name: "Mage",
	Emote: "🧙‍♂️",
	AdvanceClasses: types.Ascension{
		One: Wizard, 
		Two: Sorcerer, 
		Three: Summoner, 
		Four: types.AdvAscension{
			First: Shaman, 
			Second: Sage,
		},
		Five: types.AdvAscension{
			First: Necromancer,
			Second: Archmage,
		},
	},
}

var Wizard = types.AdvanceClass{
	Name: "Wizard",
	Type: "Mage",
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

var Sorcerer = types.AdvanceClass{
	Name: "Sorcerer",
	Type: "Mage",
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

var Summoner = types.AdvanceClass{
	Name: "Summoner",
	Type: "Mage",
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

var Shaman = types.AdvanceClass{
	Name: "Shaman",
	Type: "Mage",
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

var Sage = types.AdvanceClass{
	Name: "Sage",
	Type: "Mage",
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

var Necromancer = types.AdvanceClass{
	Name: "Necromancer",
	Type: "Mage",
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

var Archmage = types.AdvanceClass{
	Name: "Archmage",
	Type: "Mage",
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
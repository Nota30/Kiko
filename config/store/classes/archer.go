package classes

import "github.com/Nota30/Kiko/types"

// Archer Class
var BaseArcher = types.Class{
	Name: "Archer",
	Emote: "🏹",
	AdvanceClasses: types.Ascension{
		One: Fowler, 
		Two: Hunter, 
		Three: Arbalist, 
		Four: types.AdvAscension{
			First: Marksman, 
			Second: Ranger,
		},
		Five: types.AdvAscension{
			First: Sniper,
			Second: Sharpshooter,
		},
	},
}

var Fowler = types.AdvanceClass{
	Name: "Fowler",
	Type: "Archer",
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

var Hunter = types.AdvanceClass{
	Name: "Hunter",
	Type: "Archer",
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

var Arbalist = types.AdvanceClass{
	Name: "Arbalist",
	Type: "Archer",
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

var Marksman = types.AdvanceClass{
	Name: "Marksman",
	Type: "Archer",
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

var Ranger = types.AdvanceClass{
	Name: "Ranger",
	Type: "Archer",
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

var Sniper = types.AdvanceClass{
	Name: "Sniper",
	Type: "Archer",
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

var Sharpshooter = types.AdvanceClass{
	Name: "Sharpshooter",
	Type: "Archer",
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
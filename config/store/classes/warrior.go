package classes

import "github.com/Nota30/Kiko/types"

// Warrior Class
var BaseWarrior = types.Class{
	Name: "Warrior",
	Emote: "⚔️",
	AdvanceClasses: types.Ascension{
		One: Mercenary, 
		Two: Soldier, 
		Three: Knight, 
		Four: types.AdvAscension{
			First: DarkKnight, 
			Second: HolyKnight,
		},
		Five: types.AdvAscension{
			First: DeathKnight,
			Second: Paladin,
		},
	},
}

var Mercenary = types.AdvanceClass{
	Name: "Mercenary",
	Type: "Warrior",
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

var Soldier = types.AdvanceClass{
	Name: "Soldier",
	Type: "Warrior",
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

var Knight = types.AdvanceClass{
	Name: "Knight",
	Type: "Warrior",
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

var DarkKnight = types.AdvanceClass{
	Name: "Dark Knight",
	Type: "Warrior",
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

var HolyKnight = types.AdvanceClass{
	Name: "Holy Knight",
	Type: "Warrior",
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

var DeathKnight = types.AdvanceClass{
	Name: "DeathKnight",
	Type: "Warrior",
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

var Paladin = types.AdvanceClass{
	Name: "Paladin",
	Type: "Warrior",
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
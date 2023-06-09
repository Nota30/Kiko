package classes

// Warrior Class
var BaseWarrior = TClass{
	Name: "Warrior",
	Emote: "⚔️",
	AdvanceClasses: Ascension{
		One: Mercenary, 
		Two: Soldier, 
		Three: Knight, 
		Four: AdvAscension{
			First: DarkKnight, 
			Second: HolyKnight,
		},
		Five: AdvAscension{
			First: DeathKnight,
			Second: Paladin,
		},
	},
}

var Mercenary = AdvanceClass{
	Name: "Mercenary",
	Type: "Warrior",
	Ascension: 1,
	Stats: Stats{
		Strength: 4,
		Agility: 2,
		Mana: 2,
		Health: 0,
		Defence: 0,
		Luck: 0,
	},
}

var Soldier = AdvanceClass{
	Name: "Soldier",
	Type: "Warrior",
	Ascension: 2,
	Stats: Stats{
		Strength: 4,
		Agility: 2,
		Mana: 2,
		Health: 0,
		Defence: 0,
		Luck: 0,
	},
}

var Knight = AdvanceClass{
	Name: "Knight",
	Type: "Warrior",
	Ascension: 3,
	Stats: Stats{
		Strength: 4,
		Agility: 2,
		Mana: 2,
		Health: 0,
		Defence: 0,
		Luck: 0,
	},
}

var DarkKnight = AdvanceClass{
	Name: "Dark Knight",
	Type: "Warrior",
	Ascension: 4,
	Stats: Stats{
		Strength: 4,
		Agility: 2,
		Mana: 2,
		Health: 0,
		Defence: 0,
		Luck: 0,
	},
}

var HolyKnight = AdvanceClass{
	Name: "Holy Knight",
	Type: "Warrior",
	Ascension: 4,
	Stats: Stats{
		Strength: 4,
		Agility: 2,
		Mana: 2,
		Health: 0,
		Defence: 0,
		Luck: 0,
	},
}

var DeathKnight = AdvanceClass{
	Name: "DeathKnight",
	Type: "Warrior",
	Ascension: 5,
	Stats: Stats{
		Strength: 4,
		Agility: 2,
		Mana: 2,
		Health: 0,
		Defence: 0,
		Luck: 0,
	},
}

var Paladin = AdvanceClass{
	Name: "Paladin",
	Type: "Warrior",
	Ascension: 5,
	Stats: Stats{
		Strength: 4,
		Agility: 2,
		Mana: 2,
		Health: 0,
		Defence: 0,
		Luck: 0,
	},
}
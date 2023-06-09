package classes

// Mage Class
var BaseMage = TClass{
	Name: "Mage",
	Emote: "🧙‍♂️",
	AdvanceClasses: Ascension{
		One: Wizard, 
		Two: Sorcerer, 
		Three: Summoner, 
		Four: AdvAscension{
			First: Shaman, 
			Second: Sage,
		},
		Five: AdvAscension{
			First: Necromancer,
			Second: Archmage,
		},
	},
}

var Wizard = AdvanceClass{
	Name: "Wizard",
	Type: "Mage",
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

var Sorcerer = AdvanceClass{
	Name: "Sorcerer",
	Type: "Mage",
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

var Summoner = AdvanceClass{
	Name: "Summoner",
	Type: "Mage",
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

var Shaman = AdvanceClass{
	Name: "Shaman",
	Type: "Mage",
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

var Sage = AdvanceClass{
	Name: "Sage",
	Type: "Mage",
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

var Necromancer = AdvanceClass{
	Name: "Necromancer",
	Type: "Mage",
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

var Archmage = AdvanceClass{
	Name: "Archmage",
	Type: "Mage",
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
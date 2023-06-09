package classes

// Martial Artist Class
var BaseMartialArtist = TClass{
	Name: "Martial Artist",
	Emote: "🥋",
	AdvanceClasses: Ascension{
		One: Outlaw,
		Two: Barbarian,
		Three: Striker,
		Four: AdvAscension{
			First: Berserker, 
			Second: Brawler,
		},
		Five: AdvAscension{
			First: Warlord,
			Second: Grandmaster,
		},
	},
}

var Outlaw = AdvanceClass{
	Name: "Outlaw",
	Type: "Martial Artist",
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

var Barbarian = AdvanceClass{
	Name: "Barbarian",
	Type: "Martial Artist",
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

var Striker = AdvanceClass{
	Name: "Striker",
	Type: "Martial Artist",
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

var Berserker = AdvanceClass{
	Name: "Berserker",
	Type: "Martial Artist",
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

var Brawler = AdvanceClass{
	Name: "Brawler",
	Type: "Martial Artist",
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

var Warlord = AdvanceClass{
	Name: "Warlord",
	Type: "Martial Artist",
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

var Grandmaster = AdvanceClass{
	Name: "Grandmaster",
	Type: "Martial Artist",
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
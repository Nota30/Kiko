package classes

// Assasin Class
var BaseAssassin = TClass{
	Name: "Assassin",
	Emote: "🗡️",
	AdvanceClasses: Ascension{
		One: Thief, 
		Two: Myrmidon, 
		Three: Hitman, 
		Four: AdvAscension{
			First: Rogue, 
			Second: Executioner,
		},
		Five: AdvAscension{
			First: Slayer,
			Second: Reaper,
		},
	},
}

var Thief = AdvanceClass{
	Name: "Thief",
	Type: "Assassin",
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

var Myrmidon = AdvanceClass{
	Name: "Myrmidon",
	Type: "Assassin",
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

var Hitman = AdvanceClass{
	Name: "Hitman",
	Type: "Assassin",
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

var Rogue = AdvanceClass{
	Name: "Rogue",
	Type: "Assassin",
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

var Executioner = AdvanceClass{
	Name: "Executioner",
	Type: "Assassin",
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

var Slayer = AdvanceClass{
	Name: "Slayer",
	Type: "Assassin",
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

var Reaper = AdvanceClass{
	Name: "Reaper",
	Type: "Assassin",
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
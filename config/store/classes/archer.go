package classes

// Archer Class
var BaseArcher = TClass{
	Name: "Archer",
	Emote: "🏹",
	AdvanceClasses: Ascension{
		One: Fowler, 
		Two: Hunter, 
		Three: Arbalist, 
		Four: AdvAscension{
			First: Marksman, 
			Second: Ranger,
		},
		Five: AdvAscension{
			First: Sniper,
			Second: Sharpshooter,
		},
	},
}

var Fowler = AdvanceClass{
	Name: "Fowler",
	Type: "Archer",
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

var Hunter = AdvanceClass{
	Name: "Hunter",
	Type: "Archer",
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

var Arbalist = AdvanceClass{
	Name: "Arbalist",
	Type: "Archer",
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

var Marksman = AdvanceClass{
	Name: "Marksman",
	Type: "Archer",
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

var Ranger = AdvanceClass{
	Name: "Ranger",
	Type: "Archer",
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

var Sniper = AdvanceClass{
	Name: "Sniper",
	Type: "Archer",
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

var Sharpshooter = AdvanceClass{
	Name: "Sharpshooter",
	Type: "Archer",
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
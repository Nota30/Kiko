package weapons

import "github.com/Nota30/Kiko/types"

var BaseDaggers = types.BaseWeapon{
	Name: "Daggers",
	Weapons: []types.Weapon{Common_Dagger},
}

var Common_Dagger = types.Weapon{
	Name: "Common Dagger",
	Type: "Dagger",
	Rarity: 1,
	Description: "A common dagger mass produced by the local smithy",
	Price: 10,
	Droppable: true,
	InShop: true,
	Tradable: true,
	Stats: types.Stats{
		Strength: 1,
		Agility: 1,
		Mana: 0,
		Health: 1,
		Defence: 0,
		Luck: 0,
	},
}
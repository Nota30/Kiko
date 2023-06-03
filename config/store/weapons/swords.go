package weapons

import "github.com/Nota30/Kiko/types"

var BaseSwords = types.BaseWeapon{
	Name: "Sword",
	Weapons: []types.Weapon{Common_Sword},
}

var Common_Sword = types.Weapon{
	Name: "Common Sword",
	Type: "Sword",
	Rarity: 1,
	Description: "A common sword mass produced by the local smithy",
	Price: 10,
	Droppable: true,
	InShop: true,
	Tradable: true,
	Stats: types.ItemStats{
		Strength: 1,
		Agility: 1,
		Mana: 0,
		Health: 1,
		Defence: 0,
		Luck: 0,
		Durability: 5,
	},
}
package weapons

import "github.com/Nota30/Kiko/types"

var BaseBows = types.BaseWeapon{
	Name: "Bow",
	Weapons: []types.Weapon{Common_Bow},
}

var Common_Bow = types.Weapon{
	Name: "Common Bow",
	Type: "Bow",
	Rarity: 1,
	Description: "A common bow mass produced by the local smithy",
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
package weapons

import "github.com/Nota30/Kiko/types"

var BaseStaffs = types.BaseWeapon{
	Name: "Staff",
	Weapons: []types.Weapon{Common_Staff},
}

var Common_Staff = types.Weapon{
	Name: "Common Staff",
	Type: "Staff",
	Rarity: 1,
	Description: "A common staff mass produced by the local smithy",
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
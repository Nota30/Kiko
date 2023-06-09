package weapons

var BaseGauntlets = BaseWeapon{
	Name: "Gauntlet",
	Weapons: []TWeapon{Common_Gauntlet},
}

var Common_Gauntlet = TWeapon{
	Name: "Common Gauntlet",
	Type: "Gauntlet",
	Rarity: 1,
	Description: "A common gauntlet mass produced by the local smithy",
	Price: 10,
	Droppable: true,
	InShop: true,
	Tradable: true,
	Stats: ItemStats{
		Strength: 1,
		Agility: 1,
		Mana: 0,
		Health: 1,
		Defence: 0,
		Luck: 0,
		Durability: 5,
	},
}
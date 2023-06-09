package weapons

var BaseSwords = BaseWeapon{
	Name: "Sword",
	Weapons: []TWeapon{Common_Sword},
}

var Common_Sword = TWeapon{
	Name: "Common Sword",
	Type: "Sword",
	Rarity: 1,
	Description: "A common sword mass produced by the local smithy",
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
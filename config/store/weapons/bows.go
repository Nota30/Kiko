package weapons

var BaseBows = BaseWeapon{
	Name: "Bow",
	Weapons: []TWeapon{Common_Bow},
}

var Common_Bow = TWeapon{
	Name: "Common Bow",
	Type: "Bow",
	Rarity: 1,
	Description: "A common bow mass produced by the local smithy",
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
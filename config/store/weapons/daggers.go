package weapons

var BaseDaggers = BaseWeapon{
	Name: "Daggers",
	Weapons: []TWeapon{Common_Dagger},
}

var Common_Dagger = TWeapon{
	Name: "Common Dagger",
	Type: "Dagger",
	Rarity: 1,
	Description: "A common dagger mass produced by the local smithy",
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
package weapons

var BaseStaffs = BaseWeapon{
	Name: "Staff",
	Weapons: []TWeapon{Common_Staff},
}

var Common_Staff = TWeapon{
	Name: "Common Staff",
	Type: "Staff",
	Rarity: 1,
	Description: "A common staff mass produced by the local smithy",
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
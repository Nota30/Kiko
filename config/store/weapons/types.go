package weapons

type TWeapons struct {
	Swords BaseWeapon
	Bows BaseWeapon
	Staffs BaseWeapon
	Daggers BaseWeapon
	Gauntlets BaseWeapon
}

type BaseWeapon struct {
	Name string
	Weapons []TWeapon
}

type TWeapon struct {
	Name string
	Type string
	Rarity int
	Description string
	Price int
	Droppable bool
	InShop bool
	Tradable bool
	Stats ItemStats
}

type ItemStats struct {
	Strength float32
	Agility float32
	Mana float32
	Health float32
	Defence float32
	Luck float32
	Durability int
}
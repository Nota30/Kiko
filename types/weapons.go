package types

type Weapons struct {
	Swords BaseWeapon
	Bows BaseWeapon
	Staffs BaseWeapon
	Daggers BaseWeapon
	Gauntlets BaseWeapon
}

type BaseWeapon struct {
	Name string
	Weapons []Weapon
}

type Weapon struct {
	Name string
	Type string
	Rarity int
	Description string
	Price int
	Droppable bool
	InShop bool
	Tradable bool
	Stats Stats
}
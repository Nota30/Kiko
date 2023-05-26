package store

import (
	"github.com/Nota30/Kiko/config/store/classes"
	"github.com/Nota30/Kiko/config/store/weapons"
	"github.com/Nota30/Kiko/types"
)

var Classes = types.Classes{
	"Warrior": classes.BaseWarrior,
	"Archer": classes.BaseArcher,
	"Mage": classes.BaseMage,
	"Assassin": classes.BaseAssassin,
	"Martial Artist": classes.BaseMartialArtist,
}

var Weapons = types.Weapons{
	Swords: weapons.BaseSwords,
	Bows: weapons.BaseBows,
	Staffs: weapons.BaseStaffs,
	Daggers: weapons.BaseDaggers,
	Gauntlets: weapons.BaseGauntlets,
}
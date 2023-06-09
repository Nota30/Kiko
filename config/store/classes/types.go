package classes

type TClasses map[string] TClass

type TClass struct {
	Name string
	Emote string
	AdvanceClasses Ascension
}

type Ascension struct {
	One AdvanceClass
	Two AdvanceClass
	Three AdvanceClass
	Four AdvAscension
	Five AdvAscension
}

type AdvAscension struct {
	First AdvanceClass
	Second AdvanceClass
}

type AdvanceClass struct {
	Name string
	Type string
	Ascension int
	Stats Stats
}

type Stats struct {
	Strength int
	Agility int
	Mana int
	Health int
	Defence int
	Luck int
}
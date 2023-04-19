package config

import (
	"os"

	"github.com/pelletier/go-toml"
)

type Item struct {
	Name        string
	Price       int
	Description string
	Droppable   bool
	InShop      bool
	Tradable    bool
	Stats       struct {
		Attack  int
		Speed   int
		Defence int
		Health  int
		Stamina int
		IQ      int
		Luck    int
	}
}

var SwordsData map[string]map[string]Item

func SetupWeapons() {

	tomlData, err := os.ReadFile("../../config/store/weapons/swords.toml")
	if err != nil {
		panic(err)
	}

	// Parse TOML into a map[string]map[string][]string
	err = toml.Unmarshal(tomlData, &SwordsData)
	if err != nil {
		panic(err)
	}
}

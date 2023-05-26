package config

import (
	"os"
	"path/filepath"

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

var SwordsData, StaffsData, GauntletsData, DaggersData, BowsData map[string]map[string]Item

func SetupWeapons() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	tomlData, err := os.ReadFile(filepath.Join(cwd, "/config/store/weapons/swords.toml"))
	if err != nil {
		panic(err)
	}

	// Parse TOML into a map[string]map[string][]string
	err = toml.Unmarshal(tomlData, &SwordsData)
	if err != nil {
		panic(err)
	}

	tomlData, err = os.ReadFile(filepath.Join(cwd, "/config/store/weapons/staffs.toml"))
	if err != nil {
		panic(err)
	}

	// Parse TOML into a map[string]map[string][]string
	err = toml.Unmarshal(tomlData, &StaffsData)
	if err != nil {
		panic(err)
	}

	tomlData, err = os.ReadFile(filepath.Join(cwd, "/config/store/weapons/gauntlets.toml"))
	if err != nil {
		panic(err)
	}

	// Parse TOML into a map[string]map[string][]string
	err = toml.Unmarshal(tomlData, &GauntletsData)
	if err != nil {
		panic(err)
	}

	tomlData, err = os.ReadFile(filepath.Join(cwd, "/config/store/weapons/daggers.toml"))
	if err != nil {
		panic(err)
	}

	// Parse TOML into a map[string]map[string][]string
	err = toml.Unmarshal(tomlData, &DaggersData)
	if err != nil {
		panic(err)
	}

	tomlData, err = os.ReadFile(filepath.Join(cwd, "/config/store/weapons/bows.toml"))
	if err != nil {
		panic(err)
	}

	// Parse TOML into a map[string]map[string][]string
	err = toml.Unmarshal(tomlData, &BowsData)
	if err != nil {
		panic(err)
	}

}

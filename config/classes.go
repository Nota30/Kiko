package config

import (
	"os"

	"github.com/pelletier/go-toml"
)

var ClassData map[string]map[string][]string

func SetupClasses() {
	tomlData, err := os.ReadFile("../../config/store/classes.toml")
	if err != nil {
		panic(err)
	}

	// Parse TOML into a map[string]map[string][]string
	err = toml.Unmarshal(tomlData, &ClassData)
	if err != nil {
		panic(err)
	}
}

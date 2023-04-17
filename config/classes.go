package config

import (
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml"
)

var Data map[string]map[string][]string

func SetupClasses() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	tomlData, err := os.ReadFile(filepath.Join(cwd, "../../config/store/classes.toml"))
	if err != nil {
		panic(err)
	}

	// Parse TOML into a map[string]map[string][]string
	err = toml.Unmarshal(tomlData, &Data)
	if err != nil {
		panic(err)
	}
}

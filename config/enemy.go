package config

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
)

const (
	min = 1
	max = 10
)

type EnemyStruct struct {
	Name     string
	Strength int
	Mana     int
	Health   int
	Defence  int
}

var hollow = EnemyStruct{
	Name:     "Hollow",
	Strength: 100,
	Mana:     500,
	Health:   1500,
	Defence:  300,
}

var angel = EnemyStruct{
	Name:     "Angel",
	Strength: 300,
	Mana:     800,
	Health:   2000,
	Defence:  500,
}

var enemies = []EnemyStruct{hollow, angel}

func GetRandomEnemy() *EnemyStruct {
	rand.Seed(time.Now().UnixNano())
	spawnChance := (rand.Intn(max-min+1) + min)
	// 1 out of 10 chance for enemy to spawn
	if spawnChance == 5 {
		return &enemies[rand.Intn(len(enemies))]
	}
	return nil

}

func GetEnemyEmbed(enemy *EnemyStruct) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title:       "Enemy Spawned!",
		Color:       Color.Default,
		Description: "**" + enemy.Name + "**:\n\tStength:" + fmt.Sprint(enemy.Strength) + "\n\tHealth:" + fmt.Sprint(enemy.Health) + "\n\tMana:" + fmt.Sprint(enemy.Mana) + "\n\tDefence:" + fmt.Sprint(enemy.Defence),
	}
}

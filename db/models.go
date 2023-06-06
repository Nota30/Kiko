package database

import "time"

type User struct {
	ID        		int64
	DiscordId 		string
	Coins     		int
	Level			int
	Xp				int
	Kills			int
	Class     		string
	Subclass  		string
	ClassAscension 	string
	Strength 		int
	Agility 		int
	Mana 			int
	Health 			int
	Defence 		int
	Luck 			int
	CreatedAt 		time.Time
	UpdatedAt 		time.Time
}

type Inventory struct {
	ID 			int64
	DiscordId 	string
	ItemName 	string
	ItemType 	string
	Active 		bool
	Durability 	int
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
}

type Floor struct {
	ID 			int64
	DiscordId 	string
	Floor		int
	Exploration float32
	Active 		bool
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
}
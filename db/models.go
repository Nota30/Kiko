package database

import "time"

type User struct {
	ID        		int64
	DiscordId 		string
	Coins     		int
	Class     		string
	Subclass  		string
	ClassAscension 	string
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
package database

import "time"

type User struct {
	ID        int64
	DiscordId string
	Coins     int
	Class     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

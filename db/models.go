package database

import "time"

type User struct {
	ID 			int64
	DiscordId	string
	Class		string
	CreatedAt   time.Time
  	UpdatedAt   time.Time
}
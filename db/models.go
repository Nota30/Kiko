package database

import "time"

type RoleplayCount struct {
	ID 			int64
	DiscordId	string
	Type		string
	Sent		int `gorm:"default:0"`
  	Received	int `gorm:"default:0"`
	CreatedAt   time.Time
  	UpdatedAt   time.Time
}
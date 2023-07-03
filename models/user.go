package models

import (
	"time"

	"gorm.io/gorm"
)

type TUser struct {
	ID        		uint	`gorm:"type:bigserial;primaryKey"`
	DiscordId 		string	`gorm:"type:varchar;not null;unique;uniqueIndex"`
	Coins     		int		`gorm:"type:integer;default:0;not null"`
	Level			int		`gorm:"type:integer;default:1;not null"`
	Xp				int		`gorm:"type:integer;default:0;not null"`
	Kills			int		`gorm:"type:integer;default:0;not null"`
	Class     		string	`gorm:"type:varchar;not null"`
	Subclass  		string	`gorm:"type:varchar;not null"`
	ClassAscension 	string	`gorm:"type:varchar;default:'1';not null"`
	Strength 		int		`gorm:"type:integer;default:1;not null"`
	Agility 		int		`gorm:"type:integer;default:1;not null"`
	Mana 			int		`gorm:"type:integer;default:1;not null"`
	Health 			int		`gorm:"type:integer;default:1;not null"`
	Defence 		int		`gorm:"type:integer;default:1;not null"`
	Luck 			int		`gorm:"type:integer;default:1;not null"`
	CreatedAt 		time.Time
	UpdatedAt 		time.Time
}

type TUsers []TUser

type UserModel struct {
	db *gorm.DB
}

func NewUserModel(db *gorm.DB) *UserModel {
	return &UserModel{db: db}
}

func (m *UserModel) CreateUser(user *TUser) (result *TUser, err error) {
	err = m.db.Create(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (m *UserModel) SaveUser(user *TUser) (result *TUser, err error) {
	err = m.db.Save(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (m *UserModel) FindAllUsers() (result *TUsers, err error) {
	var users TUsers

	err = m.db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return &users, nil
}

func (m *UserModel) FindUser(discordId string) (result *TUser, err error) {
	var user TUser

	err = m.db.Where(&TUser{DiscordId: discordId}).Take(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
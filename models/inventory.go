package models

import (
	"time"

	"gorm.io/gorm"
)

type TInventory struct {
	ID        		uint	`gorm:"type:bigserial;primaryKey"`
	DiscordId 		string	`gorm:"type:varchar;not null"`
	ItemName     	string	`gorm:"type:varchar;not null"`
	ItemType		string	`gorm:"type:varchar;not null"`
	Durability		int		`gorm:"type:integer;not null"`
	Active			bool	`gorm:"type:boolean;default:false;not null"`
	CreatedAt 		time.Time
	UpdatedAt 		time.Time
}

type TInventories []TInventory

type InventoryModel struct {
	db *gorm.DB
}

func NewInventoryModel(db *gorm.DB) *InventoryModel {
	return &InventoryModel{db: db}
}

func (m *InventoryModel) CreateInventory(inv *TInventory) (result *TInventory, err error) {
	err = m.db.Create(inv).Error
	if err != nil {
		return nil, err
	}

	return inv, nil
}

func (m *InventoryModel) SaveInventory(inv *TInventory) (result *TInventory, err error) {
	err = m.db.Save(inv).Error
	if err != nil {
		return nil, err
	}

	return inv, nil
}

func (m *InventoryModel) FindAllInventories() (result *TInventories, err error) {
	var invs TInventories

	err = m.db.Find(&invs).Error
	if err != nil {
		return nil, err
	}

	return &invs, nil
}

func (m *InventoryModel) FindInventory(discordId string) (result *TInventory, err error) {
	var inv TInventory

	err = m.db.Where(&TInventory{DiscordId: discordId}).Take(&inv).Error
	if err != nil {
		return nil, err
	}

	return &inv, nil
}
package models

import (
	"time"

	"gorm.io/gorm"
)

type TFloor struct {
	ID        		uint	`gorm:"type:bigserial;primaryKey"`
	DiscordId 		string	`gorm:"type:varchar;not null;uniqueIndex:idx_floors_number"`
	Floor     		int		`gorm:"type:integer;default:1;not null;uniqueIndex:idx_floors_number"`
	Exploration		float32	`gorm:"type:real;default:0;not null"`
	Active			bool	`gorm:"type:boolean;default:false;not null"`
	CreatedAt 		time.Time
	UpdatedAt 		time.Time
}

type TFloors []TFloor

type FloorModel struct {
	db *gorm.DB
}

func NewFloorModel(db *gorm.DB) *FloorModel {
	return &FloorModel{db: db}
}

func (m *FloorModel) CreateFloor(floor *TFloor) (result *TFloor, err error) {
	err = m.db.Create(floor).Error
	if err != nil {
		return nil, err
	}

	return floor, nil
}

func (m *FloorModel) SaveFloor(floor *TFloor) (result *TFloor, err error) {
	err = m.db.Save(floor).Error
	if err != nil {
		return nil, err
	}

	return floor, nil
}

func (m *FloorModel) FindAllFloors() (result *TFloors, err error) {
	var floors TFloors

	err = m.db.Find(&floors).Error
	if err != nil {
		return nil, err
	}

	return &floors, nil
}

func (m *FloorModel) FindFloor(discordId string, floorNumber int) (result *TFloor, err error) {
	var floor TFloor

	err = m.db.Where(&TFloor{DiscordId: discordId, Floor: floorNumber}).Take(&floor).Error
	if err != nil {
		return nil, err
	}

	return &floor, nil
}

func (m *FloorModel) FindActiveFloor(discordId string) (result *TFloor, err error) {
	var floor TFloor

	err = m.db.Where(&TFloor{DiscordId: discordId, Active: true}).Take(&floor).Error
	if err != nil {
		return nil, err
	}

	return &floor, nil
}
package database

import (
	"github.com/Nota30/Kiko/models"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Db  *gorm.DB
	err error
)

func ConnDB(dbConURL string) {
	Db, err = gorm.Open(postgres.Open(dbConURL))
	if err != nil {
		logrus.Fatal(err)
	}

	psql, err := Db.DB()
	if err != nil {
		logrus.Fatal(err)
	}
	// pong the database
	err = psql.Ping()
	if err != nil {
		logrus.Fatal(err)
	}
	// migration
	err = Db.AutoMigrate(&models.TUser{}, &models.TFloor{}, &models.TInventory{})
	if err != nil {
		logrus.Fatal(err)
	}
	
	models.User = models.NewUserModel(Db)
	models.Floor = models.NewFloorModel(Db)
	models.Inventory = models.NewInventoryModel(Db)
}
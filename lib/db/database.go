package database

import (
	"github.com/Nota30/Kiko/tools"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Db *gorm.DB
	err error
)

func Connect() {
	dsn := tools.GetEnv("DB_URL")
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
    	panic("failed to connect database")
  	}

	psqlDB, err := Db.DB()
	if err != nil {
    	panic("failed to connect database")
  	}
	
	psqlDB.SetMaxIdleConns(5)
	psqlDB.SetMaxOpenConns(0)
}
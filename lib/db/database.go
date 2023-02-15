package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Db *gorm.DB
	err error
)

func Connect() {
	dsn := "host=localhost user=kiko password=1234 dbname=kiko port=5432"
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
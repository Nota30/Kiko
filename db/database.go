package database

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	Db  *gorm.DB
	err error
)

func Connect() {
	dsn := os.Getenv("DB_URL")
	env := os.Getenv("env")
	log := logger.Default.LogMode(logger.Silent)
	if env == "development" {
		log = logger.Default.LogMode(logger.Warn)
	}
	
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: log,
	})

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

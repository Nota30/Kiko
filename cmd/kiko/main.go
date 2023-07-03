package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Nota30/Kiko/bot"
	"github.com/Nota30/Kiko/cache"
	"github.com/Nota30/Kiko/database"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		logrus.Fatal("Error loading .env file")
	} else {
		logrus.Info("Loaded the environment variables!!")
	}

	cache.Connect()
	dbUrl := os.Getenv("DB_URL")
	if len(dbUrl) == 0 {
		logrus.Fatal("Database connection string was not found")
	}

	database.ConnDB(dbUrl)

	bot.DiscordConnect()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	bot.Dg.Close()
	defer cache.Client.Close()
}

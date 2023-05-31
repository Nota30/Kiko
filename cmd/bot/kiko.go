package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Nota30/Kiko/bot"
	"github.com/Nota30/Kiko/cache"
	database "github.com/Nota30/Kiko/db"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	err := godotenv.Load("../../.env")

	if err != nil {
		logrus.Fatalf("Error loading .env file")
	} else {
		logrus.Info("Loaded the environment variables!!")
	}

	cache.Connect()
	database.Connect()
	bot.DiscordConnect()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	bot.Dg.Close()
	defer cache.Client.Close()
}

package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Nota30/Kiko/lib/bot"
	"github.com/Nota30/Kiko/lib/cache"
	database "github.com/Nota30/Kiko/lib/db"
)

func main() {
	cache.Connect()
	database.Connect()
	bot.DiscordConnect()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	bot.Sharder.StopAll()
	defer cache.Client.Close()
}
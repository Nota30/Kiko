package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Nota30/Kiko/bot"
	"github.com/Nota30/Kiko/cache"
	database "github.com/Nota30/Kiko/db"
)

func main() {
	cache.Connect()
	database.Connect()
	bot.DiscordConnect()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	bot.Dg.Close()
	defer cache.Client.Close()
}

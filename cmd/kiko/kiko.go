package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Nota30/Kiko/lib/bot"
)

func main() {
	bot.DiscordConnect()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	bot.Sharder.StopAll()
}
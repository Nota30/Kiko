package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/Nota30/Kiko/bot"
	"github.com/Nota30/Kiko/cache"
	database "github.com/Nota30/Kiko/db"
	"github.com/Nota30/Kiko/lib/slash"
)

func main() {

	slashPtr := flag.Bool("s", false, "Flag to indicate whether to update the slash commands or not")

	flag.Parse()

	if *slashPtr {
		slash.UpdateSlashCMDS()
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

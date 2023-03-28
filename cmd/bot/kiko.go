package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/Nota30/Kiko/lib/bot"
	"github.com/Nota30/Kiko/lib/cache"
	database "github.com/Nota30/Kiko/lib/db"
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

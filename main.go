package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"tg_bot_eye/bot"
	"tg_bot_eye/config"
)

func abort(b *bot.Bot) {
	log.Println("Abort telegram bot")
	bot.SaveTasks(b.Tasks)
	os.Exit(1)

}

func main() {
	log.Println("Start telegram bot")

	fileConfig := flag.String("c", "config/config.json", "# config file")
	flag.Parse()

	conf := config.New(*fileConfig)

	tg_bot := bot.New(conf.Token)

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		abort(tg_bot)
	}()

	tg_bot.Start()
}

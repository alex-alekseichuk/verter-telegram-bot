package main

import (
	"github.com/alex-alekseichuk/verter-telegram-bot/internal/game/bangladesh"
	"github.com/alex-alekseichuk/verter-telegram-bot/internal/telegram"
	"log"
)

func main() {
	bot, err := telegram.NewTelegramBot()
	if err != nil {
		log.Fatal(err)
	}
	bangladesh.NewBotController().WithNewGame().WithBot(bot)
	bot.Run()
}

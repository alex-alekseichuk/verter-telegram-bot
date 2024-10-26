package main

import (
	guessNumber "github.com/alex-alekseichuk/verter-telegram-bot/internal/game/guess-number"
	"github.com/alex-alekseichuk/verter-telegram-bot/internal/telegram"
	"log"
)

func main() {
	bot, err := telegram.NewTelegramBot()
	if err != nil {
		log.Fatal(err)
	}

	guessNumber.NewBotController().WithBot(bot)
	bot.Run()
}

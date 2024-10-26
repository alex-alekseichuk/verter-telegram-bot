package main

import (
	"github.com/alex-alekseichuk/verter-telegram-bot/internal/game/bangladesh"
	"github.com/alex-alekseichuk/verter-telegram-bot/internal/telegram"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	bot, err := telegram.NewTelegramBot()
	if err != nil {
		log.Fatal(err)
	}
	bangladesh.NewBotController().WithNewGame().WithBot(bot)
	bot.Run()
}

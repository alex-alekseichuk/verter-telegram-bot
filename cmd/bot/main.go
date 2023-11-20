package main

import (
	"github.com/alex-alekseichuk/verter-telegram-bot/internal/pkg/bot"
	"github.com/alex-alekseichuk/verter-telegram-bot/internal/service"
	"github.com/alex-alekseichuk/verter-telegram-bot/internal/service/bangladesh"
)

func main() {
	var game *service.Game = bangladesh.NewService()
	games := []*service.Game{game}
	telbot, err := bot.NewTelegramBot(games)
	if err == nil {
		telbot.Run()
	}
}

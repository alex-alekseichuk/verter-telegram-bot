package service

import (
	"github.com/alex-alekseichuk/verter-telegram-bot/internal/pkg/bot"
)

type Game interface {
	Init(*bot.Bot)
}

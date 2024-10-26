package main

import (
	"github.com/alex-alekseichuk/verter-telegram-bot/internal/game/bangladesh"
)

func main() {
	host := bangladesh.NewCliHost()
	host.Run()
}

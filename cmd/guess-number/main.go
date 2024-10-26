package main

import (
	guessNumber "github.com/alex-alekseichuk/verter-telegram-bot/internal/game/guess-number"
)

func main() {
	host := guessNumber.NewCliHost().WithNewGame()
	host.Run()
}

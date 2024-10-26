package guess_number

import (
	bot "github.com/alex-alekseichuk/verter-telegram-bot/internal/bot"
)

// GameHost implements game/guess_number/Host
type GameHost struct {
	controller *BotController
	game       *Game
	chatId     bot.ChatId
}

func (host *GameHost) OnStarted(message string) {
	host.controller.sendMessage(host.chatId, message)
}

func (host *GameHost) OnRightGuess(message string) {
	host.controller.sendMessage(host.chatId, message)
}

func (host *GameHost) OnFailedGuess(message string) {
	host.controller.sendMessage(host.chatId, message)
}

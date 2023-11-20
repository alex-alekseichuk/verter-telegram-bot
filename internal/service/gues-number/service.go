package gues-number

import (
	"fmt"
	"github.com/alex-alekseichuk/verter-telegram-bot/internal/pkg/bot"
)

const (
	STARTED = iota
)

type Game struct {
	playerId int64
	mode     int
}

func NewGame(playerId int64) *Game {
	return &Game{playerId: playerId, mode: STARTED}
}

type Service struct {
	bot   *bot.Bot
	games map[int64]*Game
}

func NewService(bot *bot.Bot) *Service {
	return &Service{bot: bot, games: map[int64]*Game{}}
}

func (service *Service) startGame(playerId int64) {
	bot := *service.bot
	_, found := service.games[playerId]
	if found {
		bot.SendPrivate(playerId, "The game is in progress. Finish it first.")
		return
	}
	service.games[playerId] = NewGame(playerId)
	bot.SendPrivate(playerId, "Ok. Let's play the bangladesh game.\nGues a number?")
}

func (service *Service) guesNumber(playerId int64, num int) {
	bot := *service.bot
	_, found := service.games[playerId]
	if !found {
		bot.SendPrivate(playerId, "First, start the game.")
		return
	}
	service.games[playerId] = NewGame(playerId)
	myNum := num + 1
	bot.SendPrivate(playerId, fmt.Sprintf("Well, my number is %d. It's bigger than yours, so I won.", myNum))
}

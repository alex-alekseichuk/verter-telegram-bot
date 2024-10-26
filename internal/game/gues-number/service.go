package gues

//import (
//	"fmt"
//	"github.com/alex-alekseichuk/verter-telegram-bot/internal/bot"
//)
//
//const (
//	STARTED = iota
//)
//
//type Game struct {
//	playerId int64
//	mode     int
//}
//
//func NewGame(playerId int64) *Game {
//	return &Game{playerId: playerId, mode: STARTED}
//}
//
//type Service struct {
//	bot   bot.Bot
//	game map[int64]*Game
//}
//
//func NewService(bot bot.Bot) *Service {
//	return &Service{
//		bot:   bot,
//		game: map[int64]*Game{},
//	}
//}
//
//func (service *Service) startGame(playerId int64) {
//	_, found := service.game[playerId]
//	if found {
//		service.bot.SendPrivate(playerId, "The game is in progress. Finish it first.")
//		return
//	}
//	service.game[playerId] = NewGame(playerId)
//	service.bot.SendPrivate(playerId, "Ok. Let's play the bangladesh game.\nGues a number?")
//}
//
//func (service *Service) guesNumber(playerId int64, num int) {
//	_, found := service.game[playerId]
//	if !found {
//		service.bot.SendPrivate(playerId, "First, start the game.")
//		return
//	}
//	service.game[playerId] = NewGame(playerId)
//	myNum := num + 1
//	service.bot.SendPrivate(playerId, fmt.Sprintf("Well, my number is %d. It's bigger than yours, so I won.", myNum))
//}

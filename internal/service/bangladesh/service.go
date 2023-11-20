package bangladesh

import (
	"fmt"

	"github.com/alex-alekseichuk/verter-telegram-bot/internal/pkg/bot"
)

type Service struct {
	bot *bot.Bot
}

func NewService() *Service {
	return &Service{bot: nil}
}

func (service *Service) Init2(bot *bot.Bot) {
	service.bot = bot
}

func (service Service) Init(bot *bot.Bot) {
	service.bot = bot
}

func (service *Service) startGame(playerId int64) {
	bot := *service.bot
	bot.SendPrivate(playerId, "Ok. Let's play the bangladesh game.\nGues a number?")
}

func (service *Service) guesNumber(playerId int64, num int) {
	bot := *service.bot
	myNum := num + 1
	bot.SendPrivate(playerId, fmt.Sprintf("Well, my number is %d. It's bigger than yours, so I won.", myNum))
}

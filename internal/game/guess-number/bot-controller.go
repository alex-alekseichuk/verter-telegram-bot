package guess_number

import (
	bot "github.com/alex-alekseichuk/verter-telegram-bot/internal/bot"
	"strconv"
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

// BotController implements bot/Controller
type BotController struct {
	bot       bot.Bot
	gameHosts map[bot.ChatId]*GameHost
}

func NewBotController() *BotController {
	return &BotController{
		gameHosts: map[bot.ChatId]*GameHost{},
	}
}

func (ctl *BotController) WithBot(bot bot.Bot) *BotController {
	ctl.bot = bot
	bot.SetController(ctl)
	return ctl
}

func (ctl *BotController) OnCommand(cmd bot.Command) {
	switch cmd.Cmd {
	case "start":
		host, exists := ctl.gameHosts[cmd.ChatId]
		if !exists {
			host = ctl.createGame(cmd.ChatId)
		}
		host.game.Start()
	case "help":
		ctl.sendMessage(cmd.ChatId, "/start - start new game")
	}
}

func (ctl *BotController) OnMessage(message bot.InputMessage) {
	num, err := strconv.Atoi(message.Text)
	if err != nil {
		return
	}
	host, exists := ctl.gameHosts[message.ChatId]
	if !exists {
		return
	}
	host.game.GuessNumber(num)
}

func (ctl *BotController) createGame(chatId bot.ChatId) *GameHost {
	host := &GameHost{
		controller: ctl,
		chatId:     chatId,
	}
	host.game = NewGame(host)
	ctl.gameHosts[chatId] = host
	return host
}

func (ctl *BotController) sendMessage(chatId bot.ChatId, message string) {
	ctl.bot.Send(bot.OutputMessage{
		ChatId: chatId,
		Text:   message,
	})
}

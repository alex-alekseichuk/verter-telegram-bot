package bangladesh

import (
	bot "github.com/alex-alekseichuk/verter-telegram-bot/internal/bot"
	"strconv"
)

// BotController implements game/bangladesh/Host
// BotController implements bot/Controller
type BotController struct {
	game   *Game
	bot    bot.Bot
	chatId int64
}

func NewBotController() *BotController {
	return &BotController{}
}

func (c *BotController) WithNewGame() *BotController {
	c.game = NewGame(c)
	return c
}

func (c *BotController) WithBot(bot bot.Bot) *BotController {
	c.bot = bot
	bot.SetController(c)
	return c
}

func (c *BotController) OnCommand(cmd bot.Command) {
	switch cmd.Cmd {
	case "start":
		c.chatId = cmd.ChatId
		c.game.Start()
	}
}

func (c *BotController) OnMessage(message bot.InputMessage) {
	num, err := strconv.Atoi(message.Text)
	if err == nil {
		c.game.GuessNumber(num)
	}
}

func (c *BotController) OnStarted(message string) {
	c.bot.Send(bot.OutputMessage{
		ChatId: c.chatId,
		Text:   message,
	})
}

func (c *BotController) OnFinished(message string) {
	c.bot.Send(bot.OutputMessage{
		ChatId: c.chatId,
		Text:   message,
	})
}

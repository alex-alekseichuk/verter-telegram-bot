package telegram

import (
	"log"

	"github.com/alex-alekseichuk/verter-telegram-bot/internal/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramBot struct {
	bot   *tgbotapi.BotAPI
	games []*service.Game
}

func NewTelegramBot(games []*service.Game) (*TelegramBot, error) {
	bot, err := tgbotapi.NewBotAPI("1287198594:AAETZvG1fcywdHQAm4lAtsb234iUQlhYOwc")
	if err != nil {
		log.Panic(err)
		return nil, err
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	telbot := &TelegramBot{bot, games}

	for _, game := range games {
		(*game).Init(telbot)
	}

	return telbot, nil
}

func (telbot *TelegramBot) Run() {
	bot := telbot.bot
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		msg := update.Message
		log.Printf("[%s] %s", msg.From.UserName, msg.Text)

		cmd := msg.Command()
		switch cmd {
		case "help":
			newMsg := tgbotapi.NewMessage(msg.Chat.ID, "Help:")
			bot.Send(newMsg)
		default:
			newMsg := tgbotapi.NewMessage(msg.Chat.ID, msg.Text)
			newMsg.ReplyToMessageID = msg.MessageID
			bot.Send(newMsg)
		}

	}
}

func (telbot *TelegramBot) SendPrivate(playerId int64, msg string) {
	// TODO
}

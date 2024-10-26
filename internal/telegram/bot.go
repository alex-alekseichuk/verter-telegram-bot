package telegram

import (
	bot "github.com/alex-alekseichuk/verter-telegram-bot/internal/bot"
	tgBotApi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type Bot struct {
	botApi     *tgBotApi.BotAPI
	controller bot.Controller
}

func NewTelegramBot() (*Bot, error) {
	botApi, err := tgBotApi.NewBotAPI("1287198594:AAETZvG1fcywdHQAm4lAtsb234iUQlhYOwc")
	if err != nil {
		log.Panic(err)
		return nil, err
	}

	botApi.Debug = true

	log.Printf("Authorized on account %s", botApi.Self.UserName)

	telBot := &Bot{
		botApi: botApi,
	}

	return telBot, nil
}

func (tgBot *Bot) SetController(c bot.Controller) {
	tgBot.controller = c
}

func (tgBot *Bot) Run() {
	botApi := tgBot.botApi
	u := tgBotApi.UpdateConfig{
		Timeout: 60,
	}

	updates := botApi.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		msg := update.Message

		if msg.IsCommand() {
			cmd := msg.Command()
			command := bot.Command{
				Id:       msg.MessageID,
				ChatId:   bot.ChatId(msg.Chat.ID),
				PlayerId: msg.From.ID,
				Name:     msg.From.UserName,
				Cmd:      cmd,
			}
			tgBot.controller.OnCommand(command)
			continue
		}

		message := bot.InputMessage{
			Id:       msg.MessageID,
			ChatId:   bot.ChatId(msg.Chat.ID),
			PlayerId: msg.From.ID,
			Name:     msg.From.UserName,
			Text:     msg.Text,
		}
		tgBot.controller.OnMessage(message)
	}
}

func (tgBot *Bot) Send(message bot.OutputMessage) {
	tgMsg := tgBotApi.NewMessage(int64(message.ChatId), message.Text)
	if message.ReplyToMessageID != 0 {
		tgMsg.ReplyToMessageID = message.ReplyToMessageID
	}
	tgBot.botApi.Send(tgMsg)
}

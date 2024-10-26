package telegram

//type TelegramBot struct {
//	bot   *tgbotapi.BotAPI
//	game []game.Game
//}
//
//func NewTelegramBot(game []game.Game) (*TelegramBot, error) {
//	bot, err := tgbotapi.NewBotAPI("1287198594:AAETZvG1fcywdHQAm4lAtsb234iUQlhYOwc")
//	if err != nil {
//		log.Panic(err)
//		return nil, err
//	}
//
//	bot.Debug = true
//
//	log.Printf("Authorized on account %s", bot.Self.UserName)
//
//	telbot := &TelegramBot{bot, game}
//
//	for _, game := range game {
//		(game).Init(telbot)
//	}
//
//	return telbot, nil
//}
//
//func (telbot *TelegramBot) Run() {
//	bot := telbot.bot
//	u := tgbotapi.UpdateConfig{
//		Timeout: 60,
//	}
//
//	updates := bot.GetUpdatesChan(u)
//
//	for update := range updates {
//		if update.Message == nil {
//			continue
//		}
//		msg := update.Message
//		log.Printf("[%d][%s] %s", msg.From.ID, msg.From.UserName, msg.Text)
//
//		cmd := msg.Command()
//		log.Printf("command: %s", cmd)
//
//		switch cmd {
//		case "help":
//			newMsg := tgbotapi.NewMessage(msg.Chat.ID, "Help:")
//			bot.Send(newMsg)
//		default:
//			newMsg := tgbotapi.NewMessage(msg.Chat.ID, msg.Text)
//			newMsg.ReplyToMessageID = msg.MessageID
//			bot.Send(newMsg)
//		}
//
//	}
//}
//
//func (telbot *TelegramBot) SendPrivate(playerId int64, msg string) {
//	// TODO
//}

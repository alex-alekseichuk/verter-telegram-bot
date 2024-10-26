package bot

type Message struct {
	Id               int64
	PlayerId         int64
	Name             int64
	Text             string
	ReplyToMessageID int64
}

type Command struct {
	PlayerId int64
	Text     string
}

type Bot interface {
	SendPrivate(message Message)
}

type Controller interface {
	OnMessage(message Message)
	OnCommand(command Command)
}

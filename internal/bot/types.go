package bot

type ChatId int64

type InputMessage struct {
	Id       int
	ChatId   ChatId
	PlayerId int64
	Name     string
	Text     string
}

type OutputMessage struct {
	ChatId           ChatId
	Text             string
	ReplyToMessageID int
}

type Command struct {
	Id       int
	ChatId   ChatId
	PlayerId int64
	Name     string
	Cmd      string
}

type Controller interface {
	OnMessage(message InputMessage)
	OnCommand(command Command)
}

type Bot interface {
	Send(message OutputMessage)
	SetController(c Controller)
}

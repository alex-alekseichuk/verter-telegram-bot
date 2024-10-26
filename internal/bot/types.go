package bot

type InputMessage struct {
	Id       int
	ChatId   int64
	PlayerId int64
	Name     string
	Text     string
}

type OutputMessage struct {
	ChatId           int64
	Text             string
	ReplyToMessageID int
}

type Command struct {
	Id       int
	ChatId   int64
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

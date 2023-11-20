package bot

type Bot interface {
	SendPrivate(playerId int64, msg string)
}

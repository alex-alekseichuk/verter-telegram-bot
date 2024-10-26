package bangladesh

type Controller interface {
	Start(playerId int)
	Guess(guess int)
}

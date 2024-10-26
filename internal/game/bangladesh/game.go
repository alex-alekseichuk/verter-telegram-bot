package bangladesh

import (
	"fmt"
)

type Host interface {
	OnStarted(message string)
	OnFinished(message string)
}

type Status int

const (
	Started Status = iota + 1
	Finished
)

type Game struct {
	host   Host
	status Status
}

func NewGame(host Host) *Game {
	return &Game{
		host:   host,
		status: Finished,
	}
}

func (game *Game) SetHost(host Host) {
	game.host = host
}

func (game *Game) Start() {
	if game.status != Finished {
		return
	}
	game.status = Started
	game.host.OnStarted("Ok. Let's play the bangladesh game.\nGuess a number?")
}

func (game *Game) GuessNumber(num int) {
	if game.status != Started {
		return
	}
	game.status = Finished
	game.host.OnFinished(fmt.Sprintf("Well, my number is %d. It's greater than yours, so I won.", num+1))
}

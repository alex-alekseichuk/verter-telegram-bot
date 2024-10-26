package gues_number

import (
	"fmt"
	"math/rand"
)

type Host interface {
	OnStarted(message string)
	OnFinished(message string)
}

type Status int

const (
	Init Status = iota
	Started
	Finished
)

type Game struct {
	host         Host
	status       Status
	hiddenNumber int
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
	if game.status == Init || game.status == Finished {
		game.status = Started
		game.hiddenNumber = rand.Intn(10) + 1
		game.host.OnStarted("Guess a number from 1 to 10?")
	}
}

func (game *Game) GuessNumber(num int) {
	if game.status == Started {
		if num == game.hiddenNumber {
			game.status = Finished
			game.host.OnFinished(fmt.Sprintf("Exactly! It's %d", game.hiddenNumber))
			return
		}
		if game.hiddenNumber > num {
			game.host.OnFinished(fmt.Sprintf("No, it's greater than %d", game.hiddenNumber))
			return
		}
		if game.hiddenNumber < num {
			game.host.OnFinished(fmt.Sprintf("No, it's greater than %d", game.hiddenNumber))
			return
		}
	}
}

package guess_number

import "fmt"

// CliHost implements game/guess-number/Host
type CliHost struct {
	game *Game
}

func NewCliHost() *CliHost {
	return &CliHost{}
}

func (c *CliHost) Run() {
	c.game.Start()
}

func (c *CliHost) WithNewGame() *CliHost {
	c.game = NewGame(c)
	return c
}

func (c *CliHost) OnStarted(message string) {
	c.guessNumber(message)
}

func (c *CliHost) OnRightGuess(message string) {
	fmt.Println(message)
}

func (c *CliHost) OnFailedGuess(message string) {
	c.guessNumber(message)
}

func (c *CliHost) guessNumber(message string) {
	fmt.Println(message)
	c.game.GuessNumber(inputNumber())
}

func inputNumber() int {
	var num int
	if _, err := fmt.Scan(&num); err != nil {
		num = 0
	}
	return num
}

package bangladesh

import "fmt"

// CliHost implements game/bangladesh/Host
type CliHost struct {
	game *Game
}

func NewCliHost() *CliHost {
	return &CliHost{}
}

func (c *CliHost) Run() {
	c.createGame()
	c.game.Start()
}

func (c *CliHost) createGame() (g *Game) {
	c.game = NewGame(c)
	return c.game
}

func (c *CliHost) OnStarted(message string) {
	fmt.Println(message)
	var num int
	if _, err := fmt.Scan(&num); err != nil {
		num = 0
	}
	c.game.GuessNumber(num)
}

func (c *CliHost) OnFinished(message string) {
	fmt.Println(message)
}

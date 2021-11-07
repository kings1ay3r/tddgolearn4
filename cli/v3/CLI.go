package poker

import (
	"bufio"
	"io"
	"strings"

	poker "github.com/quii/learn-go-with-tests/command-line/v3"
)

type CLI struct {
	playerStore poker.PlayerStore
	in          *bufio.Scanner
}

func (cli *CLI) PlayPoker() {
	userInput := cli.readLine()
	cli.playerStore.RecordWin(extractWinner(userInput))
}
func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}
func NewCLI(store poker.PlayerStore, in io.Reader) *CLI {
	return &CLI{
		playerStore: store,
		in:          bufio.NewScanner(in),
	}
}
func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}

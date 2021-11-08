package poker

import (
	"bufio"
	"io"
	"strings"
)

type CLI struct {
	playerStore PlayerStore
	in          *bufio.Scanner
}

func (cli *CLI) PlayPoker() {
	userInput := cli.readLine()
	cli.playerStore.RecordWin(extractWinner(userInput))

}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}

func NewCLI(p PlayerStore, s io.Reader) *CLI {
	return &CLI{p, bufio.NewScanner(s)}
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}

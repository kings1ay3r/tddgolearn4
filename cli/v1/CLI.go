package poker

import (
	"io"

	poker "github.com/quii/learn-go-with-tests/command-line/v3"
)

type CLI struct {
	playerStore poker.PlayerStore
	in          io.Reader
}

func (cli *CLI) PlayPoker() {
	cli.playerStore.RecordWin("Cleo")
}

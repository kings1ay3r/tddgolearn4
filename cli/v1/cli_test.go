package poker

import (
	"strings"
	"testing"

	poker "github.com/quii/learn-go-with-tests/command-line/v3"
)

func TestCLI(t *testing.T) {
	playerStore := &poker.StubPlayerStore{}
	in := strings.NewReader("Chris wins\n")
	cli := &CLI{playerStore, in}
	cli.PlayPoker()
	if len(playerStore.WinCalls) != 1 {
		t.Fatal("expected a win call but didn't get any")
	}
	got := playerStore.WinCalls[0]
	want := "Chris"
	if got != want {
		t.Errorf("didn't record correct winner, got %q, want %q", got, want)
	}
}

package poker_test

import (
	"application/poker"
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {

	t.Run("record André win from user input", func(t *testing.T) {
		in := strings.NewReader("André wins\n")
		playerStore := &poker.StubPlayerStore{}

		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "André")
	})

	t.Run("record Eduardo win from user input", func(t *testing.T) {
		in := strings.NewReader("Eduardo wins\n")
		playerStore := &poker.StubPlayerStore{}

		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Eduardo")
	})

}

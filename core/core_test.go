package core_test

import (
	point "point/core"
	"testing"
)

var p = point.Players{
	PlayerName: "Nick",
	Money:      2000,
	Bet:        200,
	Card:       "",
	Points:     0,
}

var b = point.Bot{
	DefName: "Diler",
	Money:   2000,
	Bet:     200,
	Card:    "",
	Points:  0,
}

var w = point.Winner{
	WinnerName: "",
}

func TestGame(t *testing.T) {
	_, _, winner := point.StartGame(&p, &b, &w)
	t.Run("return winner a Nick, balance + bet", func(t *testing.T) {
		p.Points = 0
		b.Points = 50
		got := winner

		want := "Nick"

		assertWinner(t, got, want)
	})
}

func assertEqualPlayer(t testing.TB, got, want point.Players) {
	t.Helper()

	if got != want {
		t.Errorf("got %v\nwant %v", got, want)
	}
}

func assertEqualBot(t testing.TB, got, want point.Bot) {
	t.Helper()

	if got != want {
		t.Errorf("got %v\nwant %v", got, want)
	}
}

func assertWinner(t testing.TB, got point.Winner, want string) {
	if got.WinnerName != want {
		t.Errorf("got %v\nwant %v", got, want)
	}
}

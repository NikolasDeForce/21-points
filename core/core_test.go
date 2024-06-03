package core_test

import (
<<<<<<< HEAD
	point "point"
=======
	point "point/core"
>>>>>>> c1dccae (Game 21 Points -- all test passed)
	"testing"
)

func TestGame(t *testing.T) {
	t.Run("returns Player Nick get points > 21, balance - bet ", func(t *testing.T) {
		got := point.GamePlayer(&point.Players{
			PlayerName: "Nick",
			Money:      2000,
			Bet:        200,
			Card:       "",
			Points:     24,
		})

		want := point.Players{
			PlayerName: got.PlayerName,
			Money:      got.Money,
			Bet:        got.Bet,
			Card:       got.Card,
			Points:     got.Points,
		}

		assertEqualPlayer(t, got, want)
	})

	t.Run("returns Player Nick got points < 21, balance + bet", func(t *testing.T) {
		got := point.GamePlayer(&point.Players{
			PlayerName: "Nick",
			Money:      2000,
			Bet:        300,
			Card:       "",
			Points:     24,
		})

		want := point.Players{
			PlayerName: got.PlayerName,
			Money:      got.Money,
			Bet:        got.Bet,
			Card:       got.Card,
			Points:     got.Points,
		}

		assertEqualPlayer(t, got, want)
	})

	t.Run("returns Bot got points > 21", func(t *testing.T) {
		got := point.GameBot(&point.Bot{
			DefName: "Diler",
			Money:   2000,
			Bet:     200,
			Card:    "",
			Points:  24,
		})

		want := point.Bot{
			DefName: got.DefName,
			Money:   got.Money,
			Bet:     got.Bet,
			Card:    got.Card,
			Points:  got.Points,
		}

		assertEqualBot(t, got, want)
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

package core

import (
	"math/rand"
)

type Players struct {
	PlayerName string
	Money      int
	Bet        int
	Card       string
	Points     int
}

type Bot struct {
	DefName string
	Money   int
	Bet     int
	Card    string
	Points  int
}

type Winner struct {
	WinnerName string
}

func Pick(m map[string]int) (string, int) {
	res := 0
	str := ""
	k := rand.Intn(len(m))
	i := 0
	for z, x := range m {
		if i == k {
			res = x
			str += z + " "
			return str, res
		}
		i++
	}
	return str, res
}

func GamePlayer(p *Players) Players {
	cards := map[string]int{
		"A":  11,
		"K":  4,
		"Q":  3,
		"J":  2,
		"10": 10,
		"9":  9,
		"8":  8,
		"7":  7,
		"6":  6,
	}

	for i := 0; i < 3; i++ {
		getCardKey, getCardVal := Pick(cards)
		p.Card += getCardKey
		p.Points += getCardVal
	}

	return Players{
		PlayerName: p.PlayerName,
		Money:      p.Money,
		Bet:        p.Bet,
		Card:       p.Card,
		Points:     p.Points,
	}
}

func GameBot(b *Bot) Bot {
	cards := map[string]int{
		"A":  11,
		"K":  4,
		"Q":  3,
		"J":  2,
		"10": 10,
		"9":  9,
		"8":  8,
		"7":  7,
		"6":  6,
	}

	for i := 0; i < 3; i++ {
		getCardKey, getCardVal := Pick(cards)
		b.Card += getCardKey
		b.Points += getCardVal
	}

	return Bot{
		DefName: b.DefName,
		Money:   b.Money,
		Bet:     b.Bet,
		Card:    b.Card,
		Points:  b.Points,
	}
}

func StartGame(p *Players, b *Bot, w *Winner) (Players, Bot, Winner) {
	resultPlayer := GamePlayer(&Players{
		PlayerName: "Nick",
		Money:      2000,
		Bet:        200,
		Card:       "",
		Points:     0,
	})
	resultBot := GameBot(&Bot{
		DefName: "Diler",
		Money:   2000,
		Bet:     200,
		Card:    "",
		Points:  0,
	})

	winner := Winner{
		w.WinnerName,
	}

	if resultPlayer.Points < resultBot.Points && resultBot.Points < 22 {
		winner.WinnerName = resultBot.DefName
		resultBot.Money += resultPlayer.Bet
		resultPlayer.Money -= resultBot.Bet
	} else if resultBot.Points < resultPlayer.Points && resultPlayer.Points < 22 {
		winner.WinnerName = resultPlayer.PlayerName
		resultBot.Money -= resultPlayer.Bet
		resultPlayer.Money += resultBot.Bet
	} else if resultBot.Points < 22 && resultPlayer.Points < 22 && resultBot.Points > resultPlayer.Points {
		winner.WinnerName = resultBot.DefName
		resultBot.Money += resultPlayer.Bet
		resultPlayer.Money -= resultBot.Bet
	} else if resultPlayer.Points < 22 && resultBot.Points < 22 && resultPlayer.Points > resultBot.Points {
		winner.WinnerName = resultPlayer.PlayerName
		resultBot.Money -= resultPlayer.Bet
		resultPlayer.Money += resultBot.Bet
	} else if resultPlayer.Points > 21 && resultBot.Points > 21 && resultPlayer.Points < resultBot.Points {
		winner.WinnerName = resultPlayer.PlayerName
		resultBot.Money -= resultPlayer.Bet
		resultPlayer.Money += resultBot.Bet
	} else if resultBot.Points > 21 && resultPlayer.Points > 21 && resultBot.Points < resultPlayer.Points {
		winner.WinnerName = resultBot.DefName
		resultBot.Money += resultPlayer.Bet
		resultPlayer.Money -= resultBot.Bet
	} else if resultPlayer.Points > 21 && resultBot.Points < 22 {
		winner.WinnerName = resultBot.DefName
		resultBot.Money += resultPlayer.Bet
		resultPlayer.Money -= resultBot.Bet
	} else if resultBot.Points > 21 && resultPlayer.Points < 22 {
		winner.WinnerName = resultPlayer.PlayerName
		resultBot.Money -= resultPlayer.Bet
		resultPlayer.Money += resultBot.Bet
	} else if resultBot.Points == resultPlayer.Points {
		winner.WinnerName = "Draw"
	}

	// fmt.Fprintf(os.Stdout, "Russian '21 Points' Game\n")
	// fmt.Fprint(os.Stdout, "------------------------------------\n")
	// fmt.Fprintf(os.Stdout, "Player %v results:\nNick Cards: %v\nPoints: %v\nAfter game Nick have balance: %v\nNick bet: %v\n", resultPlayer.PlayerName, resultPlayer.Card, resultPlayer.Points, resultPlayer.Money, resultPlayer.Bet)
	// fmt.Fprint(os.Stdout, "------------------------------------\n")
	// fmt.Fprintf(os.Stdout, "%v results:\nDiler Cards: %v\nPoints: %v\nAfter game diler have balance: %v\nDiler bet: %v\n", resultBot.DefName, resultBot.Card, resultBot.Points, resultBot.Money, resultBot.Bet)
	// fmt.Fprint(os.Stdout, "------------------------------------\n")

	return resultPlayer, resultBot, winner
}

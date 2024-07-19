package core

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

var cards = map[string]int{
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

func GamePlayer(p *Players) Players {
	for i := 0; i < 2; i++ {
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

func StartGame(p *Players, b *Bot) (Players, Bot) {
	resultPlayer := GamePlayer(&Players{
		PlayerName: p.PlayerName,
		Money:      2000,
		Bet:        p.Bet,
		Card:       "",
		Points:     0,
	})
	resultBot := GameBot(&Bot{
		DefName: "Diler",
		Money:   2000,
		Bet:     b.Bet,
		Card:    "",
		Points:  0,
	})

	return Players{PlayerName: resultPlayer.PlayerName, Money: resultPlayer.Money, Bet: resultPlayer.Bet, Card: resultPlayer.Card, Points: resultPlayer.Points},
		Bot{DefName: resultBot.DefName, Money: resultBot.Money, Bet: resultBot.Bet, Card: resultBot.Card, Points: resultBot.Points}
}

func ResultGame(p *Players, b *Bot, w *Winner) (Players, Bot, Winner) {
	if p.Points < b.Points && b.Points < 22 {
		w.WinnerName = b.DefName
		b.Money += p.Bet
		p.Money -= b.Bet
	} else if b.Points < p.Points && p.Points < 22 {
		w.WinnerName = p.PlayerName
		b.Money -= p.Bet
		p.Money += b.Bet
	} else if b.Points < 22 && p.Points < 22 && b.Points > p.Points {
		w.WinnerName = b.DefName
		b.Money += p.Bet
		p.Money -= b.Bet
	} else if p.Points < 22 && b.Points < 22 && p.Points > b.Points {
		w.WinnerName = p.PlayerName
		b.Money -= p.Bet
		p.Money += b.Bet
	} else if p.Points > 21 && b.Points > 21 && p.Points < b.Points {
		w.WinnerName = p.PlayerName
		b.Money -= p.Bet
		p.Money += b.Bet
	} else if b.Points > 21 && p.Points > 21 && b.Points < p.Points {
		w.WinnerName = b.DefName
		b.Money += p.Bet
		p.Money -= b.Bet
	} else if p.Points > 21 && b.Points < 22 {
		w.WinnerName = b.DefName
		b.Money += p.Bet
		p.Money -= b.Bet
	} else if b.Points > 21 && p.Points < 22 {
		w.WinnerName = p.PlayerName
		b.Money -= p.Bet
		p.Money += b.Bet
	} else if b.Points == p.Points {
		w.WinnerName = "Draw"
	}

	return Players{PlayerName: p.PlayerName, Money: p.Money, Bet: p.Bet, Card: p.Card, Points: p.Points},
		Bot{DefName: b.DefName, Money: b.Money, Bet: b.Bet, Card: b.Card, Points: b.Points},
		Winner{WinnerName: w.WinnerName}
}

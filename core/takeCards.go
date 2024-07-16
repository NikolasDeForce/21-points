package core

import "math/rand"

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

func TakeOne(p *Players) Players {
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

	getCardKey, getCardVal := Pick(cards)
	p.Card += getCardKey
	p.Points += getCardVal

	return Players{
		PlayerName: p.PlayerName,
		Money:      p.Money,
		Bet:        p.Bet,
		Card:       p.Card,
		Points:     p.Points,
	}
}

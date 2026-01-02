package combination

import (
	. "pokergame/poker/game/types"
)

func IsFlush(muck []Card) (bool, HandValue) {
	maxSuit := make(map[Suit]uint8)
	var maxTotal uint8
	var suit Suit
	for _, i := range muck {
		maxSuit[i.Suit]++
		if maxSuit[i.Suit] > maxTotal {
			maxTotal = maxSuit[i.Suit]
			suit = i.Suit
		}
	}
	if maxTotal < 5 {
		return false, HandValue{}
	}

	flush := make([]Card, 0)
	for _, i := range muck {
		if i.Suit == suit {
			flush = append(flush, i)
		}
	}

	flush = flush[len(flush)-5:]

	var arrFlush [5]Rank
	for i := 0; i < len(flush); i++ {
		arrFlush[i] = flush[i].Rank
	}

	return true, HandValue{Rank: Flush, Cards: arrFlush}
}

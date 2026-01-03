package combination

import (
	"fmt"
	. "pokergame/poker/game/types"
)

func IsTwoPair(muck []Card) (bool, HandValue) {
	maxRank := make(map[Rank]uint8)

	var firstPair, secondPair, kicker Rank

	// считаем карты
	for _, c := range muck {
		maxRank[c.Rank]++
	}

	// ищем две старшие пары
	for r, c := range maxRank {
		if c >= 2 {
			if r > firstPair {
				secondPair = firstPair
				firstPair = r
			} else if r > secondPair {
				secondPair = r
			}
		}
	}

	if firstPair == 0 || secondPair == 0 {
		return false, HandValue{}
	}

	// ищем киккер
	for r := range maxRank {
		if r != firstPair && r != secondPair && r > kicker {
			kicker = r
		}
	}
	var maxTwoPair = [5]Rank{firstPair, firstPair, secondPair, secondPair, kicker}
	fmt.Printf("Max combination: maxTwoPair\nCards: %v\n", maxTwoPair)
	return true, HandValue{
		Rank:  TwoPair,
		Cards: maxTwoPair,
	}
}

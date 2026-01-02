package combination

import (
	. "pokergame/poker/game/types"
)

func IsHighCard(muck []Card) (bool, HandValue) {
	var maxHightCard = [5]Rank{}
	n := len(muck)

	// Берём 5 старших карт с конца массива
	for i := 0; i < 5 && i < n; i++ {
		maxHightCard[i] = muck[n-1-i].Rank
	}

	return true, HandValue{Rank: HighCard, Cards: maxHightCard}
}

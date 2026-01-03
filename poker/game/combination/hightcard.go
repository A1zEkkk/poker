package combination

import (
	"fmt"
	. "pokergame/poker/game/types"
)

func IsHighCard(muck []Card) (bool, HandValue) {
	var maxHightCard = [5]Rank{}
	n := len(muck)

	// Берём 5 старших карт с конца массива
	for i := 0; i < 5 && i < n; i++ {
		maxHightCard[i] = muck[n-1-i].Rank
	}

	fmt.Printf("Max combination: HighCard\nCards: %v\n", maxHightCard)
	return true, HandValue{Rank: HighCard, Cards: maxHightCard}
}

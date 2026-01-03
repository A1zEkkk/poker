package combination

import (
	"fmt"
	. "pokergame/poker/game/types"
)

func IsSet(muck []Card) (bool, HandValue) {
	maxRank := make(map[Rank]uint8)
	var setRank, maxKicker1, maxKicker2 Rank
	findSet := false

	// 1. Находим ранг сета
	for _, c := range muck {
		maxRank[c.Rank]++

		if maxRank[c.Rank] == 3 && c.Rank > setRank {
			setRank = c.Rank
			findSet = true
		}
	}

	if !findSet {
		return false, HandValue{}
	}

	// 2. Теперь корректно ищем два старших киккера
	for _, c := range muck {
		if c.Rank == setRank {
			continue
		}

		if c.Rank > maxKicker1 {
			maxKicker2 = maxKicker1
			maxKicker1 = c.Rank
		} else if c.Rank > maxKicker2 {
			maxKicker2 = c.Rank
		}
	}

	var arrSet = [5]Rank{
		setRank, setRank, setRank,
		maxKicker1, maxKicker2,
	}

	fmt.Printf("Max combination: Set\nCards: %v\n", arrSet)
	return true, HandValue{
		Rank:  Set,
		Cards: arrSet,
	}
}

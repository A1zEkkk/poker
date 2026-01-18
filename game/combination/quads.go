package combination

import (
	"fmt"
	. "poker/game/types"
)

func IsQuads(muck []Card) (bool, HandValue) {
	findPairs := make(map[Rank]int)
	var isQuad bool
	var rank Rank

	for _, card := range muck {
		findPairs[card.Rank]++

		if findPairs[card.Rank] == 4 {
			isQuad = true
			rank = card.Rank
		}
	}

	if !isQuad {
		return false, HandValue{}
	}

	kicker := findMaxKicker(findPairs, rank)
	fmt.Printf("Max combination: Quads\nCards: %v\n", []Rank{rank, rank, rank, rank, kicker})
	return true, HandValue{Rank: Quads, Cards: [5]Rank{rank, rank, rank, rank, kicker}}
}

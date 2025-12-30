package combination

import (
	. "pokergame/internal/game/types"
)

func IsPair(muck []Card) (bool, HandValue) {
	maxRank := make(map[Rank]uint8)

	var pair Rank
	var findPair bool

	for _, c := range muck {
		maxRank[c.Rank]++
		if maxRank[c.Rank] == 2 && c.Rank > pair {
			pair = c.Rank
			findPair = true
		}
	}

	if !findPair {
		return false, HandValue{}
	}

	var k1, k2, k3 Rank

	for r := range maxRank {
		if r != pair {
			if r > k1 {
				k3 = k2
				k2 = k1
				k1 = r
			} else if r > k2 {
				k3 = k2
				k2 = r
			} else if r > k3 {
				k3 = r
			}
		}
	}

	return true, HandValue{
		Rank:  Pair,
		Cards: [5]Rank{pair, pair, k1, k2, k3},
	}
}

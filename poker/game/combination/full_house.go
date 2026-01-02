package combination

import (
	. "pokergame/poker/game/types"
)

func IsFullHouse(muck []Card) (bool, HandValue) {
	findPairs := make(map[Rank]int)
	for _, card := range muck {
		findPairs[card.Rank]++
	}

	var hasTrips, hasPair bool
	var tripRank, pairRank Rank

	for rank, count := range findPairs {
		if count == 3 {
			if hasTrips {
				if rank > tripRank {
					pairRank = tripRank
					tripRank = rank
				} else {
					pairRank = rank
				}
				hasPair = true
			} else {
				hasTrips = true
				tripRank = rank
			}
		}

	}

	if !hasTrips {
		return false, HandValue{}
	}

	if !hasPair {
		for rank, count := range findPairs {
			if count >= 2 && rank != tripRank && rank > pairRank {
				pairRank = rank
				hasPair = true
			}
		}
	}

	if !hasPair {
		return false, HandValue{}
	}
	var arrFullHouse = [5]Rank{tripRank, tripRank, tripRank, pairRank, pairRank}

	return true, HandValue{Rank: FullHouse, Cards: arrFullHouse}
}

package combination

import (
	. "pokergame/internal/game/types"
)

func IsStreet(muck []Card) (bool, HandValue) {
	var maxStreet []Card
	var findStreet bool

	for start := 0; start <= len(muck)-5; start++ {
		if isStreetRow(muck[start : start+5]) {
			maxStreet = muck[start : start+5]
			findStreet = true
		}
	}

	if findStreet {
		var arrStraight [5]Rank
		for i := 0; i < 5; i++ {
			arrStraight[i] = maxStreet[i].Rank
		}
		return true, HandValue{
			Rank:  Straight,
			Cards: arrStraight,
		}
	}

	return false, HandValue{}
}

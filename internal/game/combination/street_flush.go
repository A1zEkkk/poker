package combination

import (
	. "pokergame/internal/game/types"
)

func isStreetFlush(muck []Card) ([]Rank, bool) {
	if len(muck) < 5 {
		return nil, false
	}

	var bestSuit Suit                // лучшая масть
	var maxTotal int = 0             // макс кол-во карт одной масти
	suitsCount := make(map[Suit]int) // хэш таблица
	oneSuitRank := make([]Rank, 0)   // Массив карт одной масти

	for _, c := range muck { //Перебор
		suitsCount[c.Suit]++
		bestSuit = c.Suit
	}

	for k, v := range suitsCount { //Поиск лучшей масти и кол-ва карт
		if v > maxTotal {
			maxTotal = v
			bestSuit = k
		}
	}

	if maxTotal < 5 {
		return nil, false
	}

	for _, m := range muck {
		if m.Suit == bestSuit {
			oneSuitRank = append(oneSuitRank, m.Rank)
		}
	}

	SortMyRanks(oneSuitRank)

	maxSubSequence, isFind := FindMaxSubSequence(oneSuitRank)
	if isFind {
		return maxSubSequence, true
	}
	return nil, false
}

package combination

import (
	. "pokergame/internal/game/types"
	"slices"
)

func IsRoyalFlush(muck []Card) (bool, HandValue) { //Флеш рояль
	if len(muck) < 5 {
		return false, HandValue{}
	}

	var bestSuit Suit                                 // лучшая масть
	var maxTotal int = 0                              // макс кол-во карт одной масти
	suitsCount := make(map[Suit]int)                  // хэш таблица
	oneSuitRank := make([]Rank, 0)                    // Массив карт одной масти
	royalFlush := []Rank{Ten, Jack, Queen, King, Ace} // массив для получения флеш рояля

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
		return false, HandValue{}
	}

	for _, m := range muck {
		if m.Suit == bestSuit {
			oneSuitRank = append(oneSuitRank, m.Rank)
		}
	}

	SortMyRanks(oneSuitRank)

	if len(oneSuitRank) > 5 {
		trimIndex := len(oneSuitRank) - 5
		oneSuitRank = oneSuitRank[trimIndex:]
	}

	var arr [5]Rank
	copy(arr[:], oneSuitRank)

	return slices.Equal(royalFlush, oneSuitRank), HandValue{Rank: RoyalFlush, Cards: arr}
}

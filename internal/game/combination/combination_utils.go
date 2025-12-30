package combination

import (
	. "pokergame/internal/game/types"
	"sort"
)

func SortMyRanks(myRanks []Rank) { // залупень для сортировки карт от меньшего к большему

	// Используем sort.Slice, чтобы определить логику сравнения
	sort.Slice(myRanks, func(i, j int) bool {
		// Эта функция должна возвращать true, если элемент с индексом i
		// должен стоять ПЕРЕД элементом с индексом j.

		// Сортировка по возрастанию (от меньшего ранга к большему):
		return myRanks[i] < myRanks[j]
	})
}

func FindMaxSubSequence(sequence []Rank) ([]Rank, bool) {
	// Убедимся, что sequence уже отсортирован и не содержит дубликатов!
	if len(sequence) < 5 {
		return nil, false
	}

	// streakCount - счетчик текущей последовательности
	streakCount := 1

	// Запоминаем индекс, с которого начинается текущая последовательность
	start_i := 0

	// Итерируемся до предпоследнего элемента
	for i := 0; i+1 < len(sequence); i++ {

		// Проверяем, что разница между соседними рангами равна 1
		if sequence[i+1]-sequence[i] == 1 {
			streakCount++
		} else {
			// Если последовательность прервана, сбрасываем счетчик и
			// начинаем отсчет с i+1
			streakCount = 1
			start_i = i + 1
		}

		// Если мы нашли 5 последовательных рангов
		if streakCount >= 5 {
			// Возвращаем найденную подпоследовательность
			// [start_i : текущий индекс + 1]
			return sequence[start_i : i+2], true
		}
	}

	return nil, false
}

// Функция поиска максимального кика
func findMaxKicker(cards map[Rank]int, excludeRank Rank) Rank {
	maxKicker := Rank(0)

	for rank := range cards {
		if rank == excludeRank {
			continue
		}

		if rank > maxKicker {
			maxKicker = rank
		}
	}

	return maxKicker
}

func isStreetRow(row []Card) bool {
	for i := 0; i < len(row)-1; i++ {
		if row[i+1].Rank != row[i].Rank+1 {
			return false
		}
	}
	return true
}

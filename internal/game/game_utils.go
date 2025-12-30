package game

import (
	"fmt"
	. "pokergame/internal/game/types"
	"sort"
)

// Функция для раздачи первых 3 карт на стол
func DealFlop(g *Game) {
	fmt.Println("Dealer deals the flop")
	fmt.Println("Сожжение карты и раздача 3")
	burn := len(g.Deck) - 1
	fmt.Printf("Сожжена карта: %v\n", g.Deck[burn])
	g.Deck = g.Deck[:burn]

	for i := 0; i < 3; i++ {
		pop := len(g.Deck) - 1
		fmt.Printf("Разложена карта: %v \n", g.Deck[pop])
		g.CommunityCard = append(g.CommunityCard, g.Deck[pop])
		g.Deck = g.Deck[:pop]
	}
}

// Функция для раздачи 4 карты на стол
func DealTurn(g *Game) {
	fmt.Println("Dealer deals the turn")
	burn := len(g.Deck) - 1
	fmt.Printf("Сожжена карта: %v\n", g.Deck[burn])
	g.Deck = g.Deck[:burn]

	pop := len(g.Deck) - 1
	fmt.Printf("Разложена карта: %v \n", g.Deck[pop])
	g.CommunityCard = append(g.CommunityCard, g.Deck[pop])
	g.Deck = g.Deck[:pop]
}

// Функция для раздачи 5 карты на стол
func DealRiver(g *Game) {
	fmt.Println("Dealer deals the river")
	burn := len(g.Deck) - 1
	fmt.Printf("Сожжена карта: %v\n", g.Deck[burn])
	g.Deck = g.Deck[:burn]

	pop := len(g.Deck) - 1
	fmt.Printf("Разложена карта: %v \n", g.Deck[pop])
	g.CommunityCard = append(g.CommunityCard, g.Deck[pop])
	g.Deck = g.Deck[:pop]
}

// Реализация функции сортировки для mock
type ByRankThenSuit []Card

func (a ByRankThenSuit) Len() int {
	return len(a)
}

func (a ByRankThenSuit) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByRankThenSuit) Less(i, j int) bool {
	if a[i].Rank != a[j].Rank {
		return a[i].Rank < a[j].Rank
	}
	return a[i].Suit < a[j].Suit
}

func SortMock(mock []Card) {
	fmt.Println("---------------------------------------------------------")
	fmt.Println("Карты до сортировки")
	fmt.Printf("%v\n", mock)
	sort.Sort(ByRankThenSuit(mock))
	fmt.Println("Карты после сортировки")
	fmt.Printf("%v\n", mock)
	fmt.Println("---------------------------------------------------------")
}

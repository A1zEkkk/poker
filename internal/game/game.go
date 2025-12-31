package game

import (
	"fmt"
	"math/rand"
	. "pokergame/internal/game/combination"
	. "pokergame/internal/game/types"
	"time"
)

//Получение колоды карт и ее перемешивание

type Game struct {
	GameId        int
	Deck          []Card
	Players       []User
	Dealer        int
	CommunityCard []Card
}

func (g *Game) GetDeck() { //Генерация колоды
	g.Deck = make([]Card, 0, 52)

	for r := Two; r <= Ace; r++ {
		for s := Spides; s <= Clubs; s++ {
			g.Deck = append(g.Deck, Card{Rank: r, Suit: s})
		}
	}
}

func (g *Game) ShuffleDeck() { //Перемешивание карт
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	n := len(g.Deck)
	for i := n - 1; i > 0; i-- {
		j := r.Intn(i + 1)
		g.Deck[i], g.Deck[j] = g.Deck[j], g.Deck[i]
	}

	fmt.Println("-------------------------------------------")
	fmt.Printf("Игровая колода: %v \n", g.Deck)
	fmt.Println("-------------------------------------------")
}

func (g *Game) GiveCardToHand() { // Раздача карт по рукам
	for cycle := 0; cycle < 2; cycle++ {
		for i := range g.Players {
			lastIndex := len(g.Deck) - 1
			card := g.Deck[lastIndex]
			g.Players[i].Hand = append(g.Players[i].Hand, card)
			g.Deck = g.Deck[:lastIndex]

			fmt.Printf("Игрок id: %v \n получил карту на руку: %v \n", g.Players[i].Id, card)
		}
		fmt.Println("-------------------------------------------")
	}
}

func (g *Game) DealBoard() { // Логика раздачи карт сжигание + раскрытие
	if len(g.CommunityCard) == 5 {
		return
	}

	switch len(g.CommunityCard) {
	case 0:
		dealFlop(g)
	case 3:
		dealTurn(g)

	case 4:
		dealRiver(g)

	default:
		return
	}
}

func (g *Game) GetWinners() []User {
	muck := make([]Card, 0)
	for i := range g.Players {
		muck = nil
		muck = append(muck, g.Players[i].Hand...)
		muck = append(muck, g.CommunityCard...)
		SortMock(muck)
		fmt.Printf("Hand + board: %v \n", muck)

		// Нахожденией лучшей комбинации для каждого пользователя

		if ok, с := IsRoyalFlush(muck); ok {
			g.Players[i].WinComb = с
		} else if ok, с := IsStreetFlush(muck); ok {
			g.Players[i].WinComb = с
		} else if ok, с := IsQuads(muck); ok {
			g.Players[i].WinComb = с
		} else if ok, c := IsFullHouse(muck); ok {
			g.Players[i].WinComb = c
		} else if ok, с := IsFlush(muck); ok {
			g.Players[i].WinComb = с
		} else if ok, c := IsStreet(muck); ok {
			g.Players[i].WinComb = c
		} else if ok, с := IsSet(muck); ok {
			g.Players[i].WinComb = с
		} else if ok, с := IsTwoPair(muck); ok {
			g.Players[i].WinComb = с
		} else if ok, с := IsPair(muck); ok {
			g.Players[i].WinComb = с
		} else {
			_, с := IsHighCard(muck)
			g.Players[i].WinComb = с
		}

		fmt.Printf("Win combination: %v \n for user {%v} \n", g.Players[i].WinComb, g.Players[i].Id)
	}

	var winners []User
	maxRank := g.Players[0].WinComb.Rank

	// Шаг 1: находим максимальный ранг комбинации
	for i := 1; i < len(g.Players); i++ {
		if g.Players[i].WinComb.Rank > maxRank {
			maxRank = g.Players[i].WinComb.Rank
		}
	}

	// Шаг 2: отбираем всех игроков с этим рангом
	candidates := []User{}
	for _, p := range g.Players {
		if p.WinComb.Rank == maxRank {
			candidates = append(candidates, p)
		}
	}

	// Шаг 3: если один кандидат — он победитель
	if len(candidates) == 1 {
		winners = append(winners, candidates[0])
		return winners
	}

	// Шаг 4: сравниваем киккеры
	for i, p := range candidates {
		if i == 0 {
			winners = append(winners, p)
			continue
		}

		// Сравнение карт по старшинству
		equal := true
		for j := 0; j < len(p.WinComb.Cards); j++ {
			if p.WinComb.Cards[j] > winners[0].WinComb.Cards[j] {
				// Новый игрок сильнее — заменяем всех
				winners = []User{p}
				equal = false
				break
			} else if p.WinComb.Cards[j] < winners[0].WinComb.Cards[j] {
				// Новый игрок слабее — пропускаем
				equal = false
				break
			}
		}

		if equal {
			// Киккеры одинаковые — добавляем в список
			winners = append(winners, p)
		}
	}

	return winners
}

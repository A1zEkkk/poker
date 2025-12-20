package game

import (
	"math/rand"
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
			g.Deck = append(g.Deck, Card{r, s})
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
}

func (g *Game) GiveCardToHand() { // Раздача карт по рукам
	for cycle := 0; cycle < 2; cycle++ {
		for i := range g.Players {
			lastIndex := len(g.Deck) - 1
			card := g.Deck[lastIndex]
			g.Players[i].Hand = append(g.Players[i].Hand, card)
			g.Deck = g.Deck[:lastIndex]
		}
	}
}

func (g *Game) DealBoard() { // Логика раздачи карт сжигание + раскрытие
	if len(g.CommunityCard) == 5 {
		return
	}

	switch len(g.CommunityCard) {
	case 0:
		DealFlop(g)

	case 3:
		DealTurn(g)

	case 4:
		DealRiver(g)

	default:
		return
	}
}

// func (g *Game) GetWinners() {
// 	winner := make([]User, 0)
// 	board := make([]Card, 0)
// 	for _, i := range g.Players {
// 		board = g.CommunityCard
// 		muck := append(board, i.Hand...) Нужно сделать сортировку, что бы не делать в каждой ф-ции
// 	}

// }

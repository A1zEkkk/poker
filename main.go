package main

import (
	// Импорт: Имя модуля (pokergame) + путь к папке (pokergame/internal/game)
	"fmt"
	"pokergame/internal/game"
)

func main() {
	g := game.Game{
		GameId: 10,
		Deck:   []game.Card{},
		Players: []game.User{
			{Id: 10, Balance: 4.0},
			{Id: 18, Balance: 22.0},
		},
		Dealer: 0,
	}

	// Выполняем ваши действия
	g.GetDeck()
	g.ShuffleDeck()
	g.GiveCardToHand()

	// Ваш вывод отладки будет немедленно показан
	fmt.Printf("Колода: %v\n", g.Deck)
	fmt.Printf("Игроки: %v\n", g.Players)
}

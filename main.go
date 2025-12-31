package main

import (
	"fmt"
	"pokergame/internal/game"
	. "pokergame/internal/game/types"
)

// Этот ебанный файл сейчас нихуя не показатель тут просто все тестим
// texas holdem

func main() {
	g := game.Game{
		GameId: 10,
		Deck:   []Card{},
		Players: []User{
			{Id: 10, Balance: 4.0},
			{Id: 18, Balance: 22.0},
			{Id: 24, Balance: 67.2},
		},
		Dealer: 0,
	}

	// Подготавливаем колоду и раздаём карты
	g.GetDeck()        // формируем стандартную колоду 52 карты
	g.ShuffleDeck()    // тасуем колоду
	g.GiveCardToHand() // раздаём игрокам по 2 карты

	// Раздаём флоп, терн и ривер последовательно
	for len(g.CommunityCard) < 5 {
		g.DealBoard() // вызываем метод, который раздаёт по правилам
		fmt.Printf("Карты на столе: %v\n", g.CommunityCard)
	}

	// Определяем выигрышные комбинации для каждого игрока
	winner := g.GetWinners()
	// Тут нужно будет после реализации сервера осуществить работу с балансом ставки + вывод денег. Сейчас все ок

	// Выводим результаты для проверки
	for _, player := range g.Players {
		fmt.Printf("Игрок ID: %d\n", player.Id)
		fmt.Printf("Карты на руках: %v\n", player.Hand)
		fmt.Printf("Комбинация: %v\n", player.WinComb)
		fmt.Println("-------------------------------------------------")
	}
	fmt.Printf("winner or winners: %+v \n", winner)
}

// Нужно сделать функцию и реализовать логику, которая будет позволять взаимодейстовать с балансом т.е после определения победителя или победителей иенять баланс у структуры

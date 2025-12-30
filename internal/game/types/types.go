package types

//Тут перечислины часто используемые типы данных и структуры

type Rank int
type Suit int

const ( // Масти карт
	Spides Suit = iota
	Hearts
	Diamonds
	Clubs
)

const ( // Ранг карт
	Two Rank = iota + 2
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

type HandRank int

const ( // Реализовать запись в бд с преобразованием в читамый вид
	HighCard HandRank = iota
	Pair
	TwoPair
	Set
	Straight
	Flush
	FullHouse
	Quads
	StraightFlush
	RoyalFlush
)

type Card struct { // Ранг карты от 2 до туза и масть
	Rank Rank
	Suit Suit
}

type HandValue struct { // Тип комбинации, которая выпала. Тут ранг равен виду комбинации от старшей карты до флеш рояля. Ну и массив рангов карт
	Rank  HandRank
	Cards [5]Rank
}

type User struct { // Данные о пользователе. Сейчас они предварительные
	Id      int
	Balance float64
	Hand    []Card
	WinComb HandValue
}

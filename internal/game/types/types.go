package types

type Rank int
type Suit int

const (
	Spides Suit = iota
	Hearts
	Diamonds
	Clubs
)

const (
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

type Card struct {
	Rank Rank
	Suit Suit
}

type User struct {
	Id      int
	Balance float64
	Hand    []Card
}

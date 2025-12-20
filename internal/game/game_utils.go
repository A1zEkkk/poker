package game

// Функция для раздачи первых 3 карт на стол
func DealFlop(g *Game) {
	burn := len(g.Deck) - 1
	g.Deck = g.Deck[:burn]
	i := 0
	for ; i < 3; i++ {
		pop := len(g.Deck) - 1
		g.CommunityCard = append(g.CommunityCard, g.Deck[pop])
		g.Deck = g.Deck[:pop]
	}
}

// Функция для раздачи 4 краты на стол
func DealTurn(g *Game) {
	burn := len(g.Deck) - 1
	g.Deck = g.Deck[:burn]
	pop := len(g.Deck) - 1
	g.CommunityCard = append(g.CommunityCard, g.Deck[pop])
	g.Deck = g.Deck[:pop]
}

// Функция для раздачи 5 карты на стол
func DealRiver(g *Game) {
	burn := len(g.Deck) - 1
	g.Deck = g.Deck[:burn]
	pop := len(g.Deck) - 1
	g.CommunityCard = append(g.CommunityCard, g.Deck[pop])
	g.Deck = g.Deck[:pop]
}

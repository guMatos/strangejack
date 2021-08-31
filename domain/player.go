package domain

type Player struct {
	Name       string
	Cards      []Card
	ChipsCount int
}

func NewPlayer(name string, cards []Card) *Player {
	return &Player{
		Name:       name,
		Cards:      cards,
		ChipsCount: 3,
	}
}

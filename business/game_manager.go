package business

import (
	"fmt"

	"github.com/guMatos/black-jack/domain"
	"github.com/mitchellh/mapstructure"

	"github.com/ahmetb/go-linq/v3"
)

func PlayBlackjack() {
	fmt.Println("\n##### BLACKJACK #####")

	blackjack := NewBlackjack()

	var winner domain.Player
	for winner.Name == "" {
		blackjack.PlayTurn()

		var losers []domain.Player
		linq.From(blackjack.Players).WhereT(func(p domain.Player) bool {
			return p.ChipsCount <= 0
		}).ToSlice(&losers)

		if len(blackjack.Players)-len(losers) == 1 {
			winnerI := linq.From(blackjack.Players).FirstWithT(func(p domain.Player) bool {
				return p.ChipsCount > 0
			})
			mapstructure.Decode(winnerI, &winner)
		}
	}

	fmt.Println(fmt.Sprintf("\n%v is the winner with %v chips ", winner.Name, winner.ChipsCount))
}

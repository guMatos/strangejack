package business

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/guMatos/black-jack/domain"
	"github.com/guMatos/black-jack/infra"
)

type Blackjack struct {
	Players    []domain.Player
	DeckId     string
	DeckClient *infra.DeckClient
	Reader     *bufio.Reader
	TableSum   int
}

func NewBlackjack() *Blackjack {
	client := infra.NewDeckClient()
	deck := client.Shuffle(1)

	reader := bufio.NewReader(os.Stdin)

	fmt.Print(fmt.Sprintf("How many players? "))
	number, _ := reader.ReadString('\n')
	number = strings.TrimRight(number, "\r\n")
	numberOfPlayers, err := strconv.Atoi(number)
	if err != nil {
		log.Fatalln(err)
	}
	if numberOfPlayers < 2 || numberOfPlayers > 6 {
		fmt.Println("\nMust be between 2 and 6")
		return NewBlackjack()
	}

	var players []domain.Player
	for i := 0; i < numberOfPlayers; i++ {
		fmt.Print(fmt.Sprintf("Enter name of player %v: ", i+1))
		name, _ := reader.ReadString('\n')
		name = strings.TrimRight(name, "\r\n")
		cards := client.Draw(deck.DeckId, 3).Cards

		players = append(players, *domain.NewPlayer(name, cards))
	}

	return &Blackjack{
		Players:    players,
		DeckId:     deck.DeckId,
		DeckClient: client,
		Reader:     reader,
		TableSum:   0,
	}
}

func (game *Blackjack) PlayTurn() {
	for i, player := range game.Players {
		fmt.Print(fmt.Sprintf("\n%v. %v chips", player.Name, player.ChipsCount))

		if player.ChipsCount <= 0 {
			continue
		}

		validPlay := false
		for validPlay == false {
			fmt.Print(fmt.Sprintf("\n%v, which one of your cards you want to play? ", player.Name))
			number, _ := game.Reader.ReadString('\n')
			number = strings.TrimRight(number, "\r\n")
			cardNumber, err := strconv.Atoi(number)
			if err != nil || cardNumber < 1 || cardNumber > 3 {
				fmt.Println("Must be between 1 and 3")
				continue
			}

			card := &player.Cards[cardNumber-1]
			cardNumericValue := GetCardNumericValue(*card)
			fmt.Println(fmt.Sprintf("%v of %v. Numeric value: %v", card.Value, card.Suit, cardNumericValue))

			game.TableSum += cardNumericValue
			fmt.Println(fmt.Sprintf("Table sum: %v", game.TableSum))

			if game.TableSum == 21 {
				game.TableSum = 0
				player.ChipsCount++
				fmt.Println(fmt.Sprintf("\n%v achieved 21 points and won a chip. ", player.Name))
				fmt.Println(fmt.Sprintf("\nCurrent table sum: %v", game.TableSum))
			} else if game.TableSum > 21 {
				difference := game.TableSum - 21
				game.TableSum = difference
				player.ChipsCount--
				fmt.Println(fmt.Sprintf("\n%v surpassed 21 points and lost a chip. ", player.Name))
				fmt.Println(fmt.Sprintf("\nCurrent table sum: %v", game.TableSum))
			}

			game.Players[i].Cards = append(game.Players[i].Cards[:cardNumber-1])
			drawResponse := game.DeckClient.Draw(game.DeckId, 1)
			if !drawResponse.Success {
				deckResponse := game.DeckClient.Shuffle(1)
				game.DeckId = deckResponse.DeckId
				drawResponse = game.DeckClient.Draw(game.DeckId, 1)
			}
			game.Players[i].Cards = append(game.Players[i].Cards, drawResponse.Cards...)

			validPlay = true
		}
	}
}

func GetCardNumericValue(card domain.Card) int {
	if card.Value == "A" {
		return 1
	}

	numericValue, err := strconv.Atoi(card.Value)
	if err != nil || numericValue == 0 {
		return 10
	}

	return numericValue
}

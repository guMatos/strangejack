package infra

import (
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/viper"

	"github.com/guMatos/black-jack/domain"
)

type DeckClient struct {
	Client      http.Client
	BaseAddress string
}

func NewDeckClient() *DeckClient {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	return &DeckClient{
		Client:      http.Client{Timeout: time.Second * 10},
		BaseAddress: viper.GetString("API_URL"),
	}
}

func (client DeckClient) Shuffle(count int) domain.ShuffleResponse {
	route := fmt.Sprintf("%v/api/deck/new/shuffle/?deck_count=%v", client.BaseAddress, count)
	response := Get(route, client.Client)

	defer response.Body.Close()

	var shuffle domain.ShuffleResponse
	DecodeHttpResponse(*response, &shuffle)

	return shuffle
}

func (client DeckClient) Draw(deckId string, count int) domain.DrawResponse {
	route := fmt.Sprintf("%v/api/deck/%v/draw/?count=%v", client.BaseAddress, deckId, count)
	response := Get(route, client.Client)

	defer response.Body.Close()

	var draw domain.DrawResponse
	DecodeHttpResponse(*response, &draw)

	return draw
}

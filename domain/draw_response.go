package domain

type DrawResponse struct {
	Success   bool   `json:"success"`
	Cards     []Card `json:"cards"`
	DeckId    string `json:"deck_id"`
	Remaining int    `json:"remaining"`
}

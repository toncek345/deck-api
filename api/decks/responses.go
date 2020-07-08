package decks

import "errors"

var (
	ErrNotFound      = errors.New("Object not found")
	ErrBadRequest    = errors.New("Bad request")
	ErrAllCardsDrawn = errors.New("No cards left of too much drawn")
)

type DrawResponse struct {
	Cards []*Card `json:"cards"`
}

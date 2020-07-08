package decks

import (
	"deck-api/service/deck"
)

type Suit string

const (
	SuitSpades   Suit = "SPADES"
	SuitDiamonds Suit = "DIAMONDS"
	SuitClubs    Suit = "CLUBS"
	SuitHearts   Suit = "HEARTS"
)

func SuitResp(suit deck.Suit) Suit {
	switch suit {
	case deck.SuitClubs:
		return SuitClubs
	case deck.SuitDiamonds:
		return SuitDiamonds
	case deck.SuitHearts:
		return SuitHearts
	case deck.SuitSpades:
		return SuitSpades
	}

	return ""
}

type Card struct {
	Value string `json:"value"`
	Suit  Suit   `json:"suit"`
	Code  string `json:"code"`
}

func CardsResp(cards []*deck.Card) []*Card {
	c := make([]*Card, 0, len(cards))

	for _, v := range cards {
		c = append(
			c,
			&Card{
				Value: v.Value,
				Suit:  SuitResp(v.Suit),
				Code:  v.Code,
			})
	}

	return c
}

type Deck struct {
	ID        string  `json:"deck_id"`
	Shuffled  bool    `json:"shuffled"`
	Remaining int     `json:"remaining"`
	Cards     []*Card `json:"cards"`
}

package deck

import (
	"errors"
	"math/rand"
)

type Suit string

const (
	SuitSpades   Suit = "SPADES"
	SuitDiamonds Suit = "DIAMONDS"
	SuitClubs    Suit = "CLUBS"
	SuitHearts   Suit = "HEARTS"
)

func AllSuits() []Suit {
	return []Suit{SuitSpades, SuitDiamonds, SuitClubs, SuitHearts}
}

type Deck struct {
	ID       string
	Shuffled bool

	Cards []*Card
	drawn int
}

func (d *Deck) shuffleCards() {
	for i := range d.Cards {
		// Pick random from remaining cards and swap with current card.
		pick := rand.Int() % (len(d.Cards) - i)

		card := d.Cards[i]
		d.Cards[i] = d.Cards[pick]
		d.Cards[pick] = card
	}

	d.Shuffled = true
}

func (d *Deck) RemainingCards() []*Card {
	return d.Cards[d.drawn:]
}

type Card struct {
	Value string
	Suit  Suit
	Code  string
}

var (
	ErrNotFound    = errors.New("service/deck: not found")
	ErrUnknownCode = errors.New("service/deck: unknown code")
	ErrNoCards     = errors.New("service/deck: no cards left in deck")
)

package deck

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

// Making sure that memory implements whole deck service.
var _ Service = &Memory{}

type Memory struct {
	Decks map[string]*Deck
}

func NewMemory() Service {
	return &Memory{
		Decks: make(map[string]*Deck),
	}
}

func (m *Memory) Create(ctx context.Context, shuffle bool, cardCodes ...string) (*Deck, error) {
	cards, err := generateCards(cardCodes...)
	if err != nil {
		return nil, fmt.Errorf("service/deck: error generating cards: %w", err)
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, fmt.Errorf("service/deck: error generating id: %w", err)
	}

	deck := &Deck{
		ID:    id.String(),
		Cards: cards,
	}

	if shuffle {
		deck.shuffleCards()
	}

	m.Decks[deck.ID] = deck

	return deck, nil
}

func (m *Memory) Draw(ctx context.Context, deckID string, ntimes int) ([]*Card, error) {
	deck, err := m.Open(ctx, deckID)
	if err != nil {
		return nil, fmt.Errorf("serivce/deck: error getting deck: %w", err)
	}

	if (deck.drawn + ntimes) > len(deck.Cards) {
		return nil, fmt.Errorf("service/deck: too many cards drawn or no cards left: %w",
			ErrNoCards)
	}

	cards := deck.Cards[deck.drawn : deck.drawn+ntimes]
	deck.drawn += ntimes

	return cards, nil
}

func (m *Memory) Open(ctx context.Context, deckID string) (*Deck, error) {
	deck, ok := m.Decks[deckID]
	if !ok {
		return nil, fmt.Errorf("service/deck: deck not found: %w", ErrNotFound)
	}

	return deck, nil
}

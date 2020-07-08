package deck

import (
	"context"
)

var _ Service = &Mock{}

type Service interface {
	// Create creates deck and returnes it's instance. Card codes is optional parameter
	// which contains only card codes as 4S (4 spades), KH (king hearts).
	Create(ctx context.Context, shuffle bool, cardCodes ...string) (*Deck, error)

	// Draw draws cards from the deck.
	Draw(ctx context.Context, deckID string, ntimes int) ([]*Card, error)

	// Open returns the deck with current status.
	Open(ctx context.Context, deckID string) (*Deck, error)
}

type Mock struct {
	CreateFn func(context.Context, bool, ...string) (*Deck, error)
	DrawFn   func(context.Context, string, int) ([]*Card, error)
	OpenFn   func(context.Context, string) (*Deck, error)
}

func (m *Mock) Create(ctx context.Context, shuffle bool, cardCodes ...string) (*Deck, error) {
	return m.CreateFn(ctx, shuffle, cardCodes...)
}

func (m *Mock) Draw(ctx context.Context, deckID string, ntimes int) ([]*Card, error) {
	return m.DrawFn(ctx, deckID, ntimes)
}

func (m *Mock) Open(ctx context.Context, deckID string) (*Deck, error) {
	return m.OpenFn(ctx, deckID)
}

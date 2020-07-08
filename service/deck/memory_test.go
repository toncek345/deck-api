package deck

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	t.Run("without card codes", func(t *testing.T) {
		service := NewMemory()

		deck, err := service.Create(context.Background(), false)
		assert.Nil(t, err)
		assert.Equal(
			t,
			&Deck{
				ID:       deck.ID,
				drawn:    0,
				Shuffled: false,
				Cards: []*Card{
					{Value: "ACE", Suit: "SPADES", Code: "AS"},
					{Value: "2", Suit: "SPADES", Code: "2S"},
					{Value: "3", Suit: "SPADES", Code: "3S"},
					{Value: "4", Suit: "SPADES", Code: "4S"},
					{Value: "5", Suit: "SPADES", Code: "5S"},
					{Value: "6", Suit: "SPADES", Code: "6S"},
					{Value: "7", Suit: "SPADES", Code: "7S"},
					{Value: "8", Suit: "SPADES", Code: "8S"},
					{Value: "9", Suit: "SPADES", Code: "9S"},
					{Value: "10", Suit: "SPADES", Code: "10S"},
					{Value: "JACK", Suit: "SPADES", Code: "JS"},
					{Value: "QUEEN", Suit: "SPADES", Code: "QS"},
					{Value: "KING", Suit: "SPADES", Code: "KS"},
					{Value: "ACE", Suit: "DIAMONDS", Code: "AD"},
					{Value: "2", Suit: "DIAMONDS", Code: "2D"},
					{Value: "3", Suit: "DIAMONDS", Code: "3D"},
					{Value: "4", Suit: "DIAMONDS", Code: "4D"},
					{Value: "5", Suit: "DIAMONDS", Code: "5D"},
					{Value: "6", Suit: "DIAMONDS", Code: "6D"},
					{Value: "7", Suit: "DIAMONDS", Code: "7D"},
					{Value: "8", Suit: "DIAMONDS", Code: "8D"},
					{Value: "9", Suit: "DIAMONDS", Code: "9D"},
					{Value: "10", Suit: "DIAMONDS", Code: "10D"},
					{Value: "JACK", Suit: "DIAMONDS", Code: "JD"},
					{Value: "QUEEN", Suit: "DIAMONDS", Code: "QD"},
					{Value: "KING", Suit: "DIAMONDS", Code: "KD"},
					{Value: "ACE", Suit: "CLUBS", Code: "AC"},
					{Value: "2", Suit: "CLUBS", Code: "2C"},
					{Value: "3", Suit: "CLUBS", Code: "3C"},
					{Value: "4", Suit: "CLUBS", Code: "4C"},
					{Value: "5", Suit: "CLUBS", Code: "5C"},
					{Value: "6", Suit: "CLUBS", Code: "6C"},
					{Value: "7", Suit: "CLUBS", Code: "7C"},
					{Value: "8", Suit: "CLUBS", Code: "8C"},
					{Value: "9", Suit: "CLUBS", Code: "9C"},
					{Value: "10", Suit: "CLUBS", Code: "10C"},
					{Value: "JACK", Suit: "CLUBS", Code: "JC"},
					{Value: "QUEEN", Suit: "CLUBS", Code: "QC"},
					{Value: "KING", Suit: "CLUBS", Code: "KC"},
					{Value: "ACE", Suit: "HEARTS", Code: "AH"},
					{Value: "2", Suit: "HEARTS", Code: "2H"},
					{Value: "3", Suit: "HEARTS", Code: "3H"},
					{Value: "4", Suit: "HEARTS", Code: "4H"},
					{Value: "5", Suit: "HEARTS", Code: "5H"},
					{Value: "6", Suit: "HEARTS", Code: "6H"},
					{Value: "7", Suit: "HEARTS", Code: "7H"},
					{Value: "8", Suit: "HEARTS", Code: "8H"},
					{Value: "9", Suit: "HEARTS", Code: "9H"},
					{Value: "10", Suit: "HEARTS", Code: "10H"},
					{Value: "JACK", Suit: "HEARTS", Code: "JH"},
					{Value: "QUEEN", Suit: "HEARTS", Code: "QH"},
					{Value: "KING", Suit: "HEARTS", Code: "KH"},
				},
			},
			deck,
			"decks should be same",
		)
	})

	t.Run("with card codes", func(t *testing.T) {
		service := NewMemory()

		deck, err := service.Create(context.Background(), false, "AS", "2H")
		assert.Nil(t, err)
		assert.Equal(
			t,
			&Deck{
				ID:       deck.ID,
				drawn:    0,
				Shuffled: false,
				Cards: []*Card{
					{Value: "ACE", Suit: "SPADES", Code: "AS"},
					{Value: "2", Suit: "HEARTS", Code: "2H"},
				},
			},
			deck,
		)
	})

	t.Run("with invalid card code", func(t *testing.T) {
		service := NewMemory()

		_, err := service.Create(context.Background(), false, "AS", "1H")
		assert.True(t, errors.Is(err, ErrUnknownCode))
	})

	t.Run("with shuffle", func(t *testing.T) {
		// Let's try out our luck and say this test will always pass :)
		service := NewMemory()

		deck, err := service.Create(context.Background(), true)
		assert.Nil(t, err)
		assert.Equal(
			t,
			&Deck{
				ID:       deck.ID,
				Shuffled: true,
				Cards:    deck.Cards,
			},
			deck)
		assert.NotEqual(
			t,
			&Deck{
				ID:       deck.ID,
				drawn:    0,
				Shuffled: true,
				Cards: []*Card{
					{Value: "ACE", Suit: "SPADES", Code: "AS"},
					{Value: "2", Suit: "SPADES", Code: "2S"},
					{Value: "3", Suit: "SPADES", Code: "3S"},
					{Value: "4", Suit: "SPADES", Code: "4S"},
					{Value: "5", Suit: "SPADES", Code: "5S"},
					{Value: "6", Suit: "SPADES", Code: "6S"},
					{Value: "7", Suit: "SPADES", Code: "7S"},
					{Value: "8", Suit: "SPADES", Code: "8S"},
					{Value: "9", Suit: "SPADES", Code: "9S"},
					{Value: "10", Suit: "SPADES", Code: "10S"},
					{Value: "JACK", Suit: "SPADES", Code: "JS"},
					{Value: "QUEEN", Suit: "SPADES", Code: "QS"},
					{Value: "KING", Suit: "SPADES", Code: "KS"},
					{Value: "ACE", Suit: "DIAMONDS", Code: "AD"},
					{Value: "2", Suit: "DIAMONDS", Code: "2D"},
					{Value: "3", Suit: "DIAMONDS", Code: "3D"},
					{Value: "4", Suit: "DIAMONDS", Code: "4D"},
					{Value: "5", Suit: "DIAMONDS", Code: "5D"},
					{Value: "6", Suit: "DIAMONDS", Code: "6D"},
					{Value: "7", Suit: "DIAMONDS", Code: "7D"},
					{Value: "8", Suit: "DIAMONDS", Code: "8D"},
					{Value: "9", Suit: "DIAMONDS", Code: "9D"},
					{Value: "10", Suit: "DIAMONDS", Code: "10D"},
					{Value: "JACK", Suit: "DIAMONDS", Code: "JD"},
					{Value: "QUEEN", Suit: "DIAMONDS", Code: "QD"},
					{Value: "KING", Suit: "DIAMONDS", Code: "KD"},
					{Value: "ACE", Suit: "CLUBS", Code: "AC"},
					{Value: "2", Suit: "CLUBS", Code: "2C"},
					{Value: "3", Suit: "CLUBS", Code: "3C"},
					{Value: "4", Suit: "CLUBS", Code: "4C"},
					{Value: "5", Suit: "CLUBS", Code: "5C"},
					{Value: "6", Suit: "CLUBS", Code: "6C"},
					{Value: "7", Suit: "CLUBS", Code: "7C"},
					{Value: "8", Suit: "CLUBS", Code: "8C"},
					{Value: "9", Suit: "CLUBS", Code: "9C"},
					{Value: "10", Suit: "CLUBS", Code: "10C"},
					{Value: "JACK", Suit: "CLUBS", Code: "JC"},
					{Value: "QUEEN", Suit: "CLUBS", Code: "QC"},
					{Value: "KING", Suit: "CLUBS", Code: "KC"},
					{Value: "ACE", Suit: "HEARTS", Code: "AH"},
					{Value: "2", Suit: "HEARTS", Code: "2H"},
					{Value: "3", Suit: "HEARTS", Code: "3H"},
					{Value: "4", Suit: "HEARTS", Code: "4H"},
					{Value: "5", Suit: "HEARTS", Code: "5H"},
					{Value: "6", Suit: "HEARTS", Code: "6H"},
					{Value: "7", Suit: "HEARTS", Code: "7H"},
					{Value: "8", Suit: "HEARTS", Code: "8H"},
					{Value: "9", Suit: "HEARTS", Code: "9H"},
					{Value: "10", Suit: "HEARTS", Code: "10H"},
					{Value: "JACK", Suit: "HEARTS", Code: "JH"},
					{Value: "QUEEN", Suit: "HEARTS", Code: "QH"},
					{Value: "KING", Suit: "HEARTS", Code: "KH"},
				},
			},
			deck,
		)
	})
}

func TestOpen(t *testing.T) {
	service := NewMemory()
	deck, err := service.Create(context.Background(), false)
	assert.Nil(t, err)

	t.Run("valid", func(t *testing.T) {
		d, err := service.Open(context.Background(), deck.ID)
		assert.Nil(t, err)
		assert.Equal(t, deck, d)
	})

	t.Run("invalid", func(t *testing.T) {
		_, err := service.Open(context.Background(), "invalid id")
		assert.True(t, errors.Is(err, ErrNotFound))
	})
}

func TestDraw(t *testing.T) {
	service := NewMemory()

	t.Run("valid", func(t *testing.T) {
		deck, err := service.Create(context.Background(), false, "AS", "AH", "AD")
		assert.Nil(t, err)

		cards, err := service.Draw(context.Background(), deck.ID, 2)
		assert.Nil(t, err)
		assert.Equal(
			t,
			[]*Card{
				{Value: "ACE", Suit: SuitSpades, Code: "AS"},
				{Value: "ACE", Suit: SuitHearts, Code: "AH"},
			},
			cards,
			"cards should be equal",
		)
		assert.Equal(t, 1, len(deck.RemainingCards()))
	})

	t.Run("drawn too many", func(t *testing.T) {
		deck, err := service.Create(context.Background(), false, "AS", "AH", "AD")
		assert.Nil(t, err)

		_, err = service.Draw(context.Background(), deck.ID, 4)
		assert.True(t, errors.Is(err, ErrNoCards))
	})
}

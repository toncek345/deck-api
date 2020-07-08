package decks

import (
	"bytes"
	"context"
	"deck-api/api/core"
	"deck-api/service/deck"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpenDeck(t *testing.T) {
	t.Run("no args", func(t *testing.T) {
		app := &App{
			DeckService: &deck.Mock{
				CreateFn: func(context.Context, bool, ...string) (*deck.Deck, error) {
					return &deck.Deck{
						ID:       "this is id",
						Shuffled: true,
						Cards: []*deck.Card{
							&deck.Card{Code: "c", Value: "1", Suit: deck.SuitClubs},
						},
					}, nil
				},
			},
		}

		r := httptest.NewRequest(http.MethodPost, "http://test.com/decks", bytes.NewReader([]byte("{}")))
		w := httptest.NewRecorder()
		app.Router().ServeHTTP(w, r)

		assert.Equal(t, http.StatusOK, w.Result().StatusCode)

		var newDeck Deck
		assert.Nil(t, json.NewDecoder(w.Result().Body).Decode(&newDeck))
		assert.Equal(
			t,
			Deck{
				ID:        "this is id",
				Shuffled:  true,
				Remaining: 1,
			},
			newDeck,
		)
	})

	t.Run("shuffled", func(t *testing.T) {
		app := &App{
			DeckService: &deck.Mock{
				CreateFn: func(ctx context.Context, shuffle bool, cardCodes ...string) (*deck.Deck, error) {
					assert.True(t, shuffle)

					return &deck.Deck{
						ID:       "this is id",
						Shuffled: true,
					}, nil
				},
			},
		}

		r := httptest.NewRequest(http.MethodPost, "http://test.com/decks?shuffle=true", nil)
		w := httptest.NewRecorder()
		app.Router().ServeHTTP(w, r)

		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
		var newDeck Deck
		assert.Nil(t, json.NewDecoder(w.Result().Body).Decode(&newDeck))
		assert.Equal(t, Deck{ID: "this is id", Shuffled: true}, newDeck)
	})

	t.Run("with card codes", func(t *testing.T) {
		app := &App{
			DeckService: &deck.Mock{
				CreateFn: func(ctx context.Context, shuffle bool, cardCodes ...string) (*deck.Deck, error) {
					assert.Equal(t, []string{"code1", "code2"}, cardCodes)

					return &deck.Deck{
						ID: "this is id",
					}, nil
				},
			},
		}

		r := httptest.NewRequest(http.MethodPost, "http://test.com/decks?cards=code1,code2", nil)
		w := httptest.NewRecorder()
		app.Router().ServeHTTP(w, r)

		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
		var newDeck Deck
		assert.Nil(t, json.NewDecoder(w.Result().Body).Decode(&newDeck))
		assert.Equal(t, Deck{ID: "this is id"}, newDeck)
	})
}

func TestOpen(t *testing.T) {
	t.Run("works", func(t *testing.T) {
		app := &App{
			DeckService: &deck.Mock{
				OpenFn: func(ctx context.Context, id string) (*deck.Deck, error) {
					assert.Equal(t, "deck-id", id)

					return &deck.Deck{
						ID:       "deck-id",
						Shuffled: true,
						Cards: []*deck.Card{
							&deck.Card{},
						},
					}, nil
				},
			},
		}

		r := httptest.NewRequest(http.MethodGet, "http://test.com/decks/deck-id", nil)
		w := httptest.NewRecorder()
		app.Router().ServeHTTP(w, r)

		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
		var newDeck Deck
		assert.Nil(t, json.NewDecoder(w.Result().Body).Decode(&newDeck))
		assert.Equal(
			t,
			Deck{
				ID:        "deck-id",
				Remaining: 1,
				Shuffled:  true,
				Cards:     []*Card{&Card{}},
			},
			newDeck,
		)
	})

	t.Run("deck not found", func(t *testing.T) {
		app := &App{
			DeckService: &deck.Mock{
				OpenFn: func(ctx context.Context, id string) (*deck.Deck, error) {
					assert.Equal(t, "deck-id2", id)
					return nil, deck.ErrNotFound
				},
			},
		}

		r := httptest.NewRequest(http.MethodGet, "http://test.com/decks/deck-id2", nil)
		w := httptest.NewRecorder()
		app.Router().ServeHTTP(w, r)

		assert.Equal(t, http.StatusNotFound, w.Result().StatusCode)
		var errResp core.ErrorResponse
		assert.Nil(t, json.NewDecoder(w.Result().Body).Decode(&errResp))
		assert.Equal(
			t,
			core.ErrorResponse{Msg: ErrNotFound.Error()},
			errResp,
		)
	})
}

func TestDraw(t *testing.T) {
	t.Run("without args", func(t *testing.T) {
		app := &App{
			DeckService: &deck.Mock{
				DrawFn: func(ctx context.Context, id string, amount int) ([]*deck.Card, error) {
					assert.Equal(t, 1, amount)
					assert.Equal(t, "deck-id", id)

					return []*deck.Card{
						&deck.Card{Value: "1", Code: "1", Suit: deck.SuitDiamonds},
					}, nil
				},
			},
		}

		r := httptest.NewRequest(http.MethodGet, "http://test.com/decks/deck-id/draw", nil)
		w := httptest.NewRecorder()
		app.Router().ServeHTTP(w, r)

		assert.Equal(t, http.StatusOK, w.Result().StatusCode)

		var resp DrawResponse
		assert.Nil(t, json.NewDecoder(w.Result().Body).Decode(&resp))
		assert.Equal(
			t,
			DrawResponse{
				Cards: []*Card{
					&Card{Value: "1", Code: "1", Suit: SuitDiamonds},
				},
			},
			resp,
		)
	})

	t.Run("all cards drawn", func(t *testing.T) {
		app := &App{
			DeckService: &deck.Mock{
				DrawFn: func(ctx context.Context, id string, amount int) ([]*deck.Card, error) {
					assert.Equal(t, 13, amount)
					assert.Equal(t, "deck-id", id)

					return nil, deck.ErrNoCards
				},
			},
		}

		r := httptest.NewRequest(http.MethodGet, "http://test.com/decks/deck-id/draw?amount=13", nil)
		w := httptest.NewRecorder()
		app.Router().ServeHTTP(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)

		var resp core.ErrorResponse
		assert.Nil(t, json.NewDecoder(w.Result().Body).Decode(&resp))
		assert.Equal(
			t,
			core.ErrorResponse{
				Msg: ErrAllCardsDrawn.Error(),
			},
			resp,
		)
	})

	t.Run("with draw amount", func(t *testing.T) {
		app := &App{
			DeckService: &deck.Mock{
				DrawFn: func(ctx context.Context, id string, amount int) ([]*deck.Card, error) {
					assert.Equal(t, 13, amount)
					assert.Equal(t, "deck-id", id)

					return []*deck.Card{
						&deck.Card{Value: "1", Code: "1", Suit: deck.SuitDiamonds},
						&deck.Card{Value: "12", Code: "12", Suit: deck.SuitClubs},
					}, nil
				},
			},
		}

		r := httptest.NewRequest(http.MethodGet, "http://test.com/decks/deck-id/draw?amount=13", nil)
		w := httptest.NewRecorder()
		app.Router().ServeHTTP(w, r)

		assert.Equal(t, http.StatusOK, w.Result().StatusCode)

		var resp DrawResponse
		assert.Nil(t, json.NewDecoder(w.Result().Body).Decode(&resp))
		assert.Equal(
			t,
			DrawResponse{
				Cards: []*Card{
					&Card{Value: "1", Code: "1", Suit: SuitDiamonds},
					&Card{Value: "12", Code: "12", Suit: SuitClubs},
				},
			},
			resp,
		)
	})

	t.Run("with invalid draw amount", func(t *testing.T) {
		app := &App{
			DeckService: &deck.Mock{
				DrawFn: func(ctx context.Context, id string, amount int) ([]*deck.Card, error) {
					return nil, nil
				},
			},
		}

		r := httptest.NewRequest(http.MethodGet, "http://test.com/decks/deck-id/draw?amount=-13", nil)
		w := httptest.NewRecorder()
		app.Router().ServeHTTP(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)

		var resp core.ErrorResponse
		assert.Nil(t, json.NewDecoder(w.Result().Body).Decode(&resp))
		assert.Equal(
			t,
			core.ErrorResponse{Msg: ErrBadRequest.Error()},
			resp,
		)
	})
}

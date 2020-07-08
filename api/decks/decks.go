package decks

import (
	"deck-api/api/core"
	deckService "deck-api/service/deck"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi"
)

type App struct {
	core.API

	DeckService deckService.Service
}

func (app *App) Router() chi.Router {
	r := chi.NewMux()

	r.Route("/decks", func(r chi.Router) {
		r.Route("/{deckID}", func(r chi.Router) {
			r.Get("/", app.OpenDeck)
			r.Get("/draw", app.DrawFromDeck)
		})

		r.Post("/", app.CreateDeck)
	})

	return r
}

func (app *App) OpenDeck(w http.ResponseWriter, r *http.Request) {
	d, err := app.DeckService.Open(r.Context(), chi.URLParam(r, "deckID"))
	if err != nil {
		if errors.Is(err, deckService.ErrNotFound) {
			app.RespondError(w, ErrNotFound, http.StatusNotFound)
			return
		}

		app.RespondError(
			w,
			fmt.Errorf("api/decks: error opening deck: %w", err),
			http.StatusInternalServerError,
		)
		return
	}

	app.RespondJSON(
		w,
		Deck{
			ID:        d.ID,
			Shuffled:  d.Shuffled,
			Remaining: len(d.RemainingCards()),
			Cards:     CardsResp(d.RemainingCards()),
		},
	)
}

func (app *App) DrawFromDeck(w http.ResponseWriter, r *http.Request) {
	drawAmount := r.URL.Query().Get("amount")

	var drawAmountParsed int
	if drawAmount != "" {
		var err error
		drawAmountParsed, err = strconv.Atoi(drawAmount)
		if err != nil {
			app.RespondError(w, ErrBadRequest, http.StatusBadRequest)
			return
		}

		if drawAmountParsed < 0 {
			app.RespondError(w, ErrBadRequest, http.StatusBadRequest)
			return
		}
	} else {
		drawAmountParsed = 1
	}

	cards, err := app.DeckService.Draw(
		r.Context(),
		chi.URLParam(r, "deckID"),
		drawAmountParsed,
	)
	if err != nil {
		if errors.Is(err, deckService.ErrNotFound) {
			app.RespondError(w, ErrNotFound, http.StatusNotFound)
			return
		}

		if errors.Is(err, deckService.ErrNoCards) {
			app.RespondError(w, ErrAllCardsDrawn, http.StatusBadRequest)
			return
		}

		app.RespondError(
			w,
			fmt.Errorf("api/decks: error drawing card: %w", err),
			http.StatusInternalServerError,
		)
		return
	}

	app.RespondJSON(w, &DrawResponse{Cards: CardsResp(cards)})
}

func (app *App) CreateDeck(w http.ResponseWriter, r *http.Request) {
	card := r.URL.Query().Get("cards")
	shuffle := r.URL.Query().Get("shuffle")

	var cards []string
	if card != "" {
		cards = strings.Split(card, ",")
	}

	deck, err := app.DeckService.Create(
		r.Context(),
		strings.ToLower(shuffle) == "true",
		cards...,
	)
	if err != nil {
		if errors.Is(err, deckService.ErrUnknownCode) {
			app.RespondError(w, ErrBadRequest, http.StatusBadRequest)
			return
		}

		app.RespondError(w, err, http.StatusInternalServerError)
		return
	}

	app.RespondJSON(
		w,
		map[string]interface{}{
			"deck_id":   deck.ID,
			"shuffled":  deck.Shuffled,
			"remaining": len(deck.RemainingCards()),
		},
	)
}

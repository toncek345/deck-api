package api

import (
	"deck-api/api/decks"
	"deck-api/service/deck"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func New(deckService deck.Service) http.Handler {
	decksApp := decks.App{
		DeckService: deckService,
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Mount("/", decksApp.Router())

	return r
}

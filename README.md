## Deck API
This is a small deck API which supports following actions: 

* create deck action - which accepts flag to be shuffled and/or custom card codes
* open deck action - which returns given deck id and all remaining cards
* draw card action - which draws a card from the deck with given id and optionally accepts number of cards

### Dependencies
This is written in golang version 1.14 and minimal version is 1.13. The reason behind it
is this program uses new error wrapping introduced in 1.13.

### Running

#### API
API can be run with `go run main.go server` or it can be built with `go build` and then
ran with `./deck-api server`. Program optionally supports -p (--port) flag to define
custom port otherwise it runs on port `9000`.

#### Tests
Tests can be run with `go test ./...`.

### API docs

	POST /decks?shuffle=true&cards=2H,2D,AS => create deck
	shuffle - optional parameter for shuffling cards
	cards:
		- optional parameter for specifying cards
		- cards consist of value and suit first letter or value 
			- values: [A, 2...10, J, Q, K] (ace, jack, queen, king)
			- suits: [H, D, S, C] (hearts, diamonds, spades, clubs)

	GET /decks/{deckID} => open deck
	
	
	GET /decks/{deckID}/draw?amount=13 => draw cards with optional parameter amount
	


	

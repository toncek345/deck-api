package deck

import (
	"fmt"
	"strconv"
)

func cardFromCode(code string) (*Card, error) {
	if !(len(code) == 2 || len(code) == 3) {
		return nil, ErrUnknownCode
	}

	var suit Suit
	for _, s := range AllSuits() {
		if code[len(code)-1] == s[0] {
			suit = s
			break
		}
	}

	if suit == "" {
		return nil, fmt.Errorf("service/deck: unknown suit: %w", ErrUnknownCode)
	}

	var value string
	switch code[:len(code)-1] {
	case "A":
		value = "ACE"
	case "J":
		value = "JACK"
	case "Q":
		value = "QUEEN"
	case "K":
		value = "KING"
	default:
		codeParsed, err := strconv.Atoi(code[:len(code)-1])
		if err != nil {
			return nil, fmt.Errorf("service/deck: unable to parse value: %w", ErrUnknownCode)
		}

		if codeParsed < 2 || codeParsed > 10 {
			return nil, fmt.Errorf("service/deck: unknown card value: %d: %w", codeParsed, ErrUnknownCode)
		}

		value = code[:len(code)-1]
	}

	return &Card{
		Value: value,
		Suit:  suit,
		Code:  code,
	}, nil
}

func generateCards(codes ...string) ([]*Card, error) {
	var cards []*Card

	if len(codes) != 0 {
		cards = make([]*Card, 0, len(codes))

		for _, code := range codes {
			card, err := cardFromCode(code)
			if err != nil {
				return nil, fmt.Errorf("service/deck: unknown code: %s: %w", code, err)
			}

			cards = append(cards, card)
		}

		return cards, nil
	}

	cards = make([]*Card, 0, 52)

	for _, suit := range AllSuits() {
		cards = append(cards, &Card{
			Value: "ACE",
			Suit:  suit,
			Code:  fmt.Sprintf("A%c", suit[0]),
		})

		for i := 2; i < 11; i++ {
			cards = append(cards, &Card{
				Value: strconv.Itoa(i),
				Suit:  suit,
				Code:  fmt.Sprintf("%d%c", i, suit[0]),
			})
		}

		cards = append(cards, &Card{
			Value: "JACK",
			Suit:  suit,
			Code:  fmt.Sprintf("J%c", suit[0]),
		})

		cards = append(cards, &Card{
			Value: "QUEEN",
			Suit:  suit,
			Code:  fmt.Sprintf("Q%c", suit[0]),
		})

		cards = append(cards, &Card{
			Value: "KING",
			Suit:  suit,
			Code:  fmt.Sprintf("K%c", suit[0]),
		})
	}

	return cards, nil
}

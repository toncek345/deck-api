package deck

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeckRemainingCards(t *testing.T) {
	assert.Equal(
		t,
		[]*Card{&Card{}, &Card{}},
		(&Deck{
			Cards: []*Card{&Card{}, &Card{}},
		}).RemainingCards())

	assert.Equal(
		t,
		[]*Card{&Card{}},
		(&Deck{
			drawn: 1,
			Cards: []*Card{&Card{}, &Card{}},
		}).RemainingCards(),
		"remaining cards should be equal")

	assert.Equal(
		t,
		[]*Card{},
		(&Deck{
			drawn: 2,
			Cards: []*Card{&Card{}, &Card{}},
		}).RemainingCards(),
		"remaining cards should be equal")

}

package dominionizer

import (
	"fmt"
	"sort"
	"strings"
)

// Kingdom ...
type Kingdom struct {
	cards []Card
}

func (k *Kingdom) AddCard(kc Card) {
	k.cards = append(k.cards, kc)
}

func (k *Kingdom) GetCards() []Card {
	return k.cards
}

func (k *Kingdom) SortByName() {
	sort.Slice(k.cards, func(i, j int) bool {
		left := k.cards[i]
		right := k.cards[j]

		return left.CompareName(right) < 0
	})
}

func (k *Kingdom) SortByCost() {
	sort.Slice(k.cards, func(i, j int) bool {
		left := k.cards[i]
		right := k.cards[j]

		return left.CompareCost(right) > 0
	})
}

func (k Kingdom) String() string {
	var sb strings.Builder
	for _, card := range k.cards {
		sb.WriteString(fmt.Sprintf("%s\n", card))
	}
	return sb.String()
}

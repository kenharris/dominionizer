package dominionizer

import (
	"fmt"
	"sort"
	"strings"
)

// Kingdom represents a collection of ten kingdom cards randomized based on rules
type Kingdom struct {
	cards []Card
}

<<<<<<< HEAD
// AddCard adds a card to the current kingdom
func (k *Kingdom) AddCard(c Card) {
	k.cards = append(k.cards, c)
}

// GetCards gets cards from the kingdom
=======
func (k *Kingdom) AddCard(kc Card) {
	k.cards = append(k.cards, kc)
}

>>>>>>> eb3d01a75c8b0472dd656ea922ffaea7647c4be9
func (k *Kingdom) GetCards() []Card {
	return k.cards
}

// SortByName sorts cards in kingdom by name
func (k *Kingdom) SortByName() {
	sort.Slice(k.cards, func(i, j int) bool {
		left := k.cards[i]
		right := k.cards[j]

		return left.CompareName(right) < 0
	})
}

// SortByCost sorts cards in kingdom by cost
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

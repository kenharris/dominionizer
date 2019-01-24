package dominionizer

import (
	"fmt"
	"sort"
	"strings"
)

// Kingdom is a type made from a slice of Card objects.
type Kingdom []Card

// SortByName sorts cards in kingdom by name
func (cards Kingdom) SortByName() {
	sort.Slice(cards, func(i, j int) bool {
		left := cards[i]
		right := cards[j]

		return left.CompareName(right) < 0
	})
}

// SortByCost sorts cards in kingdom by cost
func (cards Kingdom) SortByCost() {
	sort.Slice(cards, func(i, j int) bool {
		left := cards[i]
		right := cards[j]

		return left.CompareCost(right) > 0
	})
}

func (k Kingdom) String() string {
	var sb strings.Builder
	for _, card := range k {
		sb.WriteString(fmt.Sprintf("%s\n", card))
	}
	return sb.String()
}

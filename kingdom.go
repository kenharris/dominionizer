package dominionizer

import (
	"fmt"
	"sort"
	"strings"
)

// Kingdom is a type made from a slice of Card objects.
type Kingdom []Card

// RandomizeKingdom is a function which generates a randomized kingdom.
func RandomizeKingdom(numCards int) Kingdom {
	k := []Card{}
	setCards := shuffle(getSetCards(), 5)

	blacklistMap := map[string]bool{}
	for _, c := range Blacklist {
		blacklistMap[c] = true
	}

	cardIndex := 0
	for len(k) < numCards && cardIndex < len(setCards) {
		cardToConsider := setCards[cardIndex]
		cardIndex++

		if blacklistMap[cardToConsider.Name] == true {
			continue
		}

		k = append(k, cardToConsider)
	}

	return k
}

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

package models

import (
	"fmt"
	"sort"
	"strings"
)

func (k *Kingdom) SortByName() {
	sort.Slice(k.Cards, func(i, j int) bool {
		left := k.Cards[i]
		right := k.Cards[j]

		return left.CompareName(right) < 0
	})
}

func (k *Kingdom) SortByCost() {
	sort.Slice(k.Cards, func(i, j int) bool {
		left := k.Cards[i]
		right := k.Cards[j]

		return left.CompareCost(right) > 0
	})
}

func (k Kingdom) String() string {
	var sb strings.Builder
	for _, card := range k.Cards {
		sb.WriteString(fmt.Sprintf("%s\n", card))
	}
	return sb.String()
}

// Kingdom ...
type Kingdom struct {
	Cards []KingdomCard
}

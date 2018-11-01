package dominionizer

import (
	"fmt"
	"sort"
	"strings"
)

type Kingdom []Card

// Sets defines the list of sets which are to be included in the randomization.
var Sets []SetName

// Blacklist is a slice of card names for cards which shouldn't be included in randomization.
var Blacklist []string

// MustInclude is a slice of card names for cards which must be included in randomization.
var MustInclude []string

// TypeRules is a map of card types to ints specifying a minimum number of card types which should be included in randomization.
var TypeRules map[CardType]int

// CardReader is an interface which defines how to fetch card data.
var CardReader CardDataReader

var allCards []Card

func init() {
	Sets = []SetName{}
	Blacklist = []string{}
	MustInclude = []string{}
	TypeRules = map[CardType]int{}

	allCards = []Card{}
}

// SetTypeRule allows for a card type rule to be set for a minimum number of cards to be included in randomization.
func SetTypeRule(ct CardType, num int) {
	if TypeRules == nil {
		TypeRules = map[CardType]int{}
	}
	TypeRules[ct] = num
}

func getSetCards() []Card {
	sets := map[SetName]bool{}
	for _, s := range Sets {
		sets[s] = true
	}

	if len(allCards) == 0 {
		allCards = CardReader.ReadCards()
	}
	retCards := []Card{}

	for _, c := range allCards {
		if sets[c.Set] == true {
			retCards = append(retCards, c)
		}
	}

	return retCards
}

// RandomizeKingdom is a function which generates a randomized kingdom.
func RandomizeKingdom(numCards int) Kingdom {
	k := []Card{}
	setCards := shuffle(getSetCards())

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

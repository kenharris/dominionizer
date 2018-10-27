package dominionizer

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func shuffle(cards []Card) []Card {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	// We start at the end of the slice, inserting our random
	// values one at a time.
	for n := len(cards); n > 0; n-- {
		randIndex := r.Intn(n)
		// We swap the value at index n-1 and the random index
		// to move our randomly chosen value to the end of the
		// slice, and to move the value that was at n-1 into our
		// unshuffled portion of the slice.
		cards[n-1], cards[randIndex] = cards[randIndex], cards[n-1]
	}

	retCards := cards

	return retCards
}

func (c Card) CompareCost(c2 Card) int {
	return c.Cost.Compare(c2.Cost)
}

func (c Card) CompareName(c2 Card) int {
	return strings.Compare(c.Name, c2.Name)
}

func (c Card) String() string {
	return fmt.Sprintf("Name: %s - Set: %s - Cost: %s", c.Name, c.Set, c.Cost)
}

func (c Card) hasType(t CardType) bool {
	for _, ct := range c.Types {
		if ct == t {
			return true
		}
	}

	return false
}

func (c Card) hasCategory(cat CardCategory) bool {
	for _, cc := range c.Categories {
		if cc == cat {
			return true
		}
	}

	return false
}

// Card ...Card
type Card struct {
	Name       string
	Cost       CardCost
	Set        SetName
	Types      []CardType
	Categories []CardCategory
}

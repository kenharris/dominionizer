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

// CompareCost compares two cards' costs
func (kc Card) CompareCost(kc2 Card) int {
	return kc.Cost.Compare(kc2.Cost)
}

// CompareName compares two cards' names
func (kc Card) CompareName(kc2 Card) int {
	return strings.Compare(kc.Name, kc2.Name)
}

func (kc Card) String() string {
	return fmt.Sprintf("Name: %s - Set: %s - Cost: %s", kc.Name, kc.Set, kc.Cost)
}

// Card represents a card from a Dominion set.
type Card struct {
	Name string
	Cost CardCost
	Set  SetName
}

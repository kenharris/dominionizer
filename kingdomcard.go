package dominionizer

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func shuffle(cards []KingdomCard) []KingdomCard {
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

func (kc KingdomCard) CompareCost(kc2 KingdomCard) int {
	return kc.Cost.Compare(kc2.Cost)
}

func (kc KingdomCard) CompareName(kc2 KingdomCard) int {
	return strings.Compare(kc.Name, kc2.Name)
}

func (kc KingdomCard) String() string {
	return fmt.Sprintf("Name: %s - Kingdom: %s - Cost: %s", kc.Name, kc.Kingdom, kc.Cost)
}

// KingdomCard ...KingdomCard
type KingdomCard struct {
	Name    string
	Cost    CardCost
	Kingdom SetName
}

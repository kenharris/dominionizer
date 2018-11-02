package dominionizer

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Card represents a card from a Dominion set.
type Card struct {
	Name       string
	Cost       CardCost
	Set        SetName
	Types      []CardType
	Categories []CardCategory
	TopText    string
	BottomText string
}

// TextToHTML formats card text to HTML-friendly output.
func (c Card) TextToHTML() string {
	t := strings.Replace(c.TopText, "\n", "<br />", -1)
	b := strings.Replace(c.BottomText, "\n", "<br />", -1)

	return fmt.Sprintf("%s\n<hr />\n%s", t, b)
}

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
func (c Card) CompareCost(c2 Card) int {
	return c.Cost.Compare(c2.Cost)
}

// CompareName compares two cards' names
func (c Card) CompareName(c2 Card) int {
	return strings.Compare(c.Name, c2.Name)
}

func (c Card) String() string {
	return fmt.Sprintf("Name: %s - Set: %s - Cost: %s - Types: %s - Categories: %s", c.Name, c.Set, c.Cost, c.Types, c.Categories)
}

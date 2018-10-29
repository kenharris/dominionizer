package dominionizer_test

import (
	"testing"

	"github.com/kenharris/dominionizer"
)

func Test_CardsEqualCost(t *testing.T) {
	c1 := dominionizer.Card{
		Name: "Chapel",
		Cost: dominionizer.CardCost{Coins: 2, Potions: 3, Debt: 0}}
	c2 := dominionizer.Card{
		Name: "Woodcutter",
		Cost: dominionizer.CardCost{Coins: 2, Potions: 3, Debt: 0}}

	if c1.CompareCost(c2) != 0 {
		t.Error()
	}
}

func Test_CardsEqualName(t *testing.T) {
	c1 := dominionizer.Card{
		Name: "Chapel",
		Cost: dominionizer.CardCost{Coins: 2, Potions: 3, Debt: 0}}
	c2 := dominionizer.Card{
		Name: "Chapel",
		Cost: dominionizer.CardCost{Coins: 2, Potions: 3, Debt: 0}}

	if c1.CompareName(c2) != 0 {
		t.Error()
	}
}

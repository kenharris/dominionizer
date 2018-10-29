package dominionizer_test

import (
	"testing"

	"github.com/kenharris/dominionizer"
)

func Test_CardsEqualCost(t *testing.T) {
	kc1 := dominionizer.Card{
		Name: "Chapel",
		Cost: dominionizer.CardCost{Coins: 2, Potions: 3, Debt: 0}}
	kc2 := dominionizer.Card{
		Name: "Woodcutter",
		Cost: dominionizer.CardCost{Coins: 2, Potions: 3, Debt: 0}}

	if kc1.CompareCost(kc2) != 0 {
		t.Error()
	}
}

func Test_CardsEqualName(t *testing.T) {
	kc1 := dominionizer.Card{
		Name: "Chapel",
		Cost: dominionizer.CardCost{Coins: 2, Potions: 3, Debt: 0}}
	kc2 := dominionizer.Card{
		Name: "Chapel",
		Cost: dominionizer.CardCost{Coins: 2, Potions: 3, Debt: 0}}

	if kc1.CompareName(kc2) != 0 {
		t.Error()
	}
}

package dominionizer_test

import (
	"testing"

	"github.com/kenharris/dominionizer"
)

func Test_KingdomSortedByCost(t *testing.T) {
	kingdom := dominionizer.Kingdom{}

	kingdom.Cards = append(kingdom.Cards,
		dominionizer.KingdomCard{
			Name: "Chapel",
			Cost: dominionizer.CardCost{
				Coins:   2,
				Potions: 0,
				Debt:    0}})
	kingdom.Cards = append(kingdom.Cards,
		dominionizer.KingdomCard{
			Name: "All Card",
			Cost: dominionizer.CardCost{
				Coins:   3,
				Potions: 2,
				Debt:    1}})
	kingdom.Cards = append(kingdom.Cards,
		dominionizer.KingdomCard{
			Name: "Potion Card",
			Cost: dominionizer.CardCost{
				Coins:   0,
				Potions: 2,
				Debt:    0}})
	kingdom.Cards = append(kingdom.Cards,
		dominionizer.KingdomCard{
			Name: "Debt Card",
			Cost: dominionizer.CardCost{
				Coins:   0,
				Potions: 0,
				Debt:    3}})

	kingdom.SortByCost()

	if kingdom.Cards[0].CompareCost(kingdom.Cards[1]) < 0 ||
		kingdom.Cards[1].CompareCost(kingdom.Cards[2]) < 0 ||
		kingdom.Cards[2].CompareCost(kingdom.Cards[3]) < 0 {
		t.Error()
	}
}

func Test_KingdomSortedByName(t *testing.T) {
	kingdom := dominionizer.Kingdom{}

	kingdom.Cards = append(kingdom.Cards,
		dominionizer.KingdomCard{
			Name: "Chapel",
			Cost: dominionizer.CardCost{
				Coins:   2,
				Potions: 0,
				Debt:    0}})
	kingdom.Cards = append(kingdom.Cards,
		dominionizer.KingdomCard{
			Name: "All Card",
			Cost: dominionizer.CardCost{
				Coins:   3,
				Potions: 2,
				Debt:    1}})
	kingdom.Cards = append(kingdom.Cards,
		dominionizer.KingdomCard{
			Name: "Potion Card",
			Cost: dominionizer.CardCost{
				Coins:   0,
				Potions: 2,
				Debt:    0}})
	kingdom.Cards = append(kingdom.Cards,
		dominionizer.KingdomCard{
			Name: "Debt Card",
			Cost: dominionizer.CardCost{
				Coins:   0,
				Potions: 0,
				Debt:    3}})

	kingdom.SortByName()

	if kingdom.Cards[0].CompareName(kingdom.Cards[1]) > 0 ||
		kingdom.Cards[1].CompareName(kingdom.Cards[2]) > 0 ||
		kingdom.Cards[2].CompareName(kingdom.Cards[3]) > 0 {
		t.Error()
	}
}

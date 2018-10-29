package dominionizer_test

import (
	"testing"

	"github.com/kenharris/dominionizer"
)

func Test_KingdomSortedByCost(t *testing.T) {
	kingdom := dominionizer.Kingdom{}

	kingdom.AddCard(dominionizer.Card{
		Name: "Chapel",
		Cost: dominionizer.CardCost{
			Coins:   2,
			Potions: 0,
			Debt:    0}})
	kingdom.AddCard(dominionizer.Card{
		Name: "All Card",
		Cost: dominionizer.CardCost{
			Coins:   3,
			Potions: 2,
			Debt:    1}})
	kingdom.AddCard(dominionizer.Card{
		Name: "Potion Card",
		Cost: dominionizer.CardCost{
			Coins:   0,
			Potions: 2,
			Debt:    0}})
	kingdom.AddCard(dominionizer.Card{
		Name: "Debt Card",
		Cost: dominionizer.CardCost{
			Coins:   0,
			Potions: 0,
			Debt:    3}})

	kingdom.SortByCost()

	cards := kingdom.GetCards()
	prevCard := cards[0]
	for _, currentCard := range cards[1:] {
		if prevCard.CompareCost(currentCard) < 0 {
			t.Errorf("'%s' should not appear before '%s'", prevCard.Cost, currentCard.Cost)
		}

		prevCard = currentCard
	}
}

func Test_KingdomSortedByName(t *testing.T) {
	kingdom := dominionizer.Kingdom{}

	kingdom.AddCard(dominionizer.Card{
		Name: "Chapel",
		Cost: dominionizer.CardCost{
			Coins:   2,
			Potions: 0,
			Debt:    0}})
	kingdom.AddCard(dominionizer.Card{
		Name: "All Card",
		Cost: dominionizer.CardCost{
			Coins:   3,
			Potions: 2,
			Debt:    1}})
	kingdom.AddCard(dominionizer.Card{
		Name: "Potion Card",
		Cost: dominionizer.CardCost{
			Coins:   0,
			Potions: 2,
			Debt:    0}})
	kingdom.AddCard(dominionizer.Card{
		Name: "Debt Card",
		Cost: dominionizer.CardCost{
			Coins:   0,
			Potions: 0,
			Debt:    3}})

	kingdom.SortByName()

	cards := kingdom.GetCards()

	prevCard := cards[0]
	for _, currentCard := range cards[1:] {
		if prevCard.CompareName(currentCard) > 0 {
			t.Errorf("'%s' should not appear before '%s'", prevCard.Name, currentCard.Name)
		}

		prevCard = currentCard
	}
}

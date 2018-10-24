package models_test

import (
	"testing"

	"github.com/kenharris/dominionizer/models"
)

func Test_CardsEqualCost(t *testing.T) {
	kc1 := models.KingdomCard{Name: "Chapel", Cost: models.CardCost{Coins: 2, Potions: 3, Debt: 0}}
	kc2 := models.KingdomCard{Name: "Woodcutter", Cost: models.CardCost{Coins: 2, Potions: 3, Debt: 0}}

	if kc1.CompareCost(kc2) != 0 {
		t.Error()
	}
}

func Test_CardsEqualName(t *testing.T) {
	kc1 := models.KingdomCard{Name: "Chapel", Cost: models.CardCost{Coins: 2, Potions: 3, Debt: 0}}
	kc2 := models.KingdomCard{Name: "Chapel", Cost: models.CardCost{Coins: 2, Potions: 3, Debt: 0}}

	if kc1.CompareName(kc2) != 0 {
		t.Error()
	}
}

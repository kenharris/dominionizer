package dominionizer_test

import (
	"testing"

	"github.com/kenharris/dominionizer"
)

func Test_CardCostEqual(t *testing.T) {
	cc1 := dominionizer.CardCost{Coins: 2, Potions: 3, Debt: 0}
	cc2 := dominionizer.CardCost{Coins: 2, Potions: 3, Debt: 0}

	if cc1.Compare(cc2) != 0 {
		t.Error()
	}
}

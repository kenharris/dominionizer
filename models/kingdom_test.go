package models_test

import (
	"testing"

	"github.com/kenharris/dominionizer/models"
)

func Test_KingdomSortedByCost(t *testing.T) {
	kingdom := models.Kingdom{}

	kingdom.Cards = append(kingdom.Cards, models.KingdomCard{Name: "Chapel", Cost: models.CardCost{Coins: 2, Potions: 0, Debt: 0}})
	kingdom.Cards = append(kingdom.Cards, models.KingdomCard{Name: "All Card", Cost: models.CardCost{Coins: 3, Potions: 2, Debt: 1}})
	kingdom.Cards = append(kingdom.Cards, models.KingdomCard{Name: "Potion Card", Cost: models.CardCost{Coins: 0, Potions: 2, Debt: 0}})
	kingdom.Cards = append(kingdom.Cards, models.KingdomCard{Name: "Debt Card", Cost: models.CardCost{Coins: 0, Potions: 0, Debt: 3}})

	kingdom.SortByCost()

	if kingdom.Cards[0].CompareCost(kingdom.Cards[1]) < 0 ||
		kingdom.Cards[1].CompareCost(kingdom.Cards[2]) < 0 ||
		kingdom.Cards[2].CompareCost(kingdom.Cards[3]) < 0 {
		t.Error()
	}
}

func Test_KingdomSortedByName(t *testing.T) {
	kingdom := models.Kingdom{}

	kingdom.Cards = append(kingdom.Cards, models.KingdomCard{Name: "Chapel", Cost: models.CardCost{Coins: 2, Potions: 0, Debt: 0}})
	kingdom.Cards = append(kingdom.Cards, models.KingdomCard{Name: "All Card", Cost: models.CardCost{Coins: 3, Potions: 2, Debt: 1}})
	kingdom.Cards = append(kingdom.Cards, models.KingdomCard{Name: "Potion Card", Cost: models.CardCost{Coins: 0, Potions: 2, Debt: 0}})
	kingdom.Cards = append(kingdom.Cards, models.KingdomCard{Name: "Debt Card", Cost: models.CardCost{Coins: 0, Potions: 0, Debt: 3}})

	kingdom.SortByName()

	if kingdom.Cards[0].CompareName(kingdom.Cards[1]) > 0 ||
		kingdom.Cards[1].CompareName(kingdom.Cards[2]) > 0 ||
		kingdom.Cards[2].CompareName(kingdom.Cards[3]) > 0 {
		t.Error()
	}
}

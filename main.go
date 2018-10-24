package main

import (
	"fmt"

	"github.com/kenharris/dominion-shuffle/models"
	// "context"
	// "log"
)

func main() {
	kingdom := models.Kingdom{}
	kingdom.Cards = append(kingdom.Cards, models.KingdomCard{Name: "Chapel", Cost: models.CardCost{Coins: 2, Potions: 0, Debt: 0}})
	kingdom.Cards = append(kingdom.Cards, models.KingdomCard{Name: "All Card", Cost: models.CardCost{Coins: 3, Potions: 2, Debt: 1}})
	kingdom.Cards = append(kingdom.Cards, models.KingdomCard{Name: "Potion Card", Cost: models.CardCost{Coins: 0, Potions: 2, Debt: 0}})
	kingdom.Cards = append(kingdom.Cards, models.KingdomCard{Name: "Debt Card", Cost: models.CardCost{Coins: 0, Potions: 0, Debt: 3}})
	fmt.Println(kingdom)

	kingdom.SortByCost()
	fmt.Println(kingdom)

	kingdom.SortByName()
	fmt.Println(kingdom)

	fmt.Println("Hello world!")
}

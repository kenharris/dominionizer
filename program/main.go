package main

import (
	"fmt"

	"github.com/kenharris/dominionizer"
)

func main() {
	shuffler := dominionizer.NewShuffler()
	shuffler.IncludeSets(dominionizer.Dominion, dominionizer.Intrigue, dominionizer.Prosperity)
	shuffler.BlacklistCards("Chapel", "Bandit", "Mine", "Library", "Cellar", "Sentry", "Council Room")
	shuffler.UnblacklistCards("Chapel", "Smithy")
	shuffler.SetTypeRule(dominionizer.CardTypeAttack, 1)
	kingdom := shuffler.RandomizeKingdom(60)

	fmt.Println("Kingdom:")
	fmt.Println(kingdom)

	kingdom.SortByCost()
	fmt.Printf("Kingdom - Sorted by Cost (%d):\n", len(kingdom.GetCards()))
	fmt.Println(kingdom)

	kingdom.SortByName()
	fmt.Println("Kingdom (Sorted by Name):")
	fmt.Println(kingdom)

	fmt.Println("Hello world!")
}

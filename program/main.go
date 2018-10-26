package main

import (
	"fmt"

	"github.com/kenharris/dominionizer"
)

func main() {
	shuffler := dominionizer.NewShuffler()
	shuffler.IncludeSets(dominionizer.Dominion, dominionizer.Intrigue, dominionizer.Prosperity)
	shuffler.BlacklistCards("Chapel", "Bandit", "Mine", "Library", "Cellar", "Sentry", "Council Room")
	shuffler.UnblacklistCards("Chapel")
	kingdom := shuffler.RandomizeKingdom(10)

	fmt.Println("Kingdom:")
	fmt.Println(kingdom)

	kingdom.SortByCost()
	fmt.Println("Kingdom (Sorted by Cost):")
	fmt.Println(kingdom)

	kingdom.SortByName()
	fmt.Println("Kingdom (Sorted by Name):")
	fmt.Println(kingdom)

	fmt.Println("Hello world!")
}

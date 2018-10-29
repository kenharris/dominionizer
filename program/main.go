package main

import (
	"fmt"

	"github.com/kenharris/dominionizer"
	"github.com/kenharris/dominionizer/json"
)

func main() {
	shuffler := dominionizer.Shuffler{
		CardReader: json.CardReader{FileName: "../json/cards.json"},
		Sets:       []dominionizer.SetName{dominionizer.Dominion, dominionizer.Intrigue, dominionizer.Prosperity},
		Blacklist:  []string{"Chapel", "Bandit", "Mine", "Library", "Cellar", "Sentry", "Council Room"},
	}

	shuffler.SetTypeRule(dominionizer.CardTypeAttack, 1)
	kingdom := shuffler.RandomizeKingdom(10)

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

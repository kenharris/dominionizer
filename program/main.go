package main

import (
	"fmt"

	"github.com/kenharris/dominionizer"
	"github.com/kenharris/dominionizer/json"
)

func main() {
	dominionizer.CardReader = json.CardReader{FileName: "../json/cards.json"}
	dominionizer.Sets = []dominionizer.SetName{dominionizer.SetDominion, dominionizer.SetIntrigue, dominionizer.SetProsperity}
	dominionizer.Blacklist = []string{"Chapel", "Bandit", "Mine", "Library", "Cellar", "Sentry", "Council Room"}

	kingdom := dominionizer.RandomizeKingdom(10)

	// shuffler.SetTypeRule(dominionizer.CardTypeAttack, 1)
	// kingdom := shuffler.RandomizeKingdom(10)

	fmt.Println("Kingdom:")
	fmt.Println(kingdom)

	kingdom.SortByCost()
	fmt.Printf("Kingdom - Sorted by Cost (%d):\n", len(kingdom))
	fmt.Println(kingdom)

	kingdom.SortByName()
	fmt.Println("Kingdom (Sorted by Name):")
	fmt.Println(kingdom)
}

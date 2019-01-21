package main

import (
	"encoding/json"
	"fmt"

	"github.com/kenharris/dominionizer"
	djson "github.com/kenharris/dominionizer/json"
)

func prettyPrintJSON(v interface{}) {
	prettyJSON, err := json.MarshalIndent(v, "", "    ")
	fmt.Println(string(prettyJSON))
	fmt.Println(err)
}

func main() {
	dominionizer.CardReader = djson.CardReader{FilePath: "../json"}
	// dominionizer.Sets = []dominionizer.SetName{dominionizer.SetDominion, dominionizer.SetIntrigue, dominionizer.SetSeaside}
	dominionizer.Sets = []dominionizer.SetName{dominionizer.SetDominion, dominionizer.SetAlchemy}
	dominionizer.Blacklist = []string{"Chapel", "Bandit", "Mine", "Library", "Cellar", "Sentry", "Council Room"}

	kingdom := dominionizer.RandomizeKingdom(10)
	prettyPrintJSON(kingdom)

	// fmt.Println("Kingdom:")
	// fmt.Println(kingdom)

	// kingdom.SortByCost()
	// fmt.Printf("Kingdom - Sorted by Cost (%d):\n", len(kingdom))
	// fmt.Println(kingdom)

	// kingdom.SortByName()
	// fmt.Println("Kingdom (Sorted by Name):")
	// fmt.Println(kingdom)
}

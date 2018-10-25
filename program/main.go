package main

import (
	"fmt" // "context"
	// "log"

	"github.com/kenharris/dominionizer"
)

func main() {
	kingdom := dominionizer.RandomizeKingdom()

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

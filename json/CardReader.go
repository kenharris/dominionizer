package json

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"

	"github.com/kenharris/dominionizer"
)

// CardReader is a JSON-backed mechanism for reading card data
type CardReader struct {
	dominionizer.CardDataReader
	FileName string
}

type cardRepository struct {
	random         *rand.Rand
	cards          []dominionizer.Card
	exhaustedCards []dominionizer.Card
}

type cardDataCardCost struct {
	Coins   int
	Potions int
	Debt    int
}

type cardDataCardText struct {
	AboveLine string
	BelowLine string
}

type cardDataCard struct {
	Name        string
	Types       []string
	Categories  []string
	Cost        cardDataCardCost
	Description string
	CardText    cardDataCardText
}

type cardDataSet struct {
	Name  string
	Cards []cardDataCard
}

type cardData struct {
	Sets []cardDataSet
}

func (cdcc cardDataCardCost) ToRepo() dominionizer.CardCost {
	return dominionizer.CardCost{Coins: cdcc.Coins, Potions: cdcc.Potions, Debt: cdcc.Debt}
}

func (cd cardData) ToRepo() []dominionizer.Card {
	retCards := []dominionizer.Card{}
	for _, cds := range cd.Sets {
		currentSet := cds.Name

		for _, cdc := range cds.Cards {
			var c dominionizer.Card

			c.Name = cdc.Name
			c.Set = dominionizer.SetNameFromString(currentSet)
			c.Cost = cdc.Cost.ToRepo()

			if len(cdc.Types) > 0 {
				c.Types = []dominionizer.CardType{}
				for _, ct := range cdc.Types {
					ctt := dominionizer.CardTypeFromString(ct)
					if ctt != dominionizer.CardTypeUnknown {
						c.Types = append(c.Types, ctt)
					}
				}
			}

			if len(cdc.Categories) > 0 {
				c.Categories = []dominionizer.CardCategory{}
				for _, ct := range cdc.Categories {
					ctt := dominionizer.CardCategoryFromString(ct)
					if ctt != dominionizer.CardCategoryUnknown {
						c.Categories = append(c.Categories, ctt)
					}
				}
			}

			retCards = append(retCards, c)
		}
	}
	return retCards
}

// ReadCards is a function defined to read cards from data source.
func (cr CardReader) ReadCards() []dominionizer.Card {
	var cards cardData
	file, err := os.Open(cr.FileName)
	if err != nil {
		fmt.Printf("Error loading cards from file: %v", err)
		os.Exit(1)
	}
	json.NewDecoder(file).Decode(&cards)

	return cards.ToRepo()
}

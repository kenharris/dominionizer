package json

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"

	"github.com/kenharris/dominionizer"
)

// CardReader is a JSON-backed mechanism for reading card data
type CardReader struct {
	dominionizer.CardDataReader
	FilePath string
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
	Types       []dominionizer.CardType
	Categories  []dominionizer.CardCategory
	Cost        cardDataCardCost
	Description string
	CardText    cardDataCardText
}

type cardDataSet struct {
	Name  dominionizer.SetName
	Cards []cardDataCard
}

func (cdcc cardDataCardCost) ToRepo() dominionizer.CardCost {
	return dominionizer.CardCost{Coins: cdcc.Coins, Potions: cdcc.Potions, Debt: cdcc.Debt}
}

func (cds cardDataSet) ToRepo() []dominionizer.Card {
	retCards := []dominionizer.Card{}

	for _, cdc := range cds.Cards {
		var c dominionizer.Card

		c.Name = cdc.Name
		c.Set = cds.Name
		c.Cost = cdc.Cost.ToRepo()
		c.TopText = cdc.CardText.AboveLine
		c.BottomText = cdc.CardText.BelowLine

		if len(cdc.Types) > 0 {
			c.Types = []dominionizer.CardType{}
			for _, ct := range cdc.Types {
				if ct != dominionizer.CardTypeUnknown {
					c.Types = append(c.Types, ct)
				}
			}
		}

		if len(cdc.Categories) > 0 {
			c.Categories = []dominionizer.CardCategory{}
			for _, ct := range cdc.Categories {
				if ct != dominionizer.CardCategoryUnknown {
					c.Categories = append(c.Categories, ct)
				}
			}
		}

		retCards = append(retCards, c)
	}
	return retCards
}

// ReadCards is a function defined to read cards from data source.
func (cr CardReader) ReadCards() []dominionizer.Card {
	cards := []dominionizer.Card{}

	files, err := ioutil.ReadDir(cr.FilePath)
	if err != nil {
		fmt.Printf("Error loading files in directory: %v", err)
		os.Exit(1)
	}

	for _, fi := range files {
		if fi.IsDir() || filepath.Ext(fi.Name()) != ".json" {
			continue
		}

		file, err := os.Open(fmt.Sprintf("%s/%s", cr.FilePath, fi.Name()))
		if err != nil {
			fmt.Printf("Error loading cards from file: %v", err)
			os.Exit(1)
		}

		var set cardDataSet
		json.NewDecoder(file).Decode(&set)

		for _, card := range set.ToRepo() {
			cards = append(cards, card)
		}
	}

	return cards
}

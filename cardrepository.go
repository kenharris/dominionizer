package dominionizer

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
)

type cardRepository struct {
	random         *rand.Rand
	cards          []Card
	exhaustedCards []Card
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
	Name  SetName
	Cards []cardDataCard
}

type cardData struct {
	Sets []cardDataSet
}

func (cdcc cardDataCardCost) ToRepo() CardCost {
	return CardCost{Coins: cdcc.Coins, Potions: cdcc.Potions, Debt: cdcc.Debt}
}

func (cd cardData) ToRepo() []Card {
	retCards := []Card{}
	for _, cds := range cd.Sets {
		currentSet := cds.Name

		for _, cdc := range cds.Cards {
			var c Card

			c.Name = cdc.Name
			c.Set = currentSet
			c.Cost = cdc.Cost.ToRepo()

			retCards = append(retCards, c)
		}
	}
	return retCards
}

func (repo *cardRepository) LoadCards() {
	if len(repo.cards) > 0 {
		repo.cards = repo.cards[:0]
	}

	// b := []byte(`{"Sets": [{"Name": "Dominion","Cards":[{"name": "Cellar","types": ["action"],"cost": {"coins": 2},"description": "+1 Action\nDiscard any number of cards, then draw that many."},{"name": "Chapel","types": ["action","trasher"],"cost": {"coins": 2},"description": "+1 Action\nDiscard any number of cards, then draw that many."}]}]}`)
	var cards cardData
	file, err := os.Open("../data/cards.json")
	if err != nil {
		fmt.Errorf("Error loading cards from file: %v", err)
		os.Exit(1)
	}
	json.NewDecoder(file).Decode(&cards)
	// json.Unmarshal(b, &cards)

	repo.cards = cards.ToRepo()
}

func (repo cardRepository) getSetCards(includedSets map[SetName]bool) []Card {
	retCards := []Card{}

	for _, c := range repo.cards {
		if includedSets[c.Set] == true {
			retCards = append(retCards, c)
		}
	}

	return retCards
}

func (repo *cardRepository) GetCard() {

}

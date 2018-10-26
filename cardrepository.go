package dominionizer

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
)

type cardRepository struct {
	random         *rand.Rand
	cards          []KingdomCard
	exhaustedCards []KingdomCard
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

func (cd cardData) ToRepo() []KingdomCard {
	retCards := []KingdomCard{}
	for _, cds := range cd.Sets {
		currentKingdom := cds.Name

		for _, cdc := range cds.Cards {
			var kc KingdomCard

			kc.Name = cdc.Name
			kc.Kingdom = currentKingdom
			kc.Cost = cdc.Cost.ToRepo()

			retCards = append(retCards, kc)
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

	// repo.cards = append(repo.cards, KingdomCard{Name: "Cellar", Kingdom: Dominion, Cost: CardCost{Coins: 2}})
	// repo.cards = append(repo.cards, KingdomCard{Name: "Chapel", Kingdom: Dominion, Cost: CardCost{Coins: 2}})
	// repo.cards = append(repo.cards, KingdomCard{Name: "Moat", Kingdom: Dominion, Cost: CardCost{Coins: 2}})
	// repo.cards = append(repo.cards, KingdomCard{Name: "Village", Kingdom: Dominion, Cost: CardCost{Coins: 3}})

	// repo.cards = append(repo.cards, KingdomCard{Name: "Pawn", Kingdom: Intrigue, Cost: CardCost{Coins: 2}})
	// repo.cards = append(repo.cards, KingdomCard{Name: "Masquerade", Kingdom: Intrigue, Cost: CardCost{Coins: 3}})
	// repo.cards = append(repo.cards, KingdomCard{Name: "Shanty Town", Kingdom: Intrigue, Cost: CardCost{Coins: 3}})
	// repo.cards = append(repo.cards, KingdomCard{Name: "Bridge", Kingdom: Intrigue, Cost: CardCost{Coins: 4}})

	// repo.cards = append(repo.cards, KingdomCard{Name: "Loan", Kingdom: Prosperity, Cost: CardCost{Coins: 3}})
	// repo.cards = append(repo.cards, KingdomCard{Name: "Monument", Kingdom: Prosperity, Cost: CardCost{Coins: 4}})
	// repo.cards = append(repo.cards, KingdomCard{Name: "Mountebank", Kingdom: Prosperity, Cost: CardCost{Coins: 5}})
	// repo.cards = append(repo.cards, KingdomCard{Name: "Rabble", Kingdom: Prosperity, Cost: CardCost{Coins: 5}})
}

func (repo cardRepository) getSetCards(includedKingdoms map[SetName]bool) []KingdomCard {
	retCards := []KingdomCard{}

	for _, c := range repo.cards {
		if includedKingdoms[c.Kingdom] == true {
			retCards = append(retCards, c)
		}
	}

	return retCards
}

func (repo *cardRepository) GetCard() {

}

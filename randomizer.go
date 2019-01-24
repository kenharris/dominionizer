package dominionizer

// Randomizer is a type that allows for managing of randomizing a kingdom
type Randomizer struct {
	id string
	sr StateReader
	cr CardDataReader
}

// RandomizeKingdom is a function which generates a randomized kingdom.
func (r Randomizer) RandomizeKingdom(numCards int) Kingdom {
	s := r.sr.ReadState(r.id)

	k := []Card{}
	setCards := shuffle(getSetCards(r.cr.ReadCards(), s.Sets), 5)

	blacklistMap := map[string]bool{}
	for _, c := range s.Blacklist {
		blacklistMap[c] = true
	}

	cardIndex := 0
	for len(k) < numCards && cardIndex < len(setCards) {
		cardToConsider := setCards[cardIndex]
		cardIndex++

		if blacklistMap[cardToConsider.Name] == true {
			continue
		}

		k = append(k, cardToConsider)
	}

	return k
}

func getSetCards(cards []Card, setNames []SetName) []Card {
	sets := map[SetName]bool{}
	for _, s := range setNames {
		sets[s] = true
	}

	retCards := []Card{}
	for _, c := range cards {
		if sets[c.Set] == true {
			retCards = append(retCards, c)
		}
	}

	return retCards
}

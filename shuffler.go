package dominionizer

// Shuffler is a type which is used to control generation of kingdoms by observing rules and abiding by them
type Shuffler struct {
	Sets        []SetName
	Blacklist   []string
	MustInclude []string
	TypeRules   map[CardType]int
	CardReader  CardDataReader

	allCards []Card
}

func (s *Shuffler) SetTypeRule(ct CardType, num int) {
	if s.TypeRules == nil {
		s.TypeRules = map[CardType]int{}
	}
	s.TypeRules[ct] = num
}

func (s *Shuffler) getSetCards() []Card {
	sets := map[SetName]bool{}
	for _, s := range s.Sets {
		sets[s] = true
	}

	if len(s.allCards) == 0 {
		s.allCards = s.CardReader.ReadCards()
	}
	retCards := []Card{}

	for _, c := range s.allCards {
		if sets[c.Set] == true {
			retCards = append(retCards, c)
		}
	}

	return retCards
}

// RandomizeKingdom is a function which generates a randomized kingdom
func (s *Shuffler) RandomizeKingdom(numCards int) Kingdom {
	k := Kingdom{}
	setCards := shuffle(s.getSetCards())

	blacklistMap := map[string]bool{}
	for _, c := range s.Blacklist {
		blacklistMap[c] = true
	}

	cardIndex := 0
	for len(k.cards) < numCards && cardIndex < len(setCards) {
		cardToConsider := setCards[cardIndex]
		cardIndex++

		if blacklistMap[cardToConsider.Name] == true {
			continue
		}

		k.cards = append(k.cards, cardToConsider)
	}

	return k
}

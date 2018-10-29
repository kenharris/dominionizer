package dominionizer

// Shuffler is a type which is used to control generation of kingdoms by observing rules and abiding by them
type Shuffler struct {
	IncludedSets     map[SetName]bool
	BlacklistedCards map[string]bool
	MustIncludeCards []string
	TypeRules        map[CardType]int
	AllCards         []Card
	CardReader       CardDataReader
}

func (s *Shuffler) BlacklistCards(cards ...string) {
	if s.BlacklistedCards == nil {
		s.BlacklistedCards = map[string]bool{}
	}

	for _, kc := range cards {
		s.BlacklistedCards[kc] = true
	}
}

func (s *Shuffler) UnblacklistCards(cards ...string) {
	if s.BlacklistedCards == nil {
		s.BlacklistedCards = map[string]bool{}
	}

	for _, kc := range cards {
		s.BlacklistedCards[kc] = false
	}
}

func (s *Shuffler) IncludeSets(k ...SetName) {
	if s.IncludedSets == nil {
		s.IncludedSets = map[SetName]bool{}
	}

	for _, kn := range k {
		s.IncludedSets[kn] = true
	}
}

func (s *Shuffler) ExcludeSets(k ...SetName) {
	if s.IncludedSets == nil {
		s.IncludedSets = map[SetName]bool{}
	}

	for _, kn := range k {
		s.IncludedSets[kn] = false
	}
}

func (s *Shuffler) AddMustIncludeCards(cards ...string) {
	for _, c := range cards {
		s.MustIncludeCards = append(s.MustIncludeCards, c)
	}
}

func (s *Shuffler) SetTypeRule(ct CardType, num int) {
	if s.TypeRules == nil {
		s.TypeRules = map[CardType]int{}
	}
	s.TypeRules[ct] = num
}

func excludeBlacklistCards(cards []Card, blackList []string) []Card {
	var retCards []Card

	for _, kc := range cards {
		found := false
		for _, blc := range blackList {
			if blc == kc.Name {
				found = true
				break
			}
		}

		if !found {
			retCards = append(retCards, kc)
		}
	}

	return retCards
}

func (s *Shuffler) getSetCards() []Card {
	if len(s.AllCards) == 0 {
		s.AllCards = s.CardReader.ReadCards()
	}
	retCards := []Card{}

	for _, c := range s.AllCards {
		if s.IncludedSets[c.Set] == true {
			retCards = append(retCards, c)
		}
	}

	return retCards
}

// RandomizeKingdom is a function which generates a randomized kingdom
func (s *Shuffler) RandomizeKingdom(numCards int) Kingdom {
	k := Kingdom{}
	setCards := shuffle(s.getSetCards())

	cardIndex := 0
	for len(k.cards) < numCards && cardIndex < len(setCards) {
		cardToConsider := setCards[cardIndex]
		cardIndex++

		if s.BlacklistedCards[cardToConsider.Name] == true {
			continue
		}

		k.cards = append(k.cards, cardToConsider)
	}

	return k
}

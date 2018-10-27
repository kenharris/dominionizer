package dominionizer

// Shuffler is a type which is used to control generation of kingdoms by observing ShufflerOptions rules and abiding by them
type shuffler struct {
	Options ShufflerOptions
	repo    cardRepository
}

// NewShuffler returns a shuffler configured with defaults
func NewShuffler() shuffler {
	s := shuffler{}
	s.repo = cardRepository{}
	s.repo.LoadCards()

	s.Options = ShufflerOptions{}
	s.Options.IncludedSets = map[SetName]bool{}
	s.Options.BlacklistedCards = map[string]bool{}
	s.Options.MustIncludeCards = []string{}
	s.Options.TypeRules = map[CardType]int{}

	return s
}

func (s *shuffler) BlacklistCards(cards ...string) {
	for _, kc := range cards {
		s.Options.BlacklistedCards[kc] = true
	}
}

func (s *shuffler) UnblacklistCards(cards ...string) {
	for _, kc := range cards {
		s.Options.BlacklistedCards[kc] = false
	}
}

func (s *shuffler) IncludeSets(k ...SetName) {
	for _, kn := range k {
		s.Options.IncludedSets[kn] = true
	}
}

func (s *shuffler) ExcludeSets(k ...SetName) {
	for _, kn := range k {
		s.Options.IncludedSets[kn] = false
	}
}

func (s *shuffler) AddMustIncludeCards(cards ...string) {
	for _, c := range cards {
		s.Options.MustIncludeCards = append(s.Options.MustIncludeCards, c)
	}
}

func (s *shuffler) SetTypeRule(ct CardType, num int) {
	s.Options.TypeRules[ct] = num
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

// RandomizeKingdom is a function which generates a randomized kingdom
func (s shuffler) RandomizeKingdom(numCards int) Kingdom {
	k := Kingdom{}
	setCards := shuffle(s.repo.getSetCards(s.Options.IncludedSets))

	cardIndex := 0
	for len(k.cards) < numCards && cardIndex < len(setCards) {
		cardToConsider := setCards[cardIndex]
		cardIndex++

		if s.Options.BlacklistedCards[cardToConsider.Name] == true {
			continue
		}

		k.cards = append(k.cards, cardToConsider)
	}

	return k
}

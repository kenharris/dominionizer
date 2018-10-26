package dominionizer

// SetName is a type representing name of kingdoms
type SetName int

// Set names
const (
	Dominion    SetName = 0
	Intrigue    SetName = 1
	Seaside     SetName = 2
	Alchemy     SetName = 3
	Prosperity  SetName = 4
	Cornucopia  SetName = 5
	Hinterlands SetName = 6
	DarkAges    SetName = 7
	Guilds      SetName = 8
	Adventures  SetName = 9
	Empires     SetName = 10
	Nocturne    SetName = 11
	Renaissance SetName = 12
)

const (
	firstKingdom SetName = Dominion - 1
	lastKingdom  SetName = Renaissance + 1
)

func (k SetName) String() string {
	names := []string{
		"Dominion",
		"Intrigue",
		"Seaside",
		"Alchemy",
		"Prosperity",
		"Cornucopia",
		"Hinterlands",
		"Dark Ages",
		"Guilds",
		"Adventures",
		"Empires",
		"Nocturne",
		"Renaissance",
	}

	if k <= firstKingdom || k >= lastKingdom {
		return "Unknown"
	}

	return names[k]
}

// Shuffler is a type which is used to control generation of kingdoms by observing ShufflerOptions rules and abiding by them
type shuffler struct {
	Options ShufflerOptions
	repo    cardRepository
}

func NewShuffler() shuffler {
	s := shuffler{}
	s.repo = cardRepository{}
	s.repo.LoadCards()

	s.Options = ShufflerOptions{}
	s.Options.IncludedKingdoms = make(map[SetName]bool)
	s.Options.IncludedKingdoms[Dominion] = false
	s.Options.IncludedKingdoms[Intrigue] = false
	s.Options.IncludedKingdoms[Seaside] = false
	s.Options.IncludedKingdoms[Alchemy] = false
	s.Options.IncludedKingdoms[Prosperity] = false
	s.Options.IncludedKingdoms[Cornucopia] = false
	s.Options.IncludedKingdoms[Hinterlands] = false
	s.Options.IncludedKingdoms[DarkAges] = false
	s.Options.IncludedKingdoms[Guilds] = false
	s.Options.IncludedKingdoms[Adventures] = false
	s.Options.IncludedKingdoms[Empires] = false
	s.Options.IncludedKingdoms[Nocturne] = false
	s.Options.IncludedKingdoms[Renaissance] = false

	return s
}

func (s *shuffler) BlacklistCards(kc ...string) {
	for _, kc := range kc {
		s.Options.BlacklistedCards = append(s.Options.BlacklistedCards, kc)
	}
}

func (s *shuffler) UnblacklistCards(cards ...string) {
	keep := []string{}

	for _, blc := range s.Options.BlacklistedCards {
		found := false
		for _, kc := range cards {
			if blc == kc {
				found = true
				break
			}
		}

		if !found {
			keep = append(keep, blc)
		}
	}

	s.Options.BlacklistedCards = keep
}

func (s *shuffler) IncludeSets(k ...SetName) {
	for _, kn := range k {
		s.Options.IncludedKingdoms[kn] = true
	}
}

func (s *shuffler) ExcludeSets(k ...SetName) {
	for _, kn := range k {
		s.Options.IncludedKingdoms[kn] = false
	}
}

func (s *shuffler) AddMustIncludeCards(cards ...string) {
	for _, c := range cards {
		s.Options.MustIncludeCards = append(s.Options.MustIncludeCards, c)
	}
}

func excludeBlacklistCards(cards []KingdomCard, blackList []string) []KingdomCard {
	var retCards []KingdomCard

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
	setCards := shuffle(
		excludeBlacklistCards(
			s.repo.getSetCards(s.Options.IncludedKingdoms), s.Options.BlacklistedCards))

	k.cards = setCards[0:numCards]
	return k
}

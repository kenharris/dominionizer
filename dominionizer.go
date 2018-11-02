package dominionizer

// Sets defines the list of sets which are to be included in the randomization.
var Sets []SetName

// Blacklist is a slice of card names for cards which shouldn't be included in randomization.
var Blacklist []string

// MustInclude is a slice of card names for cards which must be included in randomization.
var MustInclude []string

// TypeRules is a map of card types to ints specifying a minimum number of card types which should be included in randomization.
var TypeRules map[CardType]int

// CardReader is an interface which defines how to fetch card data.
var CardReader CardDataReader

var allCards []Card

func init() {
	Sets = []SetName{}
	Blacklist = []string{}
	MustInclude = []string{}
	TypeRules = map[CardType]int{}

	allCards = []Card{}
}

// SetTypeRule allows for a card type rule to be set for a minimum number of cards to be included in randomization.
func SetTypeRule(ct CardType, num int) {
	if TypeRules == nil {
		TypeRules = map[CardType]int{}
	}
	TypeRules[ct] = num
}

func getSetCards() []Card {
	sets := map[SetName]bool{}
	for _, s := range Sets {
		sets[s] = true
	}

	if len(allCards) == 0 {
		allCards = CardReader.ReadCards()
	}
	retCards := []Card{}

	for _, c := range allCards {
		if sets[c.Set] == true {
			retCards = append(retCards, c)
		}
	}

	return retCards
}

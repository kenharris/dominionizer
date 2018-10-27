package dominionizer

// CardType is a type declaration for enumerating card types
type CardType int

// Constants representing the different card types
const (
	CardTypeAttack CardType = iota
	CardTypeDoom
	CardTypeDuration
	CardTypeFate
	CardTypeReaction
	CardTypeNight
	CardTypeReserve
	CardTypeTreasure
	CardTypeVictory
)

// CardTypes is a slice of all valid card types for iteration/validation
var CardTypes = []CardType{
	CardTypeAttack,
	CardTypeDoom,
	CardTypeDuration,
	CardTypeFate,
	CardTypeReaction,
	CardTypeNight,
	CardTypeReserve,
	CardTypeTreasure,
	CardTypeVictory,
}

func (k CardType) String() string {
	types := []string{
		"Attack",
		"Doom",
		"Duration",
		"Fate",
		"Reaction",
		"Night",
		"Reserve",
		"Treasure",
		"Victory",
	}

	if k < CardTypeAttack || k > CardTypeVictory {
		return "Unknown"
	}

	return types[k]
}

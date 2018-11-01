package dominionizer

import (
	"regexp"
	"strings"
)

// CardType is a type declaration for enumerating card types
type CardType int

// Constants representing the different card types
const (
	CardTypeUnknown CardType = iota
	CardTypeAction
	CardTypeAttack
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
	CardTypeUnknown,
	CardTypeAction,
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

var stringCardTypes = []string{
	"Unknown",
	"Action",
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

func (k CardType) String() string {
	if k < CardTypeAction || k > CardTypeVictory {
		return "Unknown"
	}

	return stringCardTypes[k]
}

// CardTypeFromString takes a string value and converts it to appropriate CardType constant
func CardTypeFromString(s string) CardType {
	re := regexp.MustCompile("[^0-9A-Za-z]")
	for i, sn := range stringCardTypes {
		src := re.ReplaceAllString(strings.ToLower(sn), "")
		dest := re.ReplaceAllString(strings.ToLower(s), "")
		if src == dest {
			return CardTypes[i]
		}
	}

	return CardTypeUnknown
}

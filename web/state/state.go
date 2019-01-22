package state

import (
	"github.com/kenharris/dominionizer"
)

// State defines randomizer state that can be persisted between requests
type State struct {
	Sets        []dominionizer.SetName
	Blacklist   []string
	MustInclude []string
	TypeRules   map[dominionizer.CardType]int
}

type stateReader interface {
	ReadState(identifier string) State
}

type stateWriter interface {
	WriteState(state State) bool
}

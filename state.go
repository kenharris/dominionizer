package dominionizer

// State defines randomizer state that can be persisted between requests
type State struct {
	Sets        []SetName
	Blacklist   []string
	MustInclude []string
	TypeRules   map[CardType]int
}

// StateReader is an interface that allows for reading randomizer state
type StateReader interface {
	ReadState(id string) State
}

// StateWriter is an interface that allows for persisting of randomizer state
type StateWriter interface {
	WriteState(id string, state State) bool
}

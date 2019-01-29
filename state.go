package dominionizer

import "strings"

// State defines randomizer state that can be persisted between requests
type State struct {
	Sets        []SetName
	Blacklist   []string
	MustInclude []string
	TypeRules   map[CardType]int
}

//RemoveCardFromBlacklist removes card with the same name as the passed-in card name
func (s *State) RemoveCardFromBlacklist(card string) {
	for i, c := range s.Blacklist {
		if strings.EqualFold(c, card) {
			s.Blacklist = append(s.Blacklist[:i-1], s.Blacklist[i+1:]...)
			break
		}
	}
}

//AddCardToBlacklist adds the card to the blacklist of the current state object, if that card doesn't already exist on the blacklist
func (s *State) AddCardToBlacklist(card string) {
	found := false
	for _, c := range s.Blacklist {
		if strings.EqualFold(c, card) {
			found = true
			break
		}
	}

	if !found {
		s.Blacklist = append(s.Blacklist, card)
	}
}

// StateReader is an interface that allows for reading randomizer state
type StateReader interface {
	ReadState(id string) State
}

// StateWriter is an interface that allows for persisting of randomizer state
type StateWriter interface {
	WriteState(id string, state State) bool
}

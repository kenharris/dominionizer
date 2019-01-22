package json

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/kenharris/dominionizer/web/state"
)

// StateManager is a type that allows managing randomizer state in a json format
type StateManager struct {
	identifier string
}

// NewStateManager creates a new instance of a json.StateManager
func NewStateManager(identifier string) *StateManager {
	sm := StateManager{identifier: identifier}
	return &sm
}

func (sm StateManager) getFileName() string {
	return fmt.Sprintf("state-%s.json", sm.identifier)
}

func (sm StateManager) ReadState() state.State {
	b, _ := ioutil.ReadFile(sm.getFileName())

	var returnState state.State
	json.Unmarshal(b, &returnState)

	return returnState
}

func (sm StateManager) WriteState(s state.State) bool {
	bytes, _ := json.Marshal(s)
	ioutil.WriteFile(sm.getFileName(), bytes, 0644)

	return true
}

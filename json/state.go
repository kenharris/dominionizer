package json

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/kenharris/dominionizer"
)

// StateReader is a type that allows reading randomizer state in a JSON format
type StateReader struct {
	FilePath string
}

// StateWriter is a type that allows writing of randomizer state to filesystem in JSON format
type StateWriter struct {
	FilePath string
}

func getFileName(filePath string, id string) string {
	fmt.Printf("\n\n%s/state-%s.json\n\n", filePath, id)
	return fmt.Sprintf("%s/state-%s.json", filePath, id)
}

// ReadState reads the randomizer state identified by id from the filesystem in JSON format
func (sr StateReader) ReadState(id string) dominionizer.State {
	b, err := ioutil.ReadFile(getFileName(sr.FilePath, id))

	if err != nil {
		fmt.Printf("Could not read state file %v", err)
		os.Exit(-1)
	}

	var returnState dominionizer.State
	json.Unmarshal(b, &returnState)

	return returnState
}

// WriteState writes the state of the randomizer identified by id to the filesystem in JSON format
func (sw StateWriter) WriteState(id string, s dominionizer.State) bool {
	fmt.Println("\nWriting State:\n-------------------")
	bytes, err := json.Marshal(s)

	if err != nil {
		fmt.Printf("Could not marshal state: %v", err)
		os.Exit(-1)
	}

	fmt.Printf("\nMarshaled bytes: %v\n", bytes)

	err = ioutil.WriteFile(getFileName(sw.FilePath, id), bytes, 0644)
	if err != nil {
		fmt.Printf("Could not write state file: %v", err)
		os.Exit(-1)
	}

	return true
}

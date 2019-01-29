package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/kenharris/dominionizer"
	djson "github.com/kenharris/dominionizer/json"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	var count int64 = 1

	fmt.Printf("Count: %d\n", count)
	if r.Method == http.MethodPost {
		// process post stuff
		intVal, err := strconv.ParseInt(r.FormValue("count"), 10, 64)
		fmt.Printf("IntVal: %d\n", intVal)
		if err == nil {
			count = intVal
		}
	}
	fmt.Printf("Count: %d\n", count)

	// sm := json.NewStateManager("test")
	// s := sm.ReadState()
	// fmt.Printf("State blacklist: %v\n", s.Blacklist)
	// s.Blacklist = append(s.Blacklist, "Chapel")
	// sm.WriteState(s)

	randomizer := dominionizer.CreateRandomizer("test", djson.StateReader{FilePath: "./state"}, djson.CardReader{FilePath: "../json"})
	kingdom := randomizer.RandomizeKingdom(10)
	kingdom.SortByName()
	t, _ := template.ParseFiles("home.html")

	vm := toHomeViewModel(kingdom)
	fmt.Printf("Count: %d\n", count)
	vm.Count = count
	vm.NextNum = count + 1
	fmt.Printf("VM Count: %d\n", vm.Count)
	fmt.Printf("VM NextNum: %d\n", vm.NextNum)
	fmt.Printf("****************************\n\n")
	t.Execute(w, vm)
}

func blacklistHandler(w http.ResponseWriter, r *http.Request) {
	sr := djson.StateReader{FilePath: "./state"}
	s := sr.ReadState("test")

	r.ParseForm()
	card := r.Form.Get("card")
	method := r.Method

	switch method {
	case "DELETE":
		s.RemoveCardFromBlacklist(card)
	case "POST":
		s.AddCardToBlacklist(card)
	}

	sw := djson.StateWriter{FilePath: "./state"}
	sw.WriteState("test", s)

	bytes, _ := json.Marshal(s.Blacklist)

	w.WriteHeader(200)
	w.Write(bytes)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/api/blacklist", blacklistHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

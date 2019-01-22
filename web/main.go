package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/kenharris/dominionizer"
	"github.com/kenharris/dominionizer/json"
	jsonState "github.com/kenharris/dominionizer/web/state/json"
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

	sm := jsonState.NewStateManager("test")
	s := sm.ReadState()
	fmt.Printf("State blacklist: %v\n", s.Blacklist)
	s.Blacklist = append(s.Blacklist, "Chapel")
	sm.WriteState(s)

	dominionizer.CardReader = json.CardReader{FilePath: "../json"}
	dominionizer.Sets = []dominionizer.SetName{
		dominionizer.SetDominion,
		dominionizer.SetAlchemy,
		dominionizer.SetIntrigue,
		dominionizer.SetSeaside,
		dominionizer.SetProsperity,
		dominionizer.SetCornucopia,
		dominionizer.SetHinterlands,
		dominionizer.SetDarkAges,
		dominionizer.SetGuilds,
		dominionizer.SetAdventures,
		dominionizer.SetEmpires,
		dominionizer.SetNocturne,
		dominionizer.SetRenaissance,
	}
	dominionizer.Blacklist = []string{"Chapel", "Bandit", "Mine", "Library", "Cellar", "Sentry", "Council Room"}

	kingdom := dominionizer.RandomizeKingdom(10)
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

func main() {
	http.HandleFunc("/", homeHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

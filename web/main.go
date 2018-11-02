package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/kenharris/dominionizer"
	"github.com/kenharris/dominionizer/json"
)

type cardViewModel struct {
	Name        string
	Cost        string
	Set         string
	Types       []string
	Categories  []string
	Description template.HTML
}

type homeViewModel struct {
	Count   int64
	NextNum int64
	Cards   []cardViewModel
}

func toCardViewModel(c dominionizer.Card) cardViewModel {
	vm := cardViewModel{
		Name: c.Name,
		Cost: c.Cost.String(),
		Set:  c.Set.String(),
	}

	for _, cat := range c.Categories {
		vm.Categories = append(vm.Categories, cat.String())
	}

	for _, typ := range c.Types {
		vm.Types = append(vm.Types, typ.String())
	}

	sb := strings.Builder{}
	sb.WriteString(c.TopText)
	if len(c.BottomText) > 0 {
		sb.WriteString("<hr />")
		sb.WriteString(c.BottomText)
	}

	vm.Description = template.HTML(sb.String())

	return vm
}

func toHomeViewModel(k dominionizer.Kingdom) homeViewModel {
	vm := homeViewModel{}
	for _, card := range k {
		vm.Cards = append(vm.Cards, toCardViewModel(card))
	}
	return vm
}

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

	dominionizer.CardReader = json.CardReader{FilePath: "../json"}
	dominionizer.Sets = []dominionizer.SetName{dominionizer.SetDominion, dominionizer.SetAlchemy, dominionizer.SetIntrigue, dominionizer.SetCornucopia}
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

package main

import (
	"html/template"
	"strings"

	"github.com/kenharris/dominionizer"
)

type cardViewModel struct {
	Name        string
	Set         string
	Types       []string
	Categories  []string
	Description template.HTML
	Coins       int
	Potions     int
	Debt        int
}

func getNonActionType(types []string) string {
	for _, t := range types {
		if strings.ToLower(t) != "action" {
			return t
		}
	}

	return ""
}

func (c cardViewModel) GetCSSClass() string {
	if len(c.Types[0]) > 0 {
		return strings.ToLower(getNonActionType(c.Types))
	}

	return ""
}

func (c cardViewModel) GetTypesHTML() string {
	return strings.Join(c.Types, ", ")
}

func toCardViewModel(c dominionizer.Card) cardViewModel {
	vm := cardViewModel{
		Name:    c.Name,
		Set:     c.Set.String(),
		Coins:   c.Cost.Coins,
		Potions: c.Cost.Potions,
		Debt:    c.Cost.Debt,
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

package main

import "github.com/kenharris/dominionizer"

type homeViewModel struct {
	Count   int64
	NextNum int64
	Cards   []cardViewModel
}

func toHomeViewModel(k dominionizer.Kingdom) homeViewModel {
	vm := homeViewModel{}
	for _, card := range k {
		vm.Cards = append(vm.Cards, toCardViewModel(card))
	}
	return vm
}

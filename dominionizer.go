package dominionizer

func RandomizeKingdom() Kingdom {
	k := Kingdom{}
	k.AddCard(KingdomCard{Name: "Chapel", Cost: CardCost{Coins: 2, Potions: 0, Debt: 0}})
	k.AddCard(KingdomCard{Name: "All Card", Cost: CardCost{Coins: 3, Potions: 2, Debt: 1}})
	k.AddCard(KingdomCard{Name: "Potion Card", Cost: CardCost{Coins: 0, Potions: 2, Debt: 0}})
	k.AddCard(KingdomCard{Name: "Debt Card", Cost: CardCost{Coins: 0, Potions: 0, Debt: 3}})

	return k
}

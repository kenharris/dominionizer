package dominionizer

// CardCategory is a custom type
type CardCategory int

// Enumeration for CardCategory type values
const (
	CardCategoryCantrip CardCategory = iota
	CardCategoryCantripMoney
	CardCategoryCurser
	CardCategoryExtraBuy
	CardCategoryGainer
	CardCategoryHandsizeAttack
	CardCategoryNonTerminalDrawer
	CardCategorySifter
	CardCategoryTerminalDrawer
	CardCategoryTerminalSilver
	CardCategoryTrasher
	CardCategoryVillage
	CardCategoryVillageWithConditions
)

// CardCategories is a slice of CardCategory type for iteration/validation
var CardCategories = []CardCategory{
	CardCategoryCantrip,
	CardCategoryCantripMoney,
	CardCategoryCurser,
	CardCategoryExtraBuy,
	CardCategoryGainer,
	CardCategoryHandsizeAttack,
	CardCategoryNonTerminalDrawer,
	CardCategorySifter,
	CardCategoryTerminalDrawer,
	CardCategoryTerminalSilver,
	CardCategoryTrasher,
	CardCategoryVillage,
	CardCategoryVillageWithConditions,
}

func (k CardCategory) String() string {
	categories := []string{
		"Cantrip",
		"Cantrip Money",
		"Curser",
		"ExtraBuy",
		"Gainer",
		"Handsize Attack",
		"Non-Terminal Drawer",
		"Sifter",
		"Terminal Drawer",
		"Terminal Silver",
		"Trasher",
		"Village",
		"Villages w/Conditions",
	}

	if k < CardCategoryCantrip || k > CardCategoryVillageWithConditions {
		return "Unknown"
	}

	return categories[k]
}

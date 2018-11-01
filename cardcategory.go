package dominionizer

import (
	"regexp"
	"strings"
)

// CardCategory is a custom type
type CardCategory int

// Enumeration for CardCategory type values
const (
	CardCategoryUnknown CardCategory = iota
	CardCategoryCantrip
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
	CardCategoryUnknown,
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

var stringCardCategories = []string{
	"Unknown",
	"Cantrip",
	"Cantrip Money",
	"Curser",
	"Extra Buy",
	"Gainer",
	"Handsize Attack",
	"Non-Terminal Drawer",
	"Sifter",
	"Terminal Drawer",
	"Terminal Silver",
	"Trasher",
	"Village",
	"Village with Conditions",
}

func (cat CardCategory) String() string {
	if cat < CardCategoryCantrip || cat > CardCategoryVillageWithConditions {
		return "Unknown"
	}

	return stringCardCategories[cat]
}

// CardCategoryFromString takes a string parameter and converts it into a CardCategory constant.
func CardCategoryFromString(s string) CardCategory {
	re := regexp.MustCompile("[^0-9A-Za-z]")
	for i, sn := range stringCardCategories {
		src := re.ReplaceAllString(strings.ToLower(sn), "")
		dest := re.ReplaceAllString(strings.ToLower(s), "")
		if src == dest {
			return CardCategories[i]
		}
	}

	return CardCategoryUnknown
}

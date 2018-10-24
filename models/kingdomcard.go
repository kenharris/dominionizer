package models

import (
	"fmt"
)

func (kc KingdomCard) CompareCost(kc2 KingdomCard) int {
	return kc.Cost.Compare(kc2.Cost)
}

func (kc KingdomCard) CompareName(kc2 KingdomCard) int {
	if kc.Name == kc2.Name {
		return 0
	}

	if kc.Name < kc2.Name {
		return -1
	}

	return 1
}

func (kc KingdomCard) String() string {
	return fmt.Sprintf("Name: %s - Cost: %s", kc.Name, kc.Cost)
}

// KingdomCard ...KingdomCard
type KingdomCard struct {
	Name string
	Cost CardCost
}

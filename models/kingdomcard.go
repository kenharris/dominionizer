package models

import (
	"fmt"
	"strings"
)

func (kc KingdomCard) CompareCost(kc2 KingdomCard) int {
	return kc.Cost.Compare(kc2.Cost)
}

func (kc KingdomCard) CompareName(kc2 KingdomCard) int {
	return strings.Compare(kc.Name, kc2.Name)
}

func (kc KingdomCard) String() string {
	return fmt.Sprintf("Name: %s - Cost: %s", kc.Name, kc.Cost)
}

// KingdomCard ...KingdomCard
type KingdomCard struct {
	Name string
	Cost CardCost
}

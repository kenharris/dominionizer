package models

import (
	"fmt"
	"strings"
)

func (cc CardCost) Compare(cc2 CardCost) int {
	if cc.Coins == cc2.Coins {
		if cc.Potions == cc2.Potions {
			if cc.Debt == cc2.Debt {
				return 0
			}

			if cc.Debt > cc2.Debt {
				return 1
			}

			return -1
		}

		if cc.Potions > cc2.Potions {
			return 1
		}

		return -1
	}

	if cc.Coins > cc2.Coins {
		return 1
	}

	return -1
}

func (cc CardCost) String() string {
	var sb strings.Builder
	first := true

	if cc.Coins > 0 {
		sb.WriteString(fmt.Sprintf("%d Coins", cc.Coins))
		first = false
	}

	if cc.Potions > 0 {
		if !first {
			sb.WriteString(", ")
		}

		sb.WriteString(fmt.Sprintf("%d Potions", cc.Potions))
		first = false
	}

	if cc.Debt > 0 {
		if !first {
			sb.WriteString(", ")
		}

		sb.WriteString(fmt.Sprintf("%d Debt", cc.Debt))
		first = false
	}

	return sb.String()
}

//CardCost ...
type CardCost struct {
	Coins   int
	Potions int
	Debt    int
}

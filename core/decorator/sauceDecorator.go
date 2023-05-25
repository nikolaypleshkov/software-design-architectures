package sauceDecorator

import (
	"fmt"

	"github.com/nikolaypleshkov/burger-restaurant/core/burgers"
)

func AddSauce(burger burgers.Burger, sauceType string) burgers.Burger {
	switch sauceType {
	case "garlic sauce":
		return &GarlicSauce{Burger: burger}
	case "ketchup":
		return &Ketchup{Burger: burger}
	default:
		return burger
	}
}

type GarlicSauce struct {
	Burger burgers.Burger
}

func (gs *GarlicSauce) GetDescription() string {
	return fmt.Sprintf("%s with Garlic Sauce", gs.Burger.GetDescription())
}

type Ketchup struct {
	Burger burgers.Burger
}

func (k *Ketchup) GetDescription() string {
	return fmt.Sprintf("%s with Ketchup", k.Burger.GetDescription())
}

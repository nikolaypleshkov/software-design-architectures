package kitchen

import "github.com/nikolaypleshkov/burger-restaurant/core/burgers"

func GetBurger(burgerType string) burgers.Burger {
	switch burgerType {
	case "king":
		return burgers.NewKingBurger()
	case "classic":
		return burgers.NewClassicBurger()
	case "vegetarian":
		return burgers.NewVegetarianBurger()
	default:
		return nil
	}
}

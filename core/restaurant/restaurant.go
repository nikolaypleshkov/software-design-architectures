package restaurant

import (
	"github.com/nikolaypleshkov/burger-restaurant/core/cashier"
	"github.com/nikolaypleshkov/burger-restaurant/core/chef"
)

type Restaurant struct{}

var restaurant *Restaurant

func GetRestaurant() *Restaurant {
	if restaurant != nil {
		return restaurant
	}
	restaurant = &Restaurant{}

	return restaurant
}

func (r *Restaurant) PlaceOrder(burgerType, sauceType string) {
	c := cashier.NewCashier()
	ch := chef.NewChef()

	order := c.TakeOrder(burgerType, sauceType)

	ch.PrepareBurger(order)
}

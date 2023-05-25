package cashier

import (
	"fmt"

	"github.com/nikolaypleshkov/burger-restaurant/core/chef"
)

var pf = fmt.Printf

type Cashier struct{}

func NewCashier() *Cashier {
	return &Cashier{}
}

func (c *Cashier) TakeOrder(burgerType, sauceType string) chef.Order {

	order := chef.Order{
		BurgerType: burgerType,
		SauceType:  sauceType,
	}

	pf("Cashier: Received order - Burger Type: %s, Sauce Type: %s\n", burgerType, sauceType)

	return order
}

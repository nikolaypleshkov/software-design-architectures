package chef

import (
	"fmt"

	"github.com/nikolaypleshkov/burger-restaurant/core/burgers"
	sauceDecorator "github.com/nikolaypleshkov/burger-restaurant/core/decorator"
	"github.com/nikolaypleshkov/burger-restaurant/core/kitchen"
)

var pf = fmt.Printf

type Chef struct{}

func NewChef() *Chef {
	return &Chef{}
}
func (ch *Chef) ServeBurger(burger burgers.Burger) {
	pf("Chef: Serving %s\n", burger.GetDescription())
}

func (ch *Chef) ProcessOrder(order Order) {
	burger := ch.PrepareBurger(order)

	ch.ServeBurger(burger)
}

func (ch *Chef) PrepareBurger(order Order) burgers.Burger {
	burger := kitchen.GetBurger(order.BurgerType)

	burgerWithSauce := sauceDecorator.AddSauce(burger, order.SauceType)

	pf("Chef: Prepared %s Burger with %s Sauce\n", order.BurgerType, order.SauceType)

	ch.ServeBurger(burgerWithSauce)

	return burgerWithSauce
}

type Order struct {
	BurgerType string
	SauceType  string
}

package main

import (
	"fmt"

	"github.com/nikolaypleshkov/burger-restaurant/core/restaurant"
)

var pl = fmt.Println

func main() {
	restaurantInstance := restaurant.GetRestaurant()

	restaurantInstance.PlaceOrder("king", "ketchup")
	restaurantInstance.PlaceOrder("classic", "garlic sauce")
	restaurantInstance.PlaceOrder("vegetarian", "mustard")
}

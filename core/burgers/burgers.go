package burgers

type Burger interface {
	GetDescription() string
}

type KingBurger struct{}

func NewKingBurger() *KingBurger {
	return &KingBurger{}
}

func (kb *KingBurger) GetDescription() string {
	return "King Burger"
}

type ClassicBurger struct{}

func NewClassicBurger() *ClassicBurger {
	return &ClassicBurger{}
}

func (cb *ClassicBurger) GetDescription() string {
	return "Classic Burger"
}

type VegetarianBurger struct{}

func NewVegetarianBurger() *VegetarianBurger {
	return &VegetarianBurger{}
}

func (vb *VegetarianBurger) GetDescription() string {
	return "Vegetarian Burger"
}

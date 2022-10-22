package main

import "fmt"

type Liters float64
type Gallons float64
type Milliliters float64

func (l Liters) String() string {
	return fmt.Sprintf("%0.2f lit", l)
}

func (g Gallons) String() string {
	return fmt.Sprintf("%0.2f gal", g)
}

func (m Milliliters) String() string {
	return fmt.Sprintf("%0.2f mil", m)
}

func main() {
	var carFuel Gallons
	var busFuel Liters
	var milk Milliliters
	mil := Milliliters(4000)
	gal := Gallons(2)
	lit := Liters(6)
	carFuel = lit.ToGallons()
	busFuel = mil.ToLiters()
	milk = gal.ToMilliliters()
	fmt.Printf("Car fuel: %s, Bus fuel: %s, Milk: %s\n", carFuel, busFuel, milk)
}

func (m Milliliters) ToGallons() Gallons {
	return Gallons(m * 0.000264)
}

func (l Liters) ToGallons() Gallons {
	return Gallons(l * 0.264)
}

func (m Milliliters) ToLiters() Liters {
	return Liters(m * 0.001)
}

func (g Gallons) ToLiters() Liters {
	return Liters(g * 3.578)
}

func (l Liters) ToMilliliters() Milliliters {
	return Milliliters(l * 1000)
}

func (g Gallons) ToMilliliters() Milliliters {
	return Milliliters(g * 3578)
}

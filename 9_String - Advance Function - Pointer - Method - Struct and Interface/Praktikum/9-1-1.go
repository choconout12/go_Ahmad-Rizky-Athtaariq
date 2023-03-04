package main

import "fmt"

type Car struct {
	typeCar string
	fuelIn  float64
}

func (c *Car) calculateRange() float64 {
	fuelEfficiency := 1.5 // km per mill
	return fuelEfficiency * c.fuelIn
}

func main() {
	myCar := Car{"Bus", 1000}
	rangeInKM := myCar.calculateRange()
	fmt.Printf("Mobil %s ini bisa melaju hingga %.2f kilometer.\n", myCar.typeCar, rangeInKM)
}

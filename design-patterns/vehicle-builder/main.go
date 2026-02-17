package main

import (
	"fmt"

	"github.com/anupKanere/Golang/design-patterns/vehicle-builder/builder"
	"github.com/anupKanere/Golang/design-patterns/vehicle-builder/director"
)

func main() {
	d := &director.Director{}

	carBuilder := &builder.CarBuilder{}
	d.SetBuilder(carBuilder)
	car := d.Build()
	fmt.Printf("Car: %+v\n", car)

	bikeBuilder := &builder.BikeBuilder{}
	d.SetBuilder(bikeBuilder)
	bike := d.Build()
	fmt.Printf("Bike: %+v\n", bike)

}

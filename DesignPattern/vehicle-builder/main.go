package man

import (
	"fmt"

	"github.com/anupKanere/Golang/DesignPattern/vehicle-builder/builder"
	"github.com/anupKanere/Golang/DesignPattern/vehicle-builder/director"
)

func main() {
	d := &director.Director{}

	carBuilder := &builder.CarBuilder{}
	d.SetBuilder(carBuilder)
	car := d.Build()
	fmt.Println(car)

	bikeBuilder := &builder.BikeBuilder{}
	d.SetBuilder(bikeBuilder)
	bike := d.Build()
	fmt.Println(bike)

}

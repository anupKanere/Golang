package director

import (
	"github.com/anupKanere/Golang/DesignPattern/vehicle-builder/builder"
	"github.com/anupKanere/Golang/DesignPattern/vehicle-builder/model"
)

type Director struct {
	Builder builder.VehicleBuilder
}

func (d *Director) SetBuilder(b builder.VehicleBuilder) {
	d.Builder = b
}

func (d *Director) Build() model.Vehicle {
	d.Builder.SetType()
	d.Builder.SetEngine()
	d.Builder.SetWheel()
	return d.Builder.GetVehicle()
}

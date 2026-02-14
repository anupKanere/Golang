package builder

import "github.com/anupKanere/Golang/DesignPattern/vehicle-builder/model"

type BikeBuilder struct {
	Vehicle model.Vehicle
}

func (bb *BikeBuilder) GetVehicle() model.Vehicle {
	return bb.Vehicle
}

func (bb *BikeBuilder) SetType() {
	bb.Vehicle.Type = "bike"

}
func (bb *BikeBuilder) SetWheel() {
	bb.Vehicle.Wheels = 2
}

func (bb *BikeBuilder) SetEngine() {
	bb.Vehicle.Engine = "petrol"
}

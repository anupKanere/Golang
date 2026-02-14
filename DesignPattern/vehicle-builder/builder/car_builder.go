package builder

import "github.com/anupKanere/Golang/DesignPattern/vehicle-builder/model"

type CarBuilder struct {
	Vehicle model.Vehicle
}

func (cb *CarBuilder) GetVehicle() model.Vehicle {
	return cb.Vehicle
}

func (cb *CarBuilder) SetType() {
	cb.Vehicle.Type = "car"
}

func (cb *CarBuilder) SetWheel() {
	cb.Vehicle.Wheels = 4
}

func (cb *CarBuilder) SetEngine() {
	cb.Vehicle.Engine = "Diesel"
}

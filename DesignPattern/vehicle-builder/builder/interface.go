package builder

import "github.com/anupKanere/Golang/DesignPattern/vehicle-builder/model"

type VehicleBuilder interface {
	GetVehicle() model.Vehicle
	SetType()
	SetWheel()
	SetEngine()
}

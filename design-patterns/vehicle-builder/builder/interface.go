package builder

import "github.com/anupKanere/Golang/design-patterns/vehicle-builder/model"

type VehicleBuilder interface {
	GetVehicle() model.Vehicle
	SetType()
	SetWheel()
	SetEngine()
}

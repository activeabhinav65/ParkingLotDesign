package main

import (
	"goRepos/parking_lot_design/constants"
	"goRepos/parking_lot_design/parkinglotmanager"
)

func main() {
	newPlm := parkinglotmanager.GetParkingLotManager()

	newPlm.AddFloorAndSlots(1, 10)
	newPlm.AddFloorAndSlots(2, 20)
	newPlm.AddFloorAndSlots(3, 10)

	newPlm.FreeSlots(constants.TypeCar)

	newPlm.ParkVehicle(constants.TypeCar, 10)

	newPlm.FreeSlots(constants.TypeCar)
}

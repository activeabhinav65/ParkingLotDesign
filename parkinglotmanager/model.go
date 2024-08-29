package parkinglotmanager

import (
	"goRepos/parking_lot_design/floor"
	"goRepos/parking_lot_design/ticket"
	"goRepos/parking_lot_design/vehicle"
)

type ParkingLotManger struct {
	Tickets  map[int]ticket.Ticket
	Vehicles map[int]vehicle.Vehicle
	Floors   map[int]floor.Floor
}

func GetParkingLotManager() ParkingLotManger {
	return ParkingLotManger{
		Tickets:  make(map[int]ticket.Ticket),
		Vehicles: make(map[int]vehicle.Vehicle),
		Floors:   make(map[int]floor.Floor),
	}
}

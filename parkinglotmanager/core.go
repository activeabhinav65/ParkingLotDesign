package parkinglotmanager

import (
	"fmt"

	"goRepos/parking_lot_design/constants"
	"goRepos/parking_lot_design/floor"
	"goRepos/parking_lot_design/slot"
	"goRepos/parking_lot_design/ticket"
)

func (plm ParkingLotManger) AddFloorAndSlots(floorNumber int, numberOfSlots int) {
	if numberOfSlots < 3 {
		fmt.Println("minimum of number of slots have to be 3")
		return
	}

	var slots []slot.Slot

	slots = append(slots, slot.Slot{
		SlotNumber: 1,
		Type:       constants.TypeTruck,
		Status:     constants.FreeSlot,
	})

	slots = append(slots, slot.Slot{
		SlotNumber: 2,
		Type:       constants.TypeBike,
		Status:     constants.FreeSlot,
	})

	slots = append(slots, slot.Slot{
		SlotNumber: 3,
		Type:       constants.TypeBike,
		Status:     constants.FreeSlot,
	})

	l := numberOfSlots - 3

	for i := 0; i < l; i++ {
		cs := slot.Slot{
			SlotNumber: i + 4,
			Type:       constants.TypeCar,
			Status:     constants.FreeSlot,
		}
		slots = append(slots, cs)
	}

	flo := floor.Floor{}
	flo.Floor = slots
	plm.Floors[floorNumber] = flo
}

func (plm ParkingLotManger) AddSlotByFloorNumber(floorNumber int, slotType string, slotNumber int) {
	_, ok := plm.Floors[floorNumber]
	if !ok {
		fmt.Println("floor number does not exist")
		return
	}
	// check for slot number

	flooRe := plm.Floors[floorNumber]
	slots := flooRe.Floor

	for _, slotE := range slots {
		if slotE.SlotNumber == slotNumber {
			fmt.Println("slot number already exists")
			return
		}
	}

	slots = append(slots, slot.Slot{
		SlotNumber: slotNumber,
		Type:       slotType,
		Status:     constants.FreeSlot,
	})
	flooRe.Floor = slots

	plm.Floors[floorNumber] = flooRe
	fmt.Println("slot added on the respective floor")
}

func (plm ParkingLotManger) ParkVehicle(typ string, vi int) ticket.Ticket {
	var t ticket.Ticket
	t.TicketId = 0
	df := 0
	for fn, floorE := range plm.Floors {
		slo := floorE.Floor
		for index, s := range slo {
			if s.Type == typ && s.Status == constants.FreeSlot {
				// get ticket id
				fg := 0
				for ti, _ := range plm.Tickets {
					fg = ti
				}
				t = ticket.GetTicket(fg+1, vi, s.SlotNumber, fn)
				plm.Tickets[fg+1] = t
				s.Status = constants.BookedSlot
				slo[index] = s
				df = 1
				break
			}
			if df == 1 {
				break
			}
		}
		floorE.Floor = slo
	}

	if t.TicketId == 0 {
		fmt.Println("No slots are available right now")
		return ticket.Ticket{}
	}

	return t
}

func (plm ParkingLotManger) UnParkVehicle(ticketId int) {
	sn := plm.Tickets[ticketId].SlotNumber
	fn := plm.Tickets[ticketId].FloorNumber

	for floorId, floorInDb := range plm.Floors {
		if floorId == fn {
			slots := floorInDb.Floor

			for index, slotInDb := range slots {
				if slotInDb.SlotNumber == sn {
					slotInDb.Status = constants.FreeSlot
					slots[index] = slotInDb
				}
			}
			floorInDb.Floor = slots
		}
	}

	fmt.Println("successfully unparked vehicle")
}

func (plm ParkingLotManger) FreeSlots(vt string) {
	for floorId, floorInDb := range plm.Floors {
		slots := floorInDb.Floor
		nof := 0
		for _, slotInDb := range slots {
			if slotInDb.Status == constants.FreeSlot && slotInDb.Type == vt {
				nof++
			}
		}
		fmt.Println("Free slots on Floor Number ", floorId, " are ", nof, "for this vehicle type")
	}

	return
}

func (plm ParkingLotManger) OccupiedSlots(vt string) {
	for floorId, floorInDb := range plm.Floors {
		slots := floorInDb.Floor
		nof := 0
		for _, slotInDb := range slots {
			if slotInDb.Status == constants.BookedSlot && slotInDb.Type == vt {
				nof++
			}
		}
		fmt.Println("Booked slots on Floor Number ", floorId, " are ", nof, "for this vehicle type")
	}

	return
}

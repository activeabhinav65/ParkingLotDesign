package slot

import "goRepos/parking_lot_design/constants"

type Slot struct {
	SlotNumber int
	Status     string
	Type       string
}

func GetSlot(no int, typ string) Slot {
	return Slot{
		SlotNumber: no,
		Status:     constants.FreeSlot,
		Type:       typ,
	}
}

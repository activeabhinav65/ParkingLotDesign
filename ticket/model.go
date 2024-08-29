package ticket

type Ticket struct {
	TicketId    int
	VehicleId   int
	SlotNumber  int
	FloorNumber int
}

func GetTicket(ti, vi int, sn, fn int) Ticket {
	return Ticket{
		TicketId:    ti,
		VehicleId:   vi,
		SlotNumber:  sn,
		FloorNumber: fn,
	}
}

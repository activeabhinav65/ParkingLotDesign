package vehicle

type Vehicle struct {
	VehicleId          int
	Type               string
	RegistrationNumber string
	Color              string
}

func GetVehicle(vehicleId int, typ, registrationNo, color string) Vehicle {
	return Vehicle{
		VehicleId:          vehicleId,
		Type:               typ,
		RegistrationNumber: registrationNo,
		Color:              color,
	}
}

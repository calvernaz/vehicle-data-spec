package parking_braking_status


type ParkingBrakeStatus int

const (
	// Parking brake is not engaged (driving position)
	Inactive ParkingBrakeStatus = iota + 1
	// Parking brake is engaged (parking position)
	Active
	// There is a problem with the parking brake system
	Error
)




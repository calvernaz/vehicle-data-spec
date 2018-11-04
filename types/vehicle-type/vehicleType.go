package vehicle_type

// vehicle type
type VehicleType int

const (
	PassengerCarMini VehicleType = iota + 1
	PassengerCarLight
	PassengerCarCompact
	PassengerCarMedium
	PassengerCarHeavy
	SportUtilityVehicle
	PickupTruck
	Van
)

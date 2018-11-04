package vehicle_power

type VehiclePowerMode int

const (
	// Off - no power
	Off VehiclePowerMode = iota + 1
	// Accessory power 1
	Accessory1
	// Accessory power 2
	Accessory2
	// Running power
	Running
)

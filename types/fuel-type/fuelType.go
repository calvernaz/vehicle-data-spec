package fuel_type

type FuelType int

const (
	Gasoline FuelType = iota + 1
	Methanol
	Ethanol
	Diesel
	LPG
	CNG
	Electric
)

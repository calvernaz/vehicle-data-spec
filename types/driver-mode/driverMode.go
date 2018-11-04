package driver_mode


type DriveModeType int

const (
	Comfort DriveModeType = iota + 1
	Auto
	Sport
	Eco
	Manual
	Winter
)

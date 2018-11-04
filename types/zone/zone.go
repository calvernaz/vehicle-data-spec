package zone


//The Zone interface contains the constants that represent physical zones and logical zones
type ZoneType int

const (
	Front ZoneType = iota + 1
	// the middle position of a row
	Middle
	Right
	Left
	Rear
	// the center row
	Center
)

type Zone struct {
	// array of physical zones
	Value []string
	//  physical zone for logical driver
	Driver ZoneType
}


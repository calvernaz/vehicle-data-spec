package convertible_root_status

type ConvertibleRoofStatus int

const (
	// The convertible roof is closed
	Closed ConvertibleRoofStatus = iota + 1
	// The convertible roof is closing
	Closing
	// The convertible roof is opening
	Opening
	// The convertible roof is opene
	Opened
)

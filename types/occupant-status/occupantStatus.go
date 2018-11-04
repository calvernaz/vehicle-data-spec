package occupant_status

type OccupantStatus int

const (
	// Occupant is an adult
	Adult OccupantStatus = iota + 1
	// Occupant is a child
	Child
	// Seat is vacant
	Vacant
)

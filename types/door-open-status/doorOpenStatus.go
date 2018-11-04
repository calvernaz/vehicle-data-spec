package door_open_status

type DoorOpenStatus int

const (
	// Door is opened
	Open DoorOpenStatus = iota + 1
	// Door is ajar
	Ajar
	// Door is closed
	Closed
)
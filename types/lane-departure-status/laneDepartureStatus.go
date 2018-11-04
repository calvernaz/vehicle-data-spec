package lane_departure_status

type LaneDepartureStatus int

const (
	Off LaneDepartureStatus = iota + 1
	Pause
	Running
)


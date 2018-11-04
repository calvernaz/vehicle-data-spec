package airflow_direction


type AirflowDirection int

const (
	// Air flow is directed to the instrument panel outlets
	FrontPanel AirflowDirection = iota + 1
	// Air flow is directed to the floor outlets
	FloorDuct
	// Air flow is directed to the instrument panel outlets and the floor outlets
	BiLevel
	// Air flow is directed to the floor outlets and the windshield
	DefrostFloor
)
package wiper_control

type WiperControl int

const (
	// Wiper is not in operation
	Off WiperControl = iota + 1
	// Wipe single. It's a transient state and goes to the off mode
	Once
	// Wiper is on mode with the slowest speed
	Slowest
	// Wiper is on mode with slow speed
	Slow
	// Wiper is on mode with middle speed
	Middle
	// Wiper is on mode with fast speed
	Fast
	// Wiper is on mode with the fastest speed
	Fastest
	//     Wiper is on the automatic mode which controls wiping speed with accordance with the amount of rain
	Auto
)
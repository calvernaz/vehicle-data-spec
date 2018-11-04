package transmission_mode


type TransmissionMode int

const (
	// Transmission is in park
	Park TransmissionMode = iota + 1
	// Transmission is in reverse
	Reverse
	// Transmission is in neutral
	Neutral
	// Transmission is in low
	Low
	// Transmission is in drive
	Drive
	// Transmission is in overdrive
	Overdrive
)

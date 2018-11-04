package identification_type


type IdentificationType int

const (
	// Four digit pin number entered by user
	Pin IdentificationType = iota + 1
	// Identification by key fob
	Keyfob
	// Identification by Bluetooth device
	Bluethoot
	// Identification by NFC device
	NFC
	// Identification by fingerprint
	Fingerprint
	// Identification by camera
	Camera
	// Identification by voice
	Voice
)
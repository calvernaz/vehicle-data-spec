package alarm_status

type AlarmStatus int

const (
	// The alarm is not armed
	Disarmed AlarmStatus = iota + 1
	// The alarm is not armed
	PreArmed
	// The function is active
	Armed
	// The alarm is screaming
	Alarmed
)


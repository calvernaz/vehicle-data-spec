package button_event


type ButtonEventType int

const (
	Home ButtonEventType = iota + 1
	Back
	Search
	Call
	EndCall
	MediaPlay
	MediaNext
	MediaPrevious
	MediaPause
	VoiceRecognize
	Enter
	Left
	Right
	Up
	Down
	Press
	LongPress
	Release
)
package log_stash

import "encoding/json"

type Level int

const (
	DebugLevel Level = 1
	InfoLevel  Level = 2
	WarnLevel  Level = 3
	ErrorLevel Level = 4
)

func (s Level) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())

}

func (s Level) String() string {
	switch s {
	case DebugLevel:
		return "debug"
	case InfoLevel:
		return "info"
	case WarnLevel:
		return "warn"
	case ErrorLevel:
		return "error"
	default:
		return "其他"
	}
}

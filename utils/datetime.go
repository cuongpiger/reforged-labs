package utils

import (
	lfmt "fmt"
	ltime "time"
)

// Timestamp is a custom time type with formatted JSON output
type Timestamp struct {
	ltime.Time
}

// MarshalJSON formats Timestamp as "2006-01-02 15:04:05"
func (s Timestamp) MarshalJSON() ([]byte, error) {
	formatted := lfmt.Sprintf("\"%s\"", s.Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

func Now() Timestamp {
	return Timestamp{Time: ltime.Now()}
}

func TimestampFrom(t ltime.Time) Timestamp {
	return Timestamp{Time: t}
}

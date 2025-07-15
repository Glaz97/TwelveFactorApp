package types

import "time"

// GetCurrentUTCTime function returns current time in UTC with truncated milliseconds.
// Truncation used, because mongo stores only several digits after separator
func GetCurrentUTCTime() *time.Time {
	utcTime := time.Now().UTC().Truncate(time.Millisecond)
	return &utcTime
}

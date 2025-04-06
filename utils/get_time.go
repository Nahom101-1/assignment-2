package utils

import (
	"time"
)

// GetTimestamp returns the current time in "YYYYMMDD HH:MM" format
func GetTimestamp() string {
	return time.Now().UTC().Format(time.RFC3339)
}

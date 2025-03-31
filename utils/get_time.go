package utils

import (
	"time"
)

// GetTimestamp returns the current time in "YYYYMMDD HH:MM" format
func GetTimestamp() string {
	return time.Now().Format("20060102 15:04")
}

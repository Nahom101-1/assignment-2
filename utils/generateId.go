package utils

import (
	"crypto/rand"
	"encoding/hex"
)

// GenerateID returns a random 8-byte hex string (16 chars)
func GenerateID() string {
	b := make([]byte, 8)
	rand.Read(b)
	return hex.EncodeToString(b)
}

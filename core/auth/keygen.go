package auth

import (
	"crypto/rand"
	"encoding/hex"
)

// Key generates 'len' random bytes, encodes
// them in hex and returns a string
// representation of it.
func Key(len int) string {
	buff := make([]byte, len)
	_, err := rand.Read(buff)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(buff)
}

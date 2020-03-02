/*
Package auth provides functionality for authenticating
with service users.
*/
package auth

import (
	"crypto/rand"
	"encoding/base64"
	"log"
)

/*
KeyGen generates a random len byte string.
*/
func KeyGen(len int) string {
	key := make([]byte, len)
	_, err := rand.Read(key)

	if err != nil {
		log.Println(err)
	}

	return base64.StdEncoding.EncodeToString(key)
}

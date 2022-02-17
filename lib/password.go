package lib

import (
	"crypto/sha512"
	"encoding/base64"
)

func GeneratePassword(username string, secret string) string {
	hashed := sha512.Sum512([]byte(string(username) + secret))	
	expectedPassword := base64.StdEncoding.EncodeToString(hashed[:])

	return expectedPassword
}
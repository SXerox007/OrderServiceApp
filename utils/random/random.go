package random

import (
	"math/rand"
	"time"
)

const charset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ#@&"
const idLength = 32

func GenerateRandomID() string {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Create a byte slice of the required length
	id := make([]byte, idLength)

	// Fill each byte with a randomly chosen character from the charset
	for i := 0; i < idLength; i++ {
		id[i] = charset[rand.Intn(len(charset))]
	}

	// Convert the byte slice to a string and return it
	return string(id)
}

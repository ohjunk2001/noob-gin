package utils

import (
	"math/rand"
	"time"
)

func RandName(n int) string {
	seed := time.Now().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	const charset = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 10)
	for i := range b {
		b[i] = charset[rng.Intn(len(charset))]
	}
	return string(b)
}

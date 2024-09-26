package myutils

import (
	"math/rand"
	"time"
)

// RandNum generates a random number between 0 and n-1
func RandNum(n int) int {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	return rng.Intn(n)
}

package random

import (
	"math/rand"
	"time"
)

// Int ...
func Int(min int, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

// Float64 ...
func Float64(min, max float64) float64 {
	rand.Seed(time.Now().Unix())
	return rand.Float64()
}

// Date ...
func Date() time.Time {
	min := time.Date(1900, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Now().Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}

package util

import (
	"math/rand"
	"time"
)

func RandomInRange(min, max int) int {
	var s = rand.NewSource(time.Now().UnixNano())
	var r = rand.New(s)
	return r.Intn(max-min) + min
}

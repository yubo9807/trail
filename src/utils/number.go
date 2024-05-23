package utils

import (
	"math/rand"
	"time"
)

func NumberRandom(num int) int {
	var timestamp = time.Now().UnixNano()
	rand.Seed(timestamp)
	return rand.Intn(num)
}

package main

import (
	"math/rand"
	"time"
)

func randomZip() int {
	min := 10000
	max := 99999

	rand.Seed(time.Now().UnixNano())

	return rand.Intn(max-min) + min
}

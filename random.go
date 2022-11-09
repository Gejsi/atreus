package main

import (
	"math/rand"
	"strconv"
	"strings"
)

func randomZip() int {
	min := 10000
	max := 99999

	return rand.Intn(max-min) + min
}

func randomEmail(firstName *string) string {
	ext := []string{"com", "net", "org"}
	name := *firstName + strconv.Itoa(rand.Intn(300))

	return strings.ToLower(name) + "@1secmail." + ext[rand.Intn(len(ext))]
}

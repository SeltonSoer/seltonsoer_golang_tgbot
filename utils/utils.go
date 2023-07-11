package utils

import (
	"math/rand"
	"time"
)

func GetRandomNumberFromRange(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	randomNumber := generateRandomNumber(min, max)
	return randomNumber
}

func generateRandomNumber(min, max int) int {
	return rand.Intn(max-min+1) + min
}

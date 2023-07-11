package utils

import (
	"math/rand"
	"time"
)

func GetRandomNumberFromRange(min, max int) int {
	// Инициализируем генератор случайных чисел
	rand.Seed(time.Now().UnixNano())

	// Вызываем функцию генерации случайного числа
	randomNumber := generateRandomNumber(min, max)

	return randomNumber
}

func generateRandomNumber(min, max int) int {
	return rand.Intn(max-min+1) + min
}

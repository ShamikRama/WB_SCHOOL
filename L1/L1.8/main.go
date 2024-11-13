// Дана переменная int64.
// Разработать программу которая устанавливает i-й бит в 1 или 0.

package main

import (
	"fmt"
)

// Функция для установки i-го бита в 1
func setBit(n int64, i uint) int64 {
	return n | (1 << i)
}

// Функция для установки i-го бита в 0
func clearBit(n int64, i uint) int64 {
	return n &^ (1 << i)
}

func main() {
	var num int64 = 123 // Пример числа
	var i uint = 2      // Индекс бита, который нужно установить

	// Устанавливаем i-й бит в 1
	num = setBit(num, i)
	fmt.Printf("После установки %d-го бита в 1: %d\n", i, num)

	// Устанавливаем i-й бит в 0
	num = clearBit(num, i)
	fmt.Printf("После установки %d-го бита в 0: %d\n", i, num)
}

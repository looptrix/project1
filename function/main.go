package main

import (
	"fmt"
)

// Функция для сложения двух чисел

func sum(a int, b int) int {
	return a + b
}

// Функция для нахождения минимума из двух чисел
func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

// Функция для нахождения максимума из двух чисел
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func greet(name string) {
	fmt.Println("Привет,", name)
}
func main() {

	fmt.Println("Сумма 5 + 7 =", sum(5, 7))
	fmt.Println("Минимум из 10 и 3 =", min(10, 3))
	fmt.Println("Максимум из 10 и 3 =", max(10, 3))
	greet("Саня")
}

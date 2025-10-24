package main

import (
	"fmt"
)

// 1. Функция для суммы чисел
func sum(numbers []int) int { // функция для подсчета суммы
	total := 0 // переменная для хранения суммы
	for i := 0; i < len(numbers); i++ {
		total += numbers[i] // добавляем каждое число
	}
	return total // возращаем результат обратно
}

// 2. Функция для среднего арифмитического
func average(numbers []int) float64 {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return float64(sum) / float64(len(numbers))
}

// 3. Функция для нахождения минимума и максимума
func minMax(numbers []int) (int, int) {
	min := numbers[0]
	max := numbers[0]
	for i := 0; i < len(numbers); i++ {
		if numbers[i] < min {
			min = numbers[i]
		}
		if numbers[i] > max {
			max = numbers[i]
		}
	}
	return min, max

}

// 4. Функция для подсчета положительных и отрицательных чисел
func countPositiveNegative(numbers []int) (int, int) {
	positive := 0                       // счетчик положительных чисел
	negative := 0                       // счетчик отрицательных чисел
	for i := 0; i < len(numbers); i++ { // цикл по всем элементам слайса
		if numbers[i] > 0 { // если число больше нуля - увеичиваем счетчик положительных
			positive++
		} else if numbers[i] < 0 { // если число меньше 0- увеличиваем счетчик отрицательных
			negative++
		}
	}
	return positive, negative // возращаем оба значения

}

// 5. Главная функция

func main() {
	var n int
	fmt.Println("введите количество чисел:")
	fmt.Scan(&n)
	numbers := make([]int, 0)

	for i := 0; i < n; i++ {
		var num int
		fmt.Printf("введите числа:%d", i+1)
		fmt.Scan(&num)
		numbers = append(numbers, num)
	}
	fmt.Println("\nРезульаты:")
	fmt.Println("сумма чисел:", sum(numbers))
	fmt.Println("Среднее значение:", average(numbers))
	min, max := minMax(numbers)
	fmt.Println("Минимум:", min, "Максимум:", max)

	pos, neg := countPositiveNegative(numbers)
	fmt.Println("Положительных:", pos, "Отрицательных:", neg)

}

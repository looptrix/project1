//// Задача: Посчитать среднее арифметическое чисел
// и вывести числа больше и меньше среднего

package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Println("Введите количество чисел:")
	fmt.Scan(&n)

	numbers := make([]int, 0)
	for i := 0; i < n; i++ {
		var num int
		fmt.Printf("введите числа %d:", i+1)
		fmt.Scan(&num)

		numbers = append(numbers, num)
	}
	sumPos := 0
	sumNeg := 0
	countPos := 0
	countNeg := 0

	for i := 0; i < n; i++ {
		if numbers[i] > 0 {
			sumPos += numbers[i]
			countPos++
		} else if numbers[i] < 0 {
			sumNeg += numbers[i]
			countNeg++
		}
	}
	fmt.Println("Сумма положительных:", sumPos)
	fmt.Println("Количество положительных:", countPos)
	fmt.Println("Сумма отрицательных:", sumNeg)
	fmt.Println("Количество отрицательных:", countNeg)

}

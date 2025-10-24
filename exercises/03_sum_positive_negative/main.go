package main

import (
	"fmt"
)

func main() {

	var n int // переменная для ввода от пользователя

	fmt.Println("введите количество чисел")
	fmt.Scan(&n)

	numbers := make([]int, 0) // слайс для хранения чисел
	for i := 0; i < n; i++ {  // цикл для перебора чисел
		var num int // временная переменная для хранения чисел
		fmt.Println("введите числа")
		fmt.Scan(&num)
		numbers = append(numbers, num) // сохраняет числа в слайс
	}
	sumpositive := 0 //  сумма позитивных чисел
	sumnegative := 0 // сумма негативных чисел

	for i := 0; i < n; i++ {
		if numbers[i] > 0 {

			sumpositive += numbers[i]
		} else if numbers[i] < 0 {

			sumnegative += numbers[i]
		}
	}

	fmt.Println("сумма положительных чисел", sumpositive)
	fmt.Println("сумма отрицательных чисел", sumnegative)
}

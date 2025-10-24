package main

import "fmt"

func main() {

	var n int // переменная для хранения чисел типа int

	fmt.Println("введите количество чисел:")
	fmt.Scan(&n)

	numbers := make([]int, 0) // слайс

	for i := 0; i < n; i++ {
		var num int // временная переменная
		fmt.Printf("введите числа: %d", i+1)
		fmt.Scan(&num)
		numbers = append(numbers, num)
	}

	even := 0
	odd := 0
	sumEven := 0
	sumOdd := 0

	for i := 0; i < n; i++ {
		if numbers[i]%2 == 0 { // проверяет детится ли на 2
			even++
			sumEven += numbers[i]
		} else {
			odd++
			sumOdd += numbers[i]
		}
	}
	fmt.Println("Сумма четных чисел:", sumEven)
	fmt.Println("Сумма нечетных чисел:", sumOdd)
}

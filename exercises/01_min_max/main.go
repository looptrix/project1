package main

import "fmt"

func main() {

	var n int
	fmt.Println("введите количество чисел")
	fmt.Scan(&n)

	numbers := make([]int, 0)

	for i := 0; i < n; i++ {
		var num int
		fmt.Printf("введите числа:%d", i+1)
		fmt.Scan(&num)
		numbers = append(numbers, num)
	}
	min := numbers[0]
	max := numbers[0]

	for i := 0; i < n; i++ {
		if numbers[i] < min {
			min = numbers[i]
		} else if numbers[i] > max {
			max = numbers[i]
		}
	}

	fmt.Println("Минимальное число:", min)
	fmt.Println("Максимальное число:", max)
}

package main

import "fmt"

func sum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func avg(numbers []int) float64 {

	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return float64(sum) / float64(len(numbers))
}

func MinMax(numbers []int) (int, int) {

	min := numbers[0]
	max := numbers[0]
	for i := 0; i < len(numbers); i++ {
		if numbers[i] < min {
			min = numbers[i]
		} else if numbers[i] > max {
			max = numbers[i]
		}
	}
	return min, max
}

func EvenOdd(numbers []int) (int, int) {

	even := 0
	odd := 0
	for i := 0; i < len(numbers); i++ {
		if numbers[i]%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	return even, odd
}

func main() {

	var n int
	fmt.Println("введите число:")
	fmt.Scan(&n)

	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Введите число %d: ", i+1)
		fmt.Scan(&numbers[i])
	}

	fmt.Println("Сумма чисел:", sum(numbers))
	fmt.Println("Среднее значение:", avg(numbers))
	min, max := MinMax(numbers)
	fmt.Println("Минимум:", min, "Максимум:", max)
	even, odd := EvenOdd(numbers)
	fmt.Println("Четных чисел:", even, "Нечетных чисел:", odd)

}

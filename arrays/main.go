package main

import (
	"fmt"
)

func main() {
	// 1. Массив фиксированной длины
	var nums1 [3]int
	fmt.Println(nums1)

	// 2. Инициализация массива при создании
	nums2 := [5]int{10, 20, 30, 40, 50}
	fmt.Println(nums2)

	// 3. доступ к элементам по индексу
	fmt.Println("первый элемент:", nums2[0])
	fmt.Println("третий элемент:", nums2[2])
	fmt.Println("последний элемент:", nums2[len(nums2)-1])

	// 4. Изменение элемента массива
	nums2[2] = 99
	fmt.Println(nums2)

	// 5. Перебор массива через for
	for i := 0; i < len(nums2); i++ {
		fmt.Println("индекс", i, "значение", nums2[i])
	}

	// 6. Перебор массива через range
	fmt.Println("Перебор через range:")
	for index, value := range nums2 {
		fmt.Println(index, value)
	}

	// 7. Cумма всех элементов массива
	sum := 0
	for _, value := range nums2 {
		sum += value
	}
	fmt.Println(sum)

	// 8. Поиск минимального и максимального значения
	numbers := [6]int{45, 12, 78, 90, 4, 56}
	min := numbers[0]
	max := numbers[0]

	for _, value := range numbers {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	fmt.Println("Массив:", numbers)
	fmt.Println("Минимум:", min)
	fmt.Println("Максимум:", max)

}

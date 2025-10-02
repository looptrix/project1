package main

import "fmt"

func main() {
	// 1. Создание слайса
	numbers := []int{5, 10, 15, 20, 25}
	fmt.Println("Слайс numbers:", numbers)

	// 2. Добавление элементов
	numbers = append(numbers, 30)
	numbers = append(numbers, 35, 40)
	fmt.Println("После добавления:", numbers)

	// 3. Доступ к элементам
	fmt.Println("Первый элемент:", numbers[0])
	fmt.Println("Последний элемент:", numbers[len(numbers)-1])

	// 4. Изменение элементов
	numbers[2] = 99
	fmt.Println("После изменения:", numbers)

	// 5. Перебор слайса через for
	fmt.Println("Перебор через for:")
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("Индекс %d => Значение %d\n", i, numbers[i])
	}

	// 6. Перебор слайса через range
	fmt.Println("Перебор через range:")
	for index, value := range numbers {
		fmt.Printf("%d => %d\n", index, value)
	}

	// 7. Сумма всех элементов
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	fmt.Println("Сумма элементов:", sum)

	// 8. Минимум и максимум
	min, max := numbers[0], numbers[0]
	for _, value := range numbers {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	fmt.Println("Минимум:", min)
	fmt.Println("Максимум:", max)
}

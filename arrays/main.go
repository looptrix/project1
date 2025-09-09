package main

import "fmt"

func main() {

	var scores [6]int //1 создала массив из 6 элементов

	scores[0] = 85
	scores[1] = 90
	scores[2] = 78
	scores[3] = 92
	scores[4] = 88
	scores[5] = 75
	//2 инициализация массива
	fmt.Println("исходный массив баллов:", scores) //3 вывод всего массива

	scores[5] = 80
	fmt.Println("после изменения последнего элемента", scores) //4 изменение значения

	//5 подсчет суммы
	sum := 0
	for i := 0; i < len(scores); i++ {
		sum += scores[i]
	}
	fmt.Println("сумма всех баллов:", sum)

	//6 нахождение максимального значения
	max := scores[0]

	for i := 1; i < len(scores); i++ {
		if scores[i] > max {
			max = scores[i]
		}
	}

	fmt.Println("максимальный балл:", max)

}

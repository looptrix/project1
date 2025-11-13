// Объявляем пакет main- "это точка входа" программы
package main

//Подключаем пакет fmt - стандартная библиотека для форматированного вывода (печати текста)
import (
	"fmt"
)

func main() {
	var num []int
	for i := 0; i < 5; i++ {
		var nums int
		fmt.Println("введите числа")
		fmt.Scan(&nums)
		num = append(num, nums)
	}
}

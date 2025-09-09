package main

import "fmt"

func main() {

	// создала слайлс с 3 числами
	scores := []int{10, 20, 30}

	// вывод слайса, длины и ёмкости
	fmt.Println("слайс:", scores, "len =", len(scores), "cap =", cap(scores))

	// дабвисла 40
	scores = append(scores, 40)
	fmt.Println("после append:", scores, "len =", len(scores), "cap =", cap(scores))
}

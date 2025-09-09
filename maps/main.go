package main

import "fmt"

func main() {
	//создала мапу
	scores := make(map[string]int)

	//добавила данные
	scores["Alice"] = 85
	scores["Bob"] = 90

	//читаем данные
	fmt.Println("баллы Alice:", scores["Alice"])
	fmt.Println("баллы Bob:", scores["Bob"])

	//удаляем элемент
	delete(scores, "Alice")
	fmt.Println("после удаления:", scores)

}

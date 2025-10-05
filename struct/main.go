package main

import "fmt"

// Структура для человека
type Person struct {
	Name string
	Age  int
}

func main() {
	// Срез из структур Person
	people := []Person{
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 30},
		{Name: "Charlie", Age: 22},
	}

	// Добавление нового человека через ввод
	var name string
	var age int
	fmt.Println("Введите имя нового человека:")
	fmt.Scan(&name)
	fmt.Println("Введите возраст:")
	fmt.Scan(&age)

	newPerson := Person{Name: name, Age: age}
	people = append(people, newPerson)

	// Перебор среза и вывод информации
	fmt.Println("\nСписок людей:")
	for _, p := range people {
		fmt.Println("Имя:", p.Name, "Возраст:", p.Age)
	}

	// Поиск минимального и максимального возраста
	minAge := people[0].Age
	maxAge := people[0].Age
	sumAge := 0

	for _, p := range people {
		if p.Age < minAge {
			minAge = p.Age
		}
		if p.Age > maxAge {
			maxAge = p.Age
		}
		sumAge += p.Age
	}

	fmt.Println("\nМинимальный возраст:", minAge)
	fmt.Println("Максимальный возраст:", maxAge)
	fmt.Println("Средний возраст:", float64(sumAge)/float64(len(people)))
}

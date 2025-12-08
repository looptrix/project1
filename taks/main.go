// package main - исполняемая программа
package main

// подключаем пакет fmt
import "fmt"

// type - объявляем новый тип
// Speaker - имя типа
// interface - говорит:"это интерфейс"(набор методов, которые должен уметь тип)
// {...} - пишем список методов
type Speaker interface {
	Speak()
}

// создаем структуру Person
// Name string - поле структуры
type Person struct {
	Name string
}

// func - обьявление функции/ метода
// (p Person) - получатель
// p - имя переменной внутри метода
// Person - тип этой переменной
// метод принадлежит типу Person, и внутри метода тот объект, у которого его вызвали, будет называться p
// Speak() -  имя метода. Скобки пустые - метод ничего не принимает.
func (p Person) Speak() {
	fmt.Println("привет, я человек. Меня зовут", p.Name)
}

// тип Dog
type Dog struct {
	Name string
}

// (d Dog) - получатель - метод относится к типу Dog, и обьект внутри называется d
// Speak() - имя метода
func (d Dog) Speak() {
	fmt.Println("гав! Я собака, Меня зовут", d.Name)
}

// func MakeSpeak - обычная функция
// (s Speaker) - параметр функции
// s - имя переменной внутри функции
// Speaker - ее тип
// Функция Make Speak принимает любой объект, который удовлетворяет интерфейсу Speaker, то есть он умеет Speak

func MakeSpeak(s Speaker) {
	s.Speak()
}

// создаем объект типа Person, у которого Name = "Вася"
// теперь p - конкретный человек
// d - собака
func main() {
	p := Person{Name: "Вася"}
	d := Dog{Name: "Мили"}
	// функция MakeSpeak ожидает аргумент типа Speaker
	// p - типа Person
	// у Person есть метод Speak()- значит Person реализует Speaker.
	MakeSpeak(p)
	MakeSpeak(d)
}

package main

import "fmt"

func main() {
	// Пример 1: Ввод одного значения
	var name string

	fmt.Print("Введите ваше имя: ")
	fmt.Scan(&name)
	fmt.Println("Привет,", name)

	// Пример 2: Ввод числа
	var age int

	fmt.Print("Введите ваш возраст: ")
	fmt.Scan(&age)
	fmt.Printf("Вам %d лет\n", age)

	// Пример 3: Ввод нескольких значений
	var firstName, lastName string

	fmt.Print("Введите имя и фамилию (через пробел): ")
	fmt.Scan(&firstName, &lastName)
	fmt.Printf("Полное имя: %s %s\n", firstName, lastName)
}

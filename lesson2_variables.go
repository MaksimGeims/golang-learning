package main

import "fmt"

func main() {
	// Способ 1: Полное объявление
	var age int = 17
	var name string = "Максим"

	fmt.Println("Имя:", name)
	fmt.Println("Возраст:", age)

	// Способ 2: Короткое объявление
	city := "Москва"
	isStudent := true

	fmt.Println("Город:", city)
	fmt.Println("Студент:", isStudent)

	// Можем изменить значение переменной
	age = 18
	fmt.Println("Новый возраст:", age)
}

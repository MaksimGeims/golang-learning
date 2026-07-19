package main

import "fmt"

func main() {
	// Задание 1
	name := "Максим"
	age := 17
	favoriteNumber := 6
	doILoveProgramming := true

	fmt.Println("Меня зовут", name, "\n мне лет", age, "\n моё любимое число это", favoriteNumber, "\n люблю ли я програмированние? ответ", doILoveProgramming)

	// Задание 2
	score := 0
	
	fmt.Println("Начальный счёт:", score)

	score += 100

	fmt.Println("Новый счёт:", score)
}

package main

import (
	"fmt"
)

func main() {
	// Задание 1

	// Переменные чисел
	var number1, number2 float64

	// Запрашиваем числа у пользователя
	fmt.Print("Введите 2 числа через пробел (пример 55 25):")
	fmt.Scan(&number1, &number2)

	// Обрабатываем результаты
	answer1 := number1 + number2
	answer2 := number1 - number2
	answer3 := number1 * number2
	answer4 := number1 / number2

	// Выводим результаты. (ипользуем %v для более красивого ответа.)
	fmt.Println("Ответы:")
	fmt.Printf("1. %v + %v = %v\n", number1, number2, answer1)
	fmt.Printf("2. %v - %v = %v\n", number1, number2, answer2)
	fmt.Printf("3. %v * %v = %v\n", number1, number2, answer3)
	fmt.Printf("4. %v / %v = %v\n", number1, number2, answer4)

	// Задание 2

	//Переменная температуры
	var temperature float64

	//Запрашиваем тепмературу в °C у пользователя
	fmt.Print("Введите температуру в °C:")
	fmt.Scan(&temperature)

	// Конвертируем °C в °F
	result := temperature*9/5 + 32

	// Выводим результат
	fmt.Printf("Ответ: %v°C равняется %v°F", temperature, result)

	// Исправленное задание 3

	var fullName string

	fmt.Print("Введите полное имя: ")
	fmt.Scanln(&fullName)
	fmt.Println("Привет,", fullName)
}

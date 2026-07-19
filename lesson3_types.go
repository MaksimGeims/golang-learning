package main

import "fmt"

func main() {
	// Целые числа
	var age int = 17
	var temperature int8 = -10
	var population uint32 = 1500000
	
	fmt.Println("Возраст:", age)
	fmt.Println("Температура:", temperature, "°C")
	fmt.Println("Население:", population)
	
	// Числа lesson3_types.goс плавающей точкой
	var price float64 = 199.99
	var pi float32 = 3.14
	
	fmt.Println("Цена:", price, "руб")
	fmt.Println("Число Пи:", pi)
	
	// Строки
	var name string = "Максим"
	greeting := "Привет, " + name + "!"
	
	fmt.Println(greeting)
	
	// Логический тип
	var isStudent bool = true
	canCode := true
	
	fmt.Println("Студент:", isStudent)
	fmt.Println("Умеет программировать:", canCode)
}

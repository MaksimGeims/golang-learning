package main

import (
	"fmt"
	"strconv"
)

func main() {
	name := "Maksim"
	secondName := "Unzhakov"
	age := strconv.Itoa(17)

	fmt.Println("Меня зовут "+name, secondName+", мне "+age+" лет")

	age2 := 17

	result := fmt.Sprintf("Меня зовут %s %s, мне %d лет", name, secondName, age2)
	fmt.Println(result)
}

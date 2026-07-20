package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var city string
	var doesSport bool
	var age int
	var favoriteNumber float64

	fmt.Println("Заполнение анекеты нового пользователя.")
	fmt.Print("Введите ваше имя и фамилию: ")
	fullName, _ := reader.ReadString('\n')
	fullName = strings.TrimSpace(fullName) // Убираем \n и лишние пробелы

	fmt.Print("\nВведите ваш возраст: ")
	fmt.Scan(&age)

	fmt.Print("\nВведите ваш город: ")
	fmt.Scan(&city)

	fmt.Print("\nВведите ваше любимое число: ")
	fmt.Scan(&favoriteNumber)

	fmt.Print("\nЗанимаетесь ли вы спортом? если да ответьте true иначе ответьте false: ")
	fmt.Scan(&doesSport)

	whenUserWillBe100years := 100 - age
	multipleFavoriteNumber := favoriteNumber * 2
	var doesSportAnswer string

	if doesSport == true {
		doesSportAnswer = "Да"
	} else {
		doesSportAnswer = "Нет"
	}

	fmt.Println("╔══════════════════════════════╗")
	fmt.Println("║      АНКЕТА ПОЛЬЗОВАТЕЛЯ     ║")
	fmt.Println("╔══════════════════════════════╗")
	fmt.Println("║ Имя:", fullName)
	fmt.Println("║ Возраст:", age, "лет")
	fmt.Println("║ Город:", city)
	fmt.Println("║ Любимое числов:", favoriteNumber)
	fmt.Println("║ Удвоенное число:", multipleFavoriteNumber)
	fmt.Println("║ Занимается спортом:", doesSportAnswer)
	fmt.Println("║ До 100 лет осталось: ", whenUserWillBe100years, "года")
	fmt.Println("╚══════════════════════════════╝")
}

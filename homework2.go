package main

import "fmt"

func main() {
	currentYear := 2026
	birthYear := 2009
	age := currentYear - birthYear
	targetAge := 19
	yearsToMiddle := targetAge - age

	hoursPracticePerDay := 2.5
	totalHours := hoursPracticePerDay * 365 * float64(yearsToMiddle)

	fmt.Println("Текущий год:", currentYear,
		"\nГод рождения:", birthYear,
		"\nМой возраст:", age,
		"\nЦель стать Middle в", targetAge, "лет",
		"\nОсталось лет:", yearsToMiddle,
		"\nСколько часов в день я планирую уделять на обучение:", hoursPracticePerDay, "часа",
		"\nСколько часов практики пройдёт:", totalHours)
}

package main

import (
	"fmt"
)

func main() {
	// Товар
	firstProductName := "Сыр"
	firstProductPrice := 34.50
	firstProductAmount := 4

	secondProductName := "Колбаса"
	secondProductPrice := 169.99
	secondProductAmount := 2

	thirdProductName := "Яйца"
	thirdProductPrice := 19.99
	thirdProductAmount := 5

	// Вычисление цен

	firstProductSummaryPrice := firstProductPrice * float64(firstProductAmount)
	secondProductSummaryPrice := secondProductPrice * float64(secondProductAmount)
	thirdProductSummaryPrice := thirdProductPrice * float64(thirdProductAmount)
	totalPrice := firstProductSummaryPrice + secondProductSummaryPrice + thirdProductSummaryPrice

	// Вывод чека

	fmt.Println("======= Чек =======")
	fmt.Printf("1. %s x%v = %.2f руб\n", firstProductName, firstProductAmount, firstProductSummaryPrice)
	fmt.Printf("2. %s x%v = %.2f руб\n", secondProductName, secondProductAmount, secondProductSummaryPrice)
	fmt.Printf("3. %s x%v = %.2f руб\n", thirdProductName, thirdProductAmount, thirdProductSummaryPrice)
	fmt.Println("-------------------")
	fmt.Printf("ИТОГО: %.2f руб\n", totalPrice)
	fmt.Println("===================")
}

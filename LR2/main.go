package main

import (
	"fmt"
	"math"
)

var groupNumber int = 17
var variantNumber int = 2
var name = "Назар"

func main() {
	fmt.Printf("Мій номер за журналом: %d. Його Sin = %f.\n", groupNumber, math.Cos(float64(groupNumber)))
	day := 29
	month := 4
	year := 2005

	sum, product := birthdayMath(day, month, year)
	fmt.Printf("Сума(d + m + y) = %d, Добуток(d * m * y) = %d\n", sum, product)

	result := float64(variantNumber) * 3.14
	intResult := int(result)
	fmt.Printf("Добуток варіанта на 3.14 ( variantNumber * 3.14 ) = %f\n", result)
	fmt.Printf("Конвертації до int: %d\n", intResult)

	fmt.Printf("Номер в групі: %d, Варіант: %d, Ім'я: %s\n", groupNumber, variantNumber, name)
}

func birthdayMath(day int, month int, year int) (int, int) {
	sum := day + month + year
	product := day * month * year
	return sum, product
}

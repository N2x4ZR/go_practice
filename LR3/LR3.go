package main

import (
	"fmt"
	"math"
)

const name_number = 5 // Назар

func main() {
	var group_number = 18 // номер студента
	sum := group_number + name_number
	product := 1

	for i := sum; i != 1; i-- {
		product *= i
	}

	defer fmt.Println("Роботу виконав Назар Передерій")

	fmt.Printf("Кількість літер у імені Назар: %d\n", name_number)
	fmt.Printf("Добуток чисел %d: %d\n", sum, product)

	result := sqrt(group_number, name_number)
	fmt.Println(result)

	var groupNumber int // номер групи

	fmt.Print("Введіть номер групи: ")
	fmt.Scan(&groupNumber)

	switch {
	case groupNumber >= 75 && groupNumber <= 80:
		fmt.Println("I курс")
	case groupNumber >= 69 && groupNumber <= 74:
		fmt.Println("II курс")
	case groupNumber >= 63 && groupNumber <= 68:
		fmt.Println("III курс")
	case groupNumber >= 57 && groupNumber <= 62:
		fmt.Println("IV курс")
	default:
		fmt.Println("Група не знайдена в коледжі")
	}
}

func sqrt(groupNumber int, nameLength int) interface{} {
	result := math.Sqrt(float64(groupNumber))

	if sqrtGroup := result; sqrtGroup > float64(nameLength) {
		result = result + float64(nameLength)
		return fmt.Sprint(result, "s")
	} else if sqrtGroup < float64(nameLength) {
		result = result * float64(nameLength)
		return fmt.Sprint(result, "n")
	} else {
		return fmt.Sprintf("Корінь квадратний з номера групи (%.2f) рівний довжині імені (%d)", result, nameLength)
	}
}

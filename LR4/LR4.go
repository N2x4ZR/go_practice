package main

import (
	"fmt"
	"math"
)

func main() {
	var group_number = 62
	var variant_number = 3
	var year_number = 4.0

	fmt.Printf("Group: %d, Variant: %d, Course: %f\n", group_number, variant_number, year_number)

	sum, product, result := Calculations(group_number, variant_number, int(year_number))
	result = math.Round(result*1e6) / 1e6
	fmt.Printf("Group: %d, Variant: %d, Course: %f\n", int(sum), int(product), result)

	me := Person{
		Name:    "Назар",
		Surname: "Передерій",
		Group:   62,
		Variant: 2,
	}
	fmt.Printf("%s, %s, %d, %d\n", me.Name, me.Surname, me.Group, me.Variant)
	CalculateSum(&me)
	fmt.Printf("%s, %s, %d, %d\n", me.Name, me.Surname, me.Group, me.Variant)

	array := []float64{1, 2, sum, product, result}
	fmt.Println(array)
	array[0], array[1] = float64(group_number), float64(year_number)
	fmt.Println(array)
}

func Calculations(group_number, variant_number, year_number int) (sum, product, result float64) {

	sum = float64(group_number + variant_number + year_number)
	product = float64(group_number * variant_number * year_number)

	if variant_number%2 == 0 {
		result = math.Sin(sum)
	} else {
		result = math.Cos(product)
	}

	return
}

type Person struct {
	Name    string
	Surname string
	Group   int
	Variant int
}

func CalculateSum(person *Person) {
	sum := person.Group + person.Variant
	person.Variant = sum
}
